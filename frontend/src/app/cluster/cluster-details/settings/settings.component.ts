import { Component, OnDestroy, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ToastrService } from '@sdk-ui/ui';
import { Observable, of } from 'rxjs';
import { takeWhile, catchError, switchMap, tap } from 'rxjs/operators';
import { ClusterService } from '../../cluster.service';
import { fadeInRight400ms } from '@sdk-ui/animations/fade-in-right.animation';
import { ConfirmDeleteDialogueComponent } from './confirm-delete-dialogue/confirm-delete-dialogue.component';
import moment from 'moment';
import icInfo from '@iconify/icons-ic/twotone-info';
import { ClusterReleaseNoteDialogComponent } from '../overview/cluster-release-dialogs/cluster-release-note-dialog/cluster-release-note-dialog.component';

@Component({
  selector: 'kc-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.scss'],
  animations: [fadeInRight400ms]
})
export class SettingsComponent implements OnInit, OnDestroy {
  isAlive: boolean = true;
  isConfigOpened: boolean = false;
  clusterId: string = '';
  clusterName: string = '';
  clusterDetails: any;
  kubeClusterDetails: any = {};
  dataLoading: boolean = true;
  isForceDelete: boolean = false;
  icInfo = icInfo;

  //cluster updgrade settings
  serverError: boolean = false;

  errMessage: string;
  clusterReleaseInfo: any;
  updateWarnigClass: string;
  releaseUpdateInProgress: boolean = false;
  isUpgradeOpen: boolean = false;
  selectedStrategy: string = '';
  selectedTime = [];
  localFromTime: string;
  localToTime: string;

  showClusterUpdateForm: boolean = false;

  constructor(
    private toastr: ToastrService,
    private dialog: MatDialog,
    private clusterService: ClusterService
  ) {}

  ngOnInit(): void {
    this.getClusterDetails();
  }

  ngOnDestroy(): void {
    this.isAlive = false;
  }

  getClusterDetails(): void {
    this.clusterService
      .mcGetCluster(this.clusterService.clusterNameSnapshot)
      .pipe(
        takeWhile(() => this.isAlive),
        switchMap(_cluster => {
          this.clusterDetails = _cluster;
          this.clusterId = _cluster.clusterId;
          // Is Force Delete
          const currentTime = new Date().getTime();
          const updateTimeForExpire = new Date(_cluster?.updatedTime).getTime() + 300000; // UpdateTime + 20mins
          const createTimeForLonger = new Date(_cluster?.createdTime).getTime() + 3000000; // CreateTime + 50mins
          if (
            (_cluster.currentState === 'DELETING' && updateTimeForExpire < currentTime) ||
            (_cluster.currentState === 'CREATING' && createTimeForLonger < currentTime)
          ) {
            this.isForceDelete = true;
          }

          // Kube Configuration
          return this.clusterService.getClusterById(_cluster.clusterId).pipe(
            tap((res: any) => {
              if (res.status === 'success') {
                this.kubeClusterDetails = res.data;
                this.selectedStrategy = this.kubeClusterDetails?.clusterUpdateStrategy?.strategy;
                this.selectedTime[0] = this.kubeClusterDetails?.clusterUpdateStrategy?.fromTime;
                this.selectedTime[1] = this.kubeClusterDetails?.clusterUpdateStrategy?.toTime;

                if (this.kubeClusterDetails?.clusterReleaseConfig?.releaseStatus === 'FAILED') {
                  this.updateWarnigClass = 'border-danger';
                } else if (this.kubeClusterDetails?.clusterReleaseConfig?.releaseStatus === 'SUCCESS') {
                  this.updateWarnigClass = 'border-success';
                } else {
                  this.updateWarnigClass = 'border-warning';
                }
              }
            }),
            catchError(err => of({}))
          );
        })
      )
      .subscribe(
        _ => {
          this.dataLoading = false;
        },
        error => {
          this.dataLoading = false;
          this.toastr.error(error.message, 'ERROR');
        }
      );
  }

  delete() {
    let _successRedirectRoute: string = `/clusters/${this.clusterDetails.clusterName}/logs`;
    let _method: Observable<any> = this.clusterService.mcDeleteCluster(this.clusterId);
    if (this.isForceDelete || this.clusterDetails.clusterType === 'MANUAL') {
      _method = this.clusterService.mcForceDeleteCluster(this.clusterId);
      _successRedirectRoute = '/clusters';
    }
    const dialogRef = this.dialog.open(ConfirmDeleteDialogueComponent, {
      width: '600px',
      minHeight: '350px',
      data: {
        route: _successRedirectRoute,
        name: this.clusterDetails.clusterName,
        forceDelete: this.isForceDelete,
        method: _method,
        clusterId: this.clusterId,
        successTitle: this.isForceDelete ? 'Success' : 'Deleting'
      }
    });
    dialogRef.afterClosed().subscribe(status => {
      if (status === 'OK') {
        this.clusterService.changeCurrentState('DELETING');
      }
    });
  }

  toggleClusterUpdateForm(): void {
    this.showClusterUpdateForm = !this.showClusterUpdateForm;
  }

  toggleClusterUpgradeView(): void {
    this.isUpgradeOpen = !this.isUpgradeOpen;
  }

  selectType(param: string): void {
    this.selectedStrategy = param;
  }

  updateClusterUpgradeStrategy(): void {
    let payload = {};
    if (this.selectedStrategy === 'AUTOMATED') {
      payload = {
        strategy: this.selectedStrategy,
        fromTime: this.selectedTime[0],
        toTime: this.selectedTime[1],
        updateTime: this.getUtcHour(this.selectedTime[0]) + '-' + this.getUtcHour(this.selectedTime[1])
      };
    } else {
      payload = {
        strategy: this.selectedStrategy
      };
    }
    this.clusterService.updateClusterUpgradeStrategy(this.clusterId, payload).subscribe(
      res => {
        if (res) {
          this.toastr.success('Cluster Upgrade Strategy Updated Successfully');
          this.isUpgradeOpen = false;
        }
      },
      error => {
        this.toastr.error(error.message, 'ERROR');
      }
    );
  }

  onDatePicker(date, type: string) {
    // no use of this function
    const utcTime = moment(date.value).utc().format();
    const hour = new Date(utcTime).getUTCHours();
  }

  getUtcHour(time: string): number {
    return new Date(moment(time).utc().format()).getUTCHours(); // converts local time to utc and return hour only
  }

  upgradeClusterDialog(): void {
    const dialogRef = this.dialog.open(ClusterReleaseNoteDialogComponent, {
      width: '70%',
      minHeight: '40%',
      maxHeight: '80vh',
      data: this.kubeClusterDetails?.clusterReleaseConfig
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.releaseUpdateInProgress = true;
        this.clusterService.activateClusterUpgrade(this.kubeClusterDetails.id).subscribe(res => {
          if (res['status'] === 'success') {
            this.getClusterDetails();
            this.releaseUpdateInProgress = false;
          } else this.releaseUpdateInProgress = false;
        });
      }
    });
  }
}
