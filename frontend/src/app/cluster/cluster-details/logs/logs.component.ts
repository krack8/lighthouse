import { ChangeDetectorRef, Component, ElementRef, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ClusterService } from '../../cluster.service';
import icCheckCircle from '@iconify/icons-ic/check-circle';
import icCancel from '@iconify/icons-ic/cancel';
import icAutoRenew from '@iconify/icons-ic/autorenew';
import icHourGlassEmpty from '@iconify/icons-ic/hourglass-empty';
import { distinctUntilChanged, takeWhile } from 'rxjs/operators';
import { ToastrService } from '@sdk-ui/ui';
import { environment } from '@env/environment';
import * as SockJS from 'sockjs-client';
import * as Stomp from 'stompjs';
import { RequesterService } from '@core-ui/services/requester.service';
import { trackByIndex } from '@core-ui/utils';

@Component({
  selector: 'kc-logs',
  templateUrl: './logs.component.html',
  styleUrls: ['./logs.component.scss']
})
export class LogsComponent implements OnInit, OnDestroy {
  isAlive: boolean = true;
  icCheckCircle = icCheckCircle;
  icCancel = icCancel;
  icAutoRenew = icAutoRenew;
  icHourGlassEmpty = icHourGlassEmpty;
  liveLogs = '';
  isExistLogLoading!: boolean;
  /* @ts-ignore */
  @ViewChild('appLogViewContainer', { read: ElementRef, static: true }) private appLogViewContainer!: ElementRef;
  currentState: any;

  wsSubscription: any; // Cluster Log WS Subscription
  ws: any;
  stompClient: any;

  // Cluster Creation Step
  isStepLoaded!: boolean;
  stepStatusInterval: any;
  clusterStepStatus: any = 'SUCCESS';
  clusterCreationStepStatus: any[] = [];

  clusterName: string = '';
  clusterType: string = '';
  manualClusterStep: string = '';
  clusterDetails: any;

  trackByIndex = trackByIndex;

  constructor(
    private requesterService: RequesterService,
    private router: Router,
    private route: ActivatedRoute,
    private cd: ChangeDetectorRef,
    private clusterService: ClusterService,
    private toastr: ToastrService
  ) {}

  ngOnInit(): void {
    this.clusterName = this.clusterService.clusterNameSnapshot;
    this.clusterService.currentState$
      .pipe(
        takeWhile(() => this.isAlive),
        distinctUntilChanged()
      )
      .subscribe(status => {
        this.currentState = status;
        if (status === 'DELETING') {
          if (this.stepStatusInterval) {
            clearInterval(this.stepStatusInterval);
          }
        }
      });
    this.getClusterDetails();
    this.getFullLog();
    this.stepStatusInterval = setInterval(() => {
      if (this.clusterType === 'AUTOMATED') {
        this.getClusterCreationStepStatus(true);
      } else {
        this.getManualCreationStepStatus(true);
      }
    }, 3000);
  }

  ngOnDestroy() {
    this.isAlive = false;
    if (this.wsSubscription) {
      this.wsSubscription?.unsubscribe();
    }
    if (this.stompClient && this.stompClient?.connected) {
      this.stompClient.disconnect();
    }
    clearInterval(this.stepStatusInterval);
    this.ws?.close();
  }

  getFullLog() {
    this.isExistLogLoading = true;
    this.clusterService.mcGetClustersFullLog(this.clusterName).subscribe({
      next: res => {
        this.processLog(res.data);
      },
      error: err => {
        this.isExistLogLoading = false;
        this.subscribeToWSLog();
      },
      complete: () => {
        this.isExistLogLoading = false;
        this.subscribeToWSLog();
      }
    });
  }

  subscribeToWSLog(): void {
    const token = this.requesterService.get().token;
    this.ws = new SockJS(environment.multiClusterWsEndpoint + '?authToken=' + token); //  + 'web-socket?authToken=' + token
    this.stompClient = Stomp.over(this.ws);
    // this.stompClient.debug = function (str: any) {
    //   console.log("debug", str)
    // };
    this.stompClient.debug = null;
    this.stompClient.connect(
      { 'x-auth-token': token },
      () => {
        console.log('Connected to websocket');
        this.wsSubscription = this.stompClient.subscribe(
          '/topic/clusterlog/' + this.clusterName,
          message => {
            this.processLog(message.body);
          },
          { 'x-auth-token': token }
        );
      },
      (error: any) => {
        console.log('ws err: ', error);
      }
    );
  }

