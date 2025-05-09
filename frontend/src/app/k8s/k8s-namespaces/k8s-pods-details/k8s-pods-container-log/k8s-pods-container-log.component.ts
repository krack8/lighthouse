import { ChangeDetectorRef, Component, ElementRef, Inject, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { UntypedFormBuilder, UntypedFormGroup } from '@angular/forms';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import icClose from '@iconify/icons-ic/twotone-close';
import icEdit from '@iconify/icons-ic/twotone-edit';
import icExport from '@iconify/icons-ic/twotone-file-upload';
import icFilter from '@iconify/icons-ic/twotone-filter-list';
import icClear from '@iconify/icons-ic/twotone-format-clear';
import icSearch from '@iconify/icons-ic/twotone-search';
import { K8sNamespacesService } from '@k8s/k8s-namespaces/k8s-namespaces.service';
import { ToastrService } from '@sdk-ui/ui';
import { Subject, Subscription } from 'rxjs';
import { K8sPodsDetailsComponent } from '../k8s-pods-details.component';
import { K8sPodLogService } from './k8s-pod-log.service';
import { delay, retryWhen, tap } from 'rxjs/operators';

@Component({
  selector: 'kc-k8s-pods-container-log',
  templateUrl: './k8s-pods-container-log.component.html',
  styleUrls: ['./k8s-pods-container-log.component.scss'],
  providers: [K8sPodLogService],
})
export class K8sPodsContainerLogComponent implements OnInit, OnDestroy {
  @ViewChild('appLogViewContainer', { read: ElementRef, static: false }) private appLogViewContainer!: ElementRef;
  private _destroy$: Subject<void> = new Subject();
  private logSubscription: Subscription | undefined;

  //icons
  icSearch = icSearch;
  icEdit = icEdit;
  icFilter = icFilter;
  icClose = icClose;
  icExport = icExport;
  icClear = icClear;

  logForm: UntypedFormGroup;
  liveLogs: string = '';
  isLoading: boolean = true;
  allowShowPrevious = this.data?.restart > 0 ? false : true;

  constructor(
    @Inject(MAT_DIALOG_DATA) public data,
    public dialogRef: MatDialogRef<K8sPodsDetailsComponent>,
    private _namespaceService: K8sNamespacesService,
    private _k8sLogService: K8sPodLogService,
    private cd: ChangeDetectorRef,
    public snackBar: MatSnackBar,
    private _formBuilder: UntypedFormBuilder,
    private toastr: ToastrService,
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
    if (this.logSubscription) {
      this.logSubscription.unsubscribe();
    }
    this._destroy$.next();
    this._destroy$.complete();
  }

  fetchLogs(): void {
    this.isLoading = true;
    if (this.logSubscription) {
      this.logSubscription.unsubscribe();
    }
    this.clearLogs();
    if (this.logForm?.get('follow').value === true) {
      this.subscribeToLogStream();
    } else {
      this.getStaticLogs();
    }
    }

  subscribeToLogStream(): void {
    this.isLoading = true;
    const qp = this.filterLogs();
    const url = this._k8sLogService.getPodsWsUrl(this.data.pod, qp); // gets http stream url
    this.logSubscription = this._k8sLogService.getPodLogsStream(url).pipe(
      retryWhen(errors =>
        errors.pipe(
          tap(error => console.error('Error:', error)),
          delay(3000) 
        )
      )
    ).subscribe(
      (log) => {
        this.processLog(log);
      },
      (error) => {
        this.toastr.error('Failed: ', 'Could not fetch logs!');
      }
    );
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

  toggleShowPrevious(): void {
    this.liveLogs = '';
    if (this.logForm?.get('follow').value === false) {
      this.getStaticLogs();
    }
  }

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
    if (typeof data === 'string') {
      const isUpdateScrollPosition =
        this.appLogViewContainer.nativeElement.scrollHeight - this.appLogViewContainer.nativeElement.scrollTop  <=
        800;
      this.liveLogs = this.transformTerminalLog(data);
      
      // Go to Bottom Log
      if (isUpdateScrollPosition) {
        setTimeout(() => {
          this.scrollLogContainerToBottom();
        }, 100);
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