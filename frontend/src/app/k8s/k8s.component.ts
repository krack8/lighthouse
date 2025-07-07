import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute, NavigationEnd, Params, PRIMARY_OUTLET, Router, UrlSegment } from '@angular/router';
import icKeyboardBackspace from '@iconify/icons-ic/keyboard-backspace';
import { filter, map, takeWhile } from 'rxjs/operators';
import { K8sService } from './k8s.service';
import icAdd from '@iconify/icons-ic/twotone-add';
import { K8sUpdateComponent } from './k8s-update/k8s-update.component';
import { MatDialog } from '@angular/material/dialog';
import { k8sRoutesMap, k8sRoutesPermissionMap } from '@shared-ui/utils';
import icDown from '@iconify/icons-ic/twotone-keyboard-arrow-down';
import { ICluster } from '@cluster/cluster.model';
import icMoreVert from '@iconify/icons-ic/twotone-more-vert';
import icDelete from '@iconify/icons-ic/twotone-delete';
import icEdit from '@iconify/icons-ic/twotone-edit';
import { SecureDeleteDialogComponent } from '@shared-ui/ui';
import { ClusterService } from '@cluster/cluster.service';
import { ToastrService } from '@sdk-ui/ui';
import { ClusterRenameDialogComponent } from '@cluster/cluster-rename-dialog/cluster-rename-dialog.component';
import { SelectedClusterService } from '@core-ui/services/selected-cluster.service';
import { Subject, Subscription } from 'rxjs';

interface IBreadcrumb {
  label: string;
  params: Params;
  queryParams: Params;
  routePath: string;
  url: UrlSegment[];
}

@Component({
  selector: 'kc-k8s',
  templateUrl: './k8s.component.html',
  styleUrls: ['./k8s.component.scss']
})
export class K8sComponent implements OnInit, OnDestroy {
  private clusterInfoSubscription!: Subscription;
  icAdd = icAdd;
  icDown = icDown;
  icKeyboardBackspace = icKeyboardBackspace;
  icMoreVert = icMoreVert;
  icDelete = icDelete;
  icEdit = icEdit;
  isAlive: boolean = true;
  isLoading!: boolean;
  clusterId!: string;
  clusterData: any;
  namespaceQueryParam!: string;
  resourceRouteMap = k8sRoutesMap;
  resourcePermissionMap = k8sRoutesPermissionMap;

  breadcrumbs: IBreadcrumb[] = [];
  resourcesSearchTerm: any;

  constructor(
    private route: ActivatedRoute,
    private _k8sService: K8sService,
    private router: Router,
    public dialog: MatDialog,
    private _clusterService: ClusterService,
    private toastr: ToastrService,
    private _router: Router,
    private selectedClusterService: SelectedClusterService,
  ) {
    this.router.events
      .pipe(
        takeWhile(() => this.isAlive),
        filter(event => event instanceof NavigationEnd),
        map(() => {
          return this.route;
        })
      )
      .subscribe(route => {
        if (this.namespaceQueryParam) {
          this.namespaceQueryParam = '';
        }
        this.breadcrumbs = this.getBreadcrumbs(route);
      });
  }

  ngOnInit(): void {
    this._k8sService.changeClusterId(this.route.snapshot.params['clusterId']);
    this.getClusterId();
    this.clusterInfoSubscription = this._k8sService.clusterInfo$.subscribe((cluster) => {
      if (cluster ) {
          this.clusterData = this._k8sService.clusterInfoSnapshot;
      }
    });
  }

  ngOnDestroy(): void {
    this.isAlive = false;
    this._k8sService.changeClusterId('');
    this.selectedClusterService.setSelectedClusterId(this.selectedClusterService.defaultClusterId); 
    if (this.clusterInfoSubscription) {
      this.clusterInfoSubscription.unsubscribe();
    }
  }

  // GET CLUSTER ID FROM Query Params or API
  getClusterId(): void {
    this.isLoading = true;
    this.route.data.subscribe(({ cluster }: { cluster: ICluster }) => {
      this.clusterId = cluster.id;
      this._k8sService.changeClusterInfo(cluster);
      this.isLoading = false;
    });
  }

  // Create BreadCrumb
  private getBreadcrumbs(route: ActivatedRoute, url: string = '', breadcrumbs: IBreadcrumb[] = []): IBreadcrumb[] {
    const ROUTE_DATA_BREADCRUMB: string = 'breadcrumb';

    //get the child routes
    let children: ActivatedRoute[] = route.children;

    //return if there are no more children
    if (children.length === 0) {
      return breadcrumbs;
    }

    //iterate over each children
    for (let child of children) {
      //verify primary route
      if (child.outlet !== PRIMARY_OUTLET) {
        continue;
      }

      //verify the custom data property "breadcrumb" is specified on the route
      if (!child.snapshot.data.hasOwnProperty(ROUTE_DATA_BREADCRUMB)) {
        return this.getBreadcrumbs(child, url, breadcrumbs);
      }

      // get the route's URL segment
      let routeURL: string = child.snapshot.url.map(segment => segment.path).join('/');

      //append route URL to URL
      url += `/${routeURL}`;

      //add breadcrumb
      let breadcrumb: IBreadcrumb = {
        label: child.snapshot.data[ROUTE_DATA_BREADCRUMB],
        params: child.snapshot.params,
        queryParams: child.snapshot?.queryParams,
        url: child.snapshot.url,
        routePath: url
      };
      if (child.snapshot.queryParams?.namespace) {
        this.namespaceQueryParam = child.snapshot.queryParams?.namespace;
      }
      breadcrumbs.push(breadcrumb);
      //recursive
      return this.getBreadcrumbs(child, url, breadcrumbs);
    }
  }

  onCreate(): void {
    const dialog = this.dialog.open(K8sUpdateComponent, {
      minHeight: '300px',
      width: '900px',
      disableClose: true
    });
    dialog.componentInstance.applyManifestFor = 'all';
  }

  navigateToResources(resource: string) {
    this.resourcesSearchTerm = '';
    const urlSegments = window.location.href.split('/');
    const path = '/' + urlSegments[3] + '/' + urlSegments[4] + '/' + urlSegments[5] + '/namespaces/' + resource;
    this.router.navigate([path], {
      queryParams: { namespace: 'default' }
    });
  }

  deleteCluster(): void {
    const cluster = this._k8sService.clusterInfoSnapshot;
    const deleteDialog = this.dialog.open(SecureDeleteDialogComponent, {
      width: '600px',
      minHeight: '350px',
      data: {
        module: 'CLUSTER',
        route: '/clusters',
        id: cluster?.id,
        name: cluster?.name,
        method: this._clusterService.deleteCluster(cluster?.id),
        clusterStatus: cluster.cluster_status
      },
    });
    deleteDialog.afterClosed().subscribe((status: string) => {
      console.log('status', status);
      if (status === 'success') {
        this.toastr.success('Cluster deleted successfully');
        this._router.navigate(['/clusters']);
      }
    });
  }

  openClusterRenameDialog(): void {
    const cluster = this._k8sService.clusterInfoSnapshot;
    const dialogRef = this.dialog.open(ClusterRenameDialogComponent, {
      width: '600px',
      minHeight: '270px',
      data: cluster
    });
    dialogRef.afterClosed().subscribe((status: string) => {
      if (status === 'success') {
        window.location.reload();
      }
    });
  }
}