  processLog(data: string | string[]) {
    // Data From WebSocket
    if (typeof data === 'string') {
      const isUpdateScrollPosition =
        this.appLogViewContainer?.nativeElement?.scrollHeight - this.appLogViewContainer?.nativeElement?.scrollTop - 40 <=
        this.appLogViewContainer?.nativeElement?.clientHeight;
      this.liveLogs += this.transformTerminalLog(data);
      // Go to Bottom Log
      if (isUpdateScrollPosition) {
        setTimeout(() => {
          this.scrollLogContainerToBottom();
        }, 10);
      }

      // Redirect When Success
      if (data.includes('Successfully Deleted Cluster...!!!')) {
        setTimeout(() => {
          this.toastr.success('Successfully Cluster Deleted!!!', 'Deleted');
          this.router.navigate(['/clusters']);
        }, 5000);
      }
      if (this.clusterStepStatus !== 'SUCCESS' && data.includes('successfully done eks cluster create and configuration setup !!!')) {
        setTimeout(() => {
          this.router.navigate(['../overview'], { relativeTo: this.route });
        }, 5000);
      }
    } else if (Array.isArray(data)) {
      // Saved Data From API
      let _logContainer = '';
      data.forEach((_log: any) => {
        _logContainer += this.transformTerminalLog(_log);
      });
      this.liveLogs += _logContainer;
      // Go to Bottom Log
      setTimeout(() => {
        this.scrollLogContainerToBottom();
      }, 100);
    }
  }

