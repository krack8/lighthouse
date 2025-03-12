import { ChangeDetectorRef, Component, ElementRef, Inject, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { K8sNamespacesService } from '@k8s/k8s-namespaces/k8s-namespaces.service';
import icSearch from '@iconify/icons-ic/twotone-search';
import icFilter from '@iconify/icons-ic/twotone-filter-list';
import icEdit from '@iconify/icons-ic/twotone-edit';
import icExport from '@iconify/icons-ic/twotone-file-upload';
import icClear from '@iconify/icons-ic/twotone-format-clear';
import icClose from '@iconify/icons-ic/twotone-close';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { K8sPodsDetailsComponent } from '../k8s-pods-details.component';
import { WebSocketSubject, webSocket } from 'rxjs/webSocket';
import { filter, takeUntil, } from 'rxjs/operators';
import { Subject, Subscription } from 'rxjs';
import { Utils } from '@shared-ui/utils';
import { UntypedFormBuilder, UntypedFormGroup, Validators } from '@angular/forms';
import { K8sPodWebSocketService } from './k8s-pod-ws.service';
import { ToastrService } from '@sdk-ui/ui';

@Component({
  selector: 'kc-k8s-pods-container-log',
  templateUrl: './k8s-pods-container-log.component.html',
  styleUrls: ['./k8s-pods-container-log.component.scss'],
  providers: [K8sPodWebSocketService],
})
export class K8sPodsContainerLogComponent implements OnInit, OnDestroy {
  @ViewChild('appLogViewContainer', { read: ElementRef, static: false }) private appLogViewContainer!: ElementRef;
  private _destroy$: Subject<void> = new Subject();
  private  retryTimer: any;
  icSearch = icSearch;
  icEdit = icEdit;
  icFilter = icFilter;
  icClose = icClose;
  icExport = icExport;
  icClear = icClear;

  logForm: UntypedFormGroup;
  logs: string[] = [];
  wsSubject!: WebSocketSubject<any>;
  wsSubscription!: Subscription;
  isWsConnected!: boolean;
  socket: any;
  liveLogs: string = '';
  external_access_token: string;
  wsRetryCount : number = 0;
  retryTime: number = 0;

  isLoading: boolean = true;
  allowShowPrevious = this.data?.restart > 0 ? false : true;

  constructor(
    @Inject(MAT_DIALOG_DATA) public data,
    public dialogRef: MatDialogRef<K8sPodsDetailsComponent>,
    private _namespaceService: K8sNamespacesService,
    private _k8sPodWebsocketService: K8sPodWebSocketService,
    private dialog: MatDialog,
    private cd: ChangeDetectorRef,
    public snackBar: MatSnackBar,
    private _formBuilder: UntypedFormBuilder,
    private toastr: ToastrService
  ) {}

  ngOnInit(): void {
    this.logForm = this._formBuilder.group({
      container: [this.data?.container ? this.data?.container : this.data?.allContainers[0]],
      timestamps: [true],
      lines: [100],
      since: [''],
      previous: [{ value: false, disabled: this.data?.restart > 0 ? false : true }],
      follow: [true],
    });

    this.fetchLogs();
  }

  ngOnDestroy() {
    this._destroy$.next();
    this._destroy$.complete();
    if (this.wsSubscription) {
      console.log('unsubscribing wsSubscription');
      this.wsSubscription.unsubscribe();
      this.wsSubject.complete();
    }
    // if (this.retryTimer) {
    //   clearTimeout(this.retryTimer);
    // }
  }

  fetchLogs(): void {
    this.isLoading = true;
    clearTimeout(this.retryTimer);
    if (this.wsSubscription) {
      console.log('subscription is on');
      this.wsSubject.complete();
      this.wsSubscription.unsubscribe();
    }
    if (this.logForm?.get('follow').value === true) {
      this.subscribeToWs();
      // this._k8sPodWebsocketService.getToken().subscribe((res) => {
      //   if (res['status'] === 'success') {
      //     if (res['data']['external_access_token']) {
      //       this.external_access_token = res['data']['external_access_token'];
      //       this.subscribeToWs();
      //     }
      //   }
      // });
    } else {
      this.retryTime = 0;
      this.wsRetryCount = 0;
      this.isWsConnected = false;
      this.getStaticLogs();
    }
    }

  subscribeToWs(): void {
    this.isLoading = true;
    const qp = this.filterLogs();
    if (this.retryTime !== 0) {
      qp['since'] = this.retryTime;
    }
    // if (this.external_access_token) {
    //   qp['access_token'] = this.external_access_token;
    // } else {
    //   this.toastr.error('Failed: ', 'no token found!');
    //   this.isLoading = false;
    //   return;
    // }
    console.log('qp:', qp);
    const uri = this._k8sPodWebsocketService.getPodsWsUrl(this.data.pod, qp); 

    this.wsSubject = webSocket({
      url: uri,
      deserializer: (msg) => {
        if (msg.type === 'message' && msg.data) {
          return Utils.processLog(msg.data);
        }
        return null;
      },
    });

    this.wsSubscription = this.wsSubject
      .asObservable()
      .pipe(
        takeUntil(this._destroy$),
        filter((log) => log !== null),
        // retryWhen(errors =>
        //   errors.pipe(
        //     tap(err => {
        //       this.isWsConnected = false;
        //       this.applySinceFilter(5);
        //     }) ,
        //     delay(3000), // Retry after 5 seconds
        //   )
        // )
      )
      .subscribe({
        next: (log) => {
          this.retryTime = 0;
          console.log('called', this.wsRetryCount);
          this.processLog(log);
        },
        error: (err) => {
          this.isWsConnected = false;
          this.retryTime= this.retryTime + 5; // fetch logs since time in seconds
          console.log('retrying with since =' + this.retryTime + ' seconds');
          this.retryTimer = setTimeout(() => {
            this.fetchLogs();             
          } , 5000);
          },
          complete: () => {
          console.log('is complete');
          this.isWsConnected = false;
          this.isLoading = false;
        },
      });
  }

  filterLogs(): object {
    const payload = {};
    payload['namespace'] = this.data?.namespace;

    if (this.logForm?.get('lines').value >= 50 && this.logForm?.get('lines').value <= 2500) {
      payload['lines'] = this.logForm?.get('lines').value;
    } else {
      this.logForm?.get('lines').setValue(100);
      this.logForm?.get('lines').updateValueAndValidity;
      payload['lines'] = 100;
    }

    if (this.logForm?.get('container').value) {
      payload['container'] = this.logForm?.get('container').value;
    }

    if (this.logForm.get('timestamps').value === true) {
      payload['timestamps'] = 'y';
    } else {
      payload['timestamps'] = 'n';
    }
    if (this.allowShowPrevious && this.logForm?.get('previous').value === true) {
      payload['previous'] = 'y';
    } else {
      payload['previous'] = 'n';
    }
    if (this.logForm?.get('since').value && this.logForm?.get('since').value > 0) {
      payload['since'] = this.logForm?.get('since').value * 60;
    }
    if (this.retryTime == 0) this.liveLogs = '';
    return payload;
  }

  downloadLogs() {
    const blob = new Blob([this.liveLogs], { type: 'text/html' });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = 'Logs.html';
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
  }

  clearLogs() {
    this.liveLogs = '';
  }

  // applyFilter(): void {  
  //   if (this.logForm?.get('follow').value === true) {
  //     this.wsSubscription.unsubscribe();
  //     this.liveLogs = '';
  //     this.fetchLogs();
  //   } else {
  //     this.getStaticLogs();
  //   }
  // }

  // applyLinesFilter(): void {  
  //   this.liveLogs = '';
  //   if (this.logForm?.get('lines').value !== '' && this.logForm?.get('lines').value >= 50) {
  //     if (this.logForm?.get('follow').value === true) {
  //       this.wsSubscription.unsubscribe();
  //       this.fetchLogs();
  //     } else {
  //       this.getStaticLogs();
  //     } 
  //   }
  // }

  // applySinceFilter(): void {
  //   if (this.logForm?.get('since').value !== '' && this.logForm?.get('since').value >= 1) {
  //     if (this.logForm?.get('follow').value === true) {
  //       this.wsSubscription.unsubscribe();
  //       this.liveLogs = '';
  //       this.fetchLogs();
  //     } else {
  //       this.getStaticLogs();
  //     }
  //   }

  // }

  toggleShowPrevious(): void {
    this.liveLogs = '';
    if (this.logForm?.get('follow').value === false) {
      this.getStaticLogs();
    }
  }

  // toggleFollow(): void {
  //   this.liveLogs = '';
  //   if (this.logForm?.get('follow').value === true) {
  //     this.fetchLogs();
  //   } else {
  //     if (this.wsSubscription) this.wsSubscription.unsubscribe();
  //     this.getStaticLogs();
  //   }
  // }

  getStaticLogs(queryParams?): void {
    this.isLoading = true;
    this.liveLogs = '';
    const qp = this.filterLogs();
    console.log('qp:', qp);
    this._namespaceService.getLogsV1(qp, this.data.pod).subscribe((res) => {
      if (res.status === 'success') {
        this.liveLogs += this.transformTerminalLog(res.data);
        // Go to Bottom Log
        setTimeout(() => {
          this.scrollLogContainerToBottom();
        }, 100);

        this.isLoading = false;
      }
    }),
      (err) => {
        this.toastr.error('Failed: ', 'Something went wrong!');
        this.isLoading = false;
      };
  }

  processLog(data: string | string[]) {
    // Data From WebSocket
    if (typeof data === 'string') {
      const isUpdateScrollPosition =
        this.appLogViewContainer.nativeElement.scrollHeight - this.appLogViewContainer.nativeElement.scrollTop  <=
        800;
      this.liveLogs += this.transformTerminalLog(data);
      
      // Go to Bottom Log
      if (isUpdateScrollPosition) {
        this.scrollLogContainerToBottom();
        setTimeout(() => {
          //console.log('scrolling to bottom');
          this.scrollLogContainerToBottom();
        }, 10);
      }
    }
  }

  scrollLogContainerToBottom(): void {
    try {
      // clientHeight
      this.appLogViewContainer.nativeElement.scrollTop = this.appLogViewContainer.nativeElement.scrollHeight;
      this.cd.markForCheck();
    } catch (err) {
      console.warn('UPDATE Log Container Scroll: ', err);
    }
  }

  /*
   * @Params: string data;
   * @Definition: Data transform to pre html tag
   * @return: pre html tag with string type
   */
  transformTerminalLog(data: string): string {
    const match = /\r|\n/.exec(data);
    if (match) {
      data = data.replace(/\n/g, '<br>');
      data = data.replace(/\r/g, '&emsp;');
    }
    const removeUnnecessaryPrefixRegex = /\[([0-9]?[0-9])m/g;
    data = data.replace(removeUnnecessaryPrefixRegex, '');
    data = data.replace(/[\x00-\x09\x0b-\x1F]/g, ' ');
    // add margin bottom for every br tag found
    data = data.replace(/<br>/g, '<br><hr>');

    data = data.replace(/INFO/g, '<span class="text-info">INFO</span>');
    data = data.replace(/SUCCESS/g, '<span class="text-success">SUCCESS</span>');
    data = data.replace(/SUCCESSFUL/g, '<span class="text-success">SUCCESSFUL</span>');
    data = data.replace(/WARN /g, '<span class="text-warn"> WARN </span>');
    data = data.replace(/WARNING /g, '<span class="text-warn">WARNING </span>');
    data = data.replace(/error/g, '<span class="text-error mb-12">ERROR</span>');
    data = data.replace(/ERROR/g, '<span class="text-error mb-12">ERROR</span>');
    data = data.replace(/Error/g, '<span class="text-error mb-12">ERROR</span>');
    data = data.replace(/FAILED/g, '<span class="text-error mb-12">FAILED</span>');
    data = data.replace(
      /you must delete node group before delete the cluster/g,
      '<span class="text-error">you must delete node group before delete the cluster</span>'
    );
    return '<span>' + data + '</span>';
  }

}
