import { Component, OnDestroy, OnInit } from '@angular/core';
import { ClusterService } from '../../cluster.service';
import { catchError, switchMap, tap } from 'rxjs/operators';
import { of } from 'rxjs';
import { ClusterReleaseNoteDialogComponent } from './cluster-release-dialogs/cluster-release-note-dialog/cluster-release-note-dialog.component';
import { MatDialog } from '@angular/material/dialog';
import { ConfirmDialogComponent } from '@shared-ui/ui';

@Component({
  selector: 'kc-overview',
  templateUrl: './overview.component.html',
  styleUrls: ['./overview.component.scss']
})
export class OverviewComponent implements OnInit, OnDestroy {
  isAlive: boolean = true;
  serverError: boolean = false;

  errMessage: string;

  isLoaded: boolean = false;
  clusterDetails: any;
  kubeClusterDetails: any = {};
  clusterReleaseInfo: any;
  newReleaseAvailable: boolean = false;
  releaseUpdateInProgress: boolean = false;

  constructor(
    private clusterService: ClusterService,
    public dialog: MatDialog
  ) {}

  ngOnInit(): void {
    this.getClusterDetails();
    // this.getClusterReleaseInfo();
  }

  ngOnDestroy(): void {
    this.isAlive = false;
  }

  getProcessedHelmCommand(): string {
    if (this.clusterDetails && this.clusterDetails.helmCommand)
      return this.clusterDetails.helmCommand
        .split('\n')
        .map(line => line.trim().replace(/(\s+)--set/g, '--set'))
        .join('\n');
    return 'N/A';
  }

  getClusterDetails(): void {
    this.clusterService
      .mcGetCluster(this.clusterService.clusterNameSnapshot)
      .pipe(
        switchMap(_cluster => {
          this.clusterDetails = _cluster;
          return this.clusterService.getClusterById(_cluster.clusterId).pipe(
            tap((res: any) => {
              if (res.status === 'success') {
                this.kubeClusterDetails = res.data;
                console.log('kubeClusterDetails', this.kubeClusterDetails);
                if (this.kubeClusterDetails?.clusterReleaseConfig?.newReleaseId) {
                  this.newReleaseAvailable = true;
                }
                if (
                  this.kubeClusterDetails?.clusterReleaseConfig?.newReleaseId &&
                  this.kubeClusterDetails?.clusterUpdateStrategy?.active === true
                ) {
                  this.releaseUpdateInProgress = true;
                }
              }
            }),
            catchError(err => of({}))
          );
        })
      )
      .subscribe(
        _cluster => {
          this.serverError = false;
          this.isLoaded = true;
        },
        error => {
          this.serverError = true;
          this.isLoaded = true;
          this.errMessage = error.message;
        }
      );
  }

  getClusterReleaseInfo(): void {
    this.clusterService
      .mcGetCluster(this.clusterService.clusterNameSnapshot)
      .pipe(
        switchMap(_cluster => {
          console.log('_cluster', _cluster);
          return this.clusterService.getClusterReleaseInfoById().pipe(
            tap((res: any) => {
              if (res.status === 'success') {
                this.clusterReleaseInfo = res?.data?.find(cluster => cluster.clusterId === _cluster.clusterId);
                console.log('this.clusterReleaseInfo', this.clusterReleaseInfo);
                if (this.clusterReleaseInfo?.clusterReleaseConfig?.newReleaseId) {
                  this.newReleaseAvailable = true;
                }
                if (
                  this.clusterReleaseInfo?.clusterReleaseConfig?.newReleaseId &&
                  this.clusterReleaseInfo?.clusterUpdateStrategy?.active === true
                ) {
                  this.releaseUpdateInProgress = true;
                }
              }
            }),
            catchError(err => of({}))
          );
        })
      )
      .subscribe(
        _cluster => {
          this.serverError = false;
          this.isLoaded = true;
        },
        error => {
          this.serverError = true;
          this.isLoaded = false;
          this.errMessage = error.message;
        }
      );
  }

  getFormattedKey(key: string): string {
    return key
      .replace(/([a-z])([A-Z])/g, '$1 $2') // Add a space before capital letters
      .replace(/\b\w/g, char => char.toUpperCase()); // Capitalize the first letter of each word
  }

  upgradeCluster(): void {
    const dialogRef = this.dialog.open(ClusterReleaseNoteDialogComponent, {
      width: '1050px',
      //height: '700px',
      maxHeight: '770px',
      data: this.kubeClusterDetails?.clusterReleaseConfig
    });
    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.releaseUpdateInProgress = true;
        this.clusterService.activateClusterUpgrade(this.kubeClusterDetails.id).subscribe(res => {
          if (res['status'] === 'success') {
            this.getClusterDetails();
          } else {
            this.releaseUpdateInProgress = false;
          }
        });
      }
    });
  }
}