  scrollLogContainerToBottom(): void {
    try {
      // clientHeight
      if (this.appLogViewContainer) this.appLogViewContainer.nativeElement.scrollTop = this.appLogViewContainer.nativeElement.scrollHeight;
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

    data = data.replace(/INFO/g, '<span class="text-info">INFO</span>');
    data = data.replace(/SUCCESS/g, '<span class="text-success">SUCCESS</span>');
    data = data.replace(/SUCCESSFUL/g, '<span class="text-success">SUCCESSFUL</span>');
    data = data.replace(/WARN /g, '<span class="text-warn"> WARN </span>');
    data = data.replace(/WARNING /g, '<span class="text-warn">WARNING </span>');
    data = data.replace(/error/g, '<span class="text-error">ERROR</span>');
    data = data.replace(/ERROR/g, '<span class="text-error">ERROR</span>');
    data = data.replace(/Error/g, '<span class="text-error">ERROR</span>');
    data = data.replace(/FAILED/g, '<span class="text-error">FAILED</span>');
    data = data.replace(
      /you must delete node group before delete the cluster/g,
      '<span class="text-error">you must delete node group before delete the cluster</span>'
    );
    return '<pre>' + data + '</pre>';
  }

  /* -------- Cluster Step ---------- */
  getClusterCreationStepStatus(force?: boolean) {
    this.clusterService.getCreationStepStatus(this.clusterName, force).subscribe({
      next: (data: any) => {
        this.clusterStepStatus = data.clusterFinishUp;
        this.clusterCreationStepStatus = this.processStatusStep(data);
        if (!this.isStepLoaded) {
          this.isStepLoaded = true;
        }
      },
      error: err => {
        console.log(err);
      }
    });
  }

  /* --------Manual Cluster Step ---------- */
  getManualCreationStepStatus(force?: boolean) {
    this.clusterService.getManualCreationStepStatus(this.clusterName, force).subscribe({
      next: (data: any) => {
        // this.clusterStepStatus = data.clusterFinishUp;
        this.clusterCreationStepStatus = this.processManualStatusStep(data);
        this.manualClusterStep = this.clusterCreationStepStatus[this.clusterCreationStepStatus.length - 1].status;
        this.checkStepStatus(this.clusterCreationStepStatus);
        if (!this.isStepLoaded) this.isStepLoaded = true;
      },
      error: err => {
        console.log(err);
      }
    });
  }
  checkStepStatus(step: any): void {
    if (step.some(log => log.status === 'FAILED')) {
      this.manualClusterStep = 'FAILED';
      return;
    }

    // Check initiatingPlatformSetup status
    const initiatingPlatformSetupStatus = step.find(log => log.step === 'initiatingPlatformSetup').status;
    if (initiatingPlatformSetupStatus === 'SUCCESS' || initiatingPlatformSetupStatus === 'RUNNING') {
      this.manualClusterStep = 'RUNNING';
    }

    // Check setupFinishUp status
    const setupFinishUpStatus = step.find(log => log.step === 'setupFinishUp').status;
    if (setupFinishUpStatus === 'SUCCESS') {
      this.manualClusterStep = 'SUCCESS';
    }
    if (this.manualClusterStep === 'SUCCESS') {
      this.clusterService.changeCurrentState('ACTIVE');
    } else if (this.manualClusterStep === 'RUNNING') {
      this.clusterService.changeCurrentState('CREATING');
    } else {
      this.clusterService.changeCurrentState(this.manualClusterStep);
    }
  }

  processStatusStep(data: any): any[] {
    if (data['nodeGroupOperation']) {
      return [
        {
          details: 'NodeGroup Operation',
          status: data['nodeGroupOperationProgress']
        }
      ];
    }
    const m = new Map([
      ['clusterCreationJobStarted', { details: 'Starting Cluster Creation Job', status: 'PENDING' }],
      ['clusterCreation', { details: 'Cluster Creation', status: 'PENDING' }],
      ['clusterEnvironmentSetUp', { details: 'Cluster Environment Setup', status: 'PENDING' }],
      ['clusterCertManager', { details: 'Cert Manager Setup', status: 'PENDING' }],
      ['clusterReflector', { details: 'Reflector Setup', status: 'PENDING' }],
      ['clusterStorageClass', { details: 'Storage Class Creation', status: 'PENDING' }],
      // ["clusterIssuer", { details: "Cluster Issuer Setup", status: 'PENDING' }],
      // ["clusterWildCardCertificate", { details: "Generating Wildcard Certificate", status: 'PENDING' }],
      ['clusterIngressController', { details: 'Ingress Controller Setup', status: 'PENDING' }],
      ['clusterKlovercloudStack', { details: 'Deploy KC Agents', status: 'PENDING' }],
      ['clusterFinishUp', { details: 'Finishing Up', status: 'PENDING' }]
    ]);
    let hasFailed = false;
    for (let key of m.keys()) {
      if (data.hasOwnProperty(key)) {
        m.set(key, {
          ...m.get(key),
          status: !hasFailed ? data[key] : 'FAILED'
        });
        if (data[key] === 'FAILED' && !hasFailed) {
          hasFailed = true;
        }
      }
    }
    return Array.from(m.values());
  }

  processManualStatusStep(data: any): any[] {
    if (data['nodeGroupOperation']) {
      return [
        {
          details: 'NodeGroup Operation',
          status: data['nodeGroupOperationProgress']
        }
      ];
    }
    const m = new Map([
      ['initiatingPlatformSetup', { details: 'Initiating Platform Setup', status: 'PENDING', step: 'initiatingPlatformSetup' }],
      ['preparingConfigurations', { details: 'Service Configuration', status: 'PENDING', step: 'preparingConfigurations' }],
      ['preparingServices', { details: 'Setup Services', status: 'PENDING', step: 'preparingServices' }],
      ['platformSetupCompleted', { details: 'Platform Environment Setup', status: 'PENDING', step: 'platformSetupCompleted' }],
      ['setupFinishUp', { details: 'Finishing Up', status: 'PENDING', step: 'setupFinishUp' }]
    ]);
    let hasFailed = false;
    for (let key of m.keys()) {
      if (data.hasOwnProperty(key)) {
        m.set(key, {
          ...m.get(key),
          status: !hasFailed ? data[key] : 'FAILED'
        });
        if (data[key] === 'FAILED' && !hasFailed) {
          hasFailed = true;
        }
      }
    }
    return Array.from(m.values());
  }

  getClusterDetails(): void {
    this.clusterService
      .mcGetCluster(this.clusterService.clusterNameSnapshot)
      .pipe(takeWhile(() => this.isAlive))
      .subscribe(
        _cluster => {
          this.clusterDetails = _cluster;
          this.clusterType = _cluster.clusterType;
          if (this.clusterType === 'AUTOMATED') {
            this.getClusterCreationStepStatus();
          } else {
            this.getManualCreationStepStatus();
          }
        },
        err => {
          console.log(err);
          this.toastr.error(err.message, 'ERROR');
        }
      );
  }
}
