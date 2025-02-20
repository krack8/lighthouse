import { Component, OnInit } from '@angular/core';
import icAddCircleOutline from '@iconify/icons-ic/twotone-add';
import { ToastrService } from '@sdk-ui/ui';
import { ClusterService } from '../cluster.service';
import { ToolbarService } from '@sdk-ui/services/toolbar.service';
import { ICluster } from '@cluster/cluster.model';
import { Router } from '@angular/router';
import { MatDialog } from '@angular/material/dialog';
import icCircle from '@iconify/icons-ic/twotone-lens';
import { PermissionService, RequesterService } from '@core-ui/services';

enum View {
  GRID = 'grid',
  LIST = 'list'
}

@Component({
  selector: 'kc-cluster-list',
  templateUrl: './cluster-list.component.html',
  styleUrls: ['./cluster-list.component.scss']
})
export class ClusterListComponent implements OnInit {
  icAddCircleOutline = icAddCircleOutline;
  icCircle = icCircle;
  clusterList: any = [];
  dataLoading!: boolean;
  serverError: boolean = false;
  permissions: any[] = [];
  k8sRoute: string;
  k8sAuthority: string;
  userType: string;

  viewStyle: View = View.LIST;

  constructor(
    private clusterService: ClusterService,
    private toastrService: ToastrService,
    private toolbarService: ToolbarService,
    private router: Router,
    private _dialog: MatDialog,
    private permissionSvc: PermissionService,
    private requesterService: RequesterService
  ) {}

  ngOnInit(): void {
    this.userType = this.requesterService.get().userInfo?.user_type;
    this.permissions = this.permissionSvc.userPermissionsSnapshot;
    this.k8sRoute = this.setK8sRoute();
    this.toolbarService.changeData({ title: 'Cluster' });
    this.getCluster();
  }

  getCluster(): void {
    this.dataLoading = true;
    this.clusterService.getClusters().subscribe({
      next: data => {
        this.clusterList = data || [];
        this.dataLoading = false;
        this.serverError = false;
      },
      error: err => {
        this.dataLoading = false;
        this.serverError = true;
        this.toastrService.error(err.message, 'ERROR');
      }
    });
  }

  changeView(): void {
    if (this.viewStyle === View.LIST) {
      this.viewStyle = View.GRID;
    } else {
      this.viewStyle = View.LIST;
    }
  }

  routeToDetails(cluster?: ICluster): void {
    if (cluster && cluster.is_active) {
      this.router.navigate(['/clusters', cluster?.id, 'k8s', this.k8sRoute]);
      return;
    }

    import('../cluster-form/cluster-form.component').then(m => {
    const dialog = this._dialog.open(m.ClusterFormComponent, {
        width: '800px',
        data: cluster
      });
      dialog.afterClosed().subscribe((result) => {
        if (result) this.getCluster();
      })
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
    } else {
      return 'node-list';
    }
  }
}
