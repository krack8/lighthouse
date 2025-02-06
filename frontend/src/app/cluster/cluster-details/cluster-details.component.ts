import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { distinctUntilChanged, filter, map, switchMap, takeWhile } from 'rxjs/operators';
import { ClusterService } from '../cluster.service';

import icKeyboardBackspace from '@iconify/icons-ic/keyboard-backspace';
import { ToastrService } from '@sdk-ui/ui';
import { ToolbarService } from '@sdk-ui/services/toolbar.service';
import { of } from 'rxjs';
import { PermissionService } from '@core-ui/services/permission.service';
import { RequesterService } from '@core-ui/services/requester.service';

@Component({
  selector: 'kc-cluster-details',
  templateUrl: './cluster-details.component.html',
  styleUrls: ['./cluster-details.component.scss']
})
export class ClusterDetailsComponent implements OnInit, OnDestroy {
  icKeyboardBackspace = icKeyboardBackspace;
  isAlive: boolean = true;
  isLoading: boolean = true;

  clusterStepStatus: any = '';
  clusterType: any = '';
  clusterStatus: string = '';
  clusterData: any;
  clusterName = '';

  permissions: any[] = [];
  k8sRoute: string;
  k8sAuthority: string;
  userType: string;

  constructor(
    private router: Router,
    private route: ActivatedRoute,
    private clusterService: ClusterService,
    private toastr: ToastrService,
    private toolbarService: ToolbarService,
    private permissionService: PermissionService,
    private requesterService: RequesterService
  ) {}

  ngOnInit(): void {
    this.userType = this.requesterService.get().userInfo?.user_type;
    this.route.params.pipe(takeWhile(() => this.isAlive)).subscribe((params: any) => {
      this.clusterName = params['clusterName'];
      this.clusterService.changeClusterName(params['clusterName']);
    });
    this.clusterService.creationStepStatus$.pipe(takeWhile(() => this.isAlive)).subscribe((status: any) => {
      this.clusterStepStatus = status.clusterFinishUp;
    });
    this.clusterService.currentState$
      .pipe(
        takeWhile(() => this.isAlive),
        distinctUntilChanged()
      )
      .subscribe((status: string) => {
        this.clusterStatus = status;
      });
    this.getCluster();
    this.permissions = this.permissionService.userPermissionsSnapshot;
    this.k8sRoute = this.setK8sRoute();
  }

  ngOnDestroy() {
    this.isAlive = false;
    this.clusterService.changeClusterName('');
    this.clusterService.changeCreationStepStatus({});
    this.clusterService.changeCurrentState('ACTIVE');
  }

  getCluster(): void {
    this.route.data
      .pipe(
        takeWhile(() => this.isAlive),
        map(data => {
          return data['cluster'];
        }),
        filter(clusterData => !!clusterData),
        switchMap((data: any) => {
          this.clusterType = data.clusterType;
          this.toolbarService.changeData({ title: 'Cluster' });
          if (data.clusterType === 'AUTOMATED') {
            this.clusterService.changeCurrentState(data.currentState);
            return this.clusterService.mcGetClustersCreationStepStatus(this.clusterName).pipe(map(res => [data, res?.data]));
          }
          return of([data, null]);
        })
      )
      .subscribe({
        next: ([cluster, stepData]) => {
          if (cluster !== null) {
            this.clusterData = cluster;
            this.clusterService.changeCurrentState(cluster?.currentState);
            if (cluster.clusterType === 'AUTOMATED' && (cluster?.currentState !== 'ACTIVE' || stepData?.clusterFinishUp !== 'SUCCESS')) {
              this.router.navigate(['logs'], { relativeTo: this.route });
            }
          }
          this.isLoading = false;
        },
        error: err => {
          // For Creation Step Api Error
          this.isLoading = false;
          this.toastr.error(err?.message || 'Something is wrong creation step!!!');
        }
      });
  }

  setK8sRoute(): string {
    const set = new Set(this.permissions);
    if (set.has('VIEW_K8S_NODES') || this.userType === 'ADMIN' || this.userType === 'SUPER_ADMIN') {
      this.k8sAuthority = 'VIEW_K8S_NODES';
      return 'node-list';
    } else if (set.has('VIEW_K8S_NAMESPACE')) {
      this.k8sAuthority = 'VIEW_K8S_NAMESPACE';
      return 'namespaces';
    } else if (set.has('VIEW_K8S_PERSISTENT_VOLUME')) {
      this.k8sAuthority = 'VIEW_K8S_PERSISTENT_VOLUME';
      return 'persistent-volume';
    } else if (set.has('VIEW_K8S_CLUSTER_ROLE')) {
      this.k8sAuthority = 'VIEW_K8S_CLUSTER_ROLE';
      return 'cluster-role';
    } else if (set.has('VIEW_K8S_CLUSTER_ROLE_BINDING')) {
      this.k8sAuthority = 'VIEW_K8S_CLUSTER_ROLE_BINDING';
      return 'cluster-role-binding';
    } else if (set.has('VIEW_K8S_STORAGE_CLASS')) {
      this.k8sAuthority = 'VIEW_K8S_STORAGE_CLASS';
      return 'storage-class';
    } else if (set.has('VIEW_K8S_CUSTOM_RESOURCE_DEFINATION')) {
      this.k8sAuthority = 'VIEW_K8S_CUSTOM_RESOURCE_DEFINATION';
      return 'custom-resources-defination';
    }
  }
}
