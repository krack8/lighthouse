import { Component, Inject, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { distinctUntilChanged, map, share, takeWhile, throttleTime } from 'rxjs/operators';
import { K8sNamespacesService } from '../k8s-namespaces.service';
import icSearch from '@iconify/icons-ic/search';
import icInfo from '@iconify/icons-ic/twotone-info';
import icAdd from '@iconify/icons-ic/twotone-add';
import icEdit from '@iconify/icons-ic/twotone-edit';
import icRefresh from '@iconify/icons-ic/twotone-refresh';
import icCross from '@iconify/icons-ic/twotone-cancel';
import icLabel from '@iconify/icons-ic/twotone-label';
import icDelete from '@iconify/icons-ic/twotone-delete';
import icMoreHoriz from '@iconify/icons-ic/twotone-more-horiz';
import { ToolbarService } from '@sdk-ui/services/toolbar.service';
import icKeyboardBackspace from '@iconify/icons-ic/keyboard-backspace';
import { MatDialog } from '@angular/material/dialog';
import { K8sNamespacesDetailsComponent } from '../k8s-namespaces-details/k8s-namespaces-details.component';
import { K8sUpdateComponent } from '@k8s/k8s-update/k8s-update.component';
import { ConfirmDialogStaticComponent } from '@shared-ui/ui';
import { ToastrService } from '@sdk-ui/ui';
import { PermissionService } from '@core-ui/services/permission.service';
import { RequesterService } from '@core-ui/services/requester.service';
import { DOCUMENT } from '@angular/common';
import { fromEvent } from 'rxjs';

@Component({
  selector: 'kc-k8s-namespaces-list',
  templateUrl: './k8s-namespaces-list.component.html',
  styleUrls: ['./k8s-namespaces-list.component.scss']
})
export class K8sNamespacesListComponent implements OnInit, OnDestroy {
  icKeyboardBackspace = icKeyboardBackspace;
  icSearch = icSearch;
  icInfo = icInfo;
  icAdd = icAdd;
  icCross = icCross;
  icLabel = icLabel;
  icEdit = icEdit;
  icDelete = icDelete;
  icMoreHoriz = icMoreHoriz;
  icRefresh = icRefresh;
  isAlive: boolean = true;
  namespaceList: any[] = [];
  selectedNamespace: string = '';
  isLoading!: boolean;
  searchTerm = '';
  searchBy = 'name';
  title: any = 'Namespaces';
  userPermissions: any[] = [];
  userType: string;
  loadingSpanner: boolean = true;
  data: any;
  resourceToken: string = '';
  loadMoreData: boolean = false;
  remaining;
  tokenReceiveTime: Date;

  constructor(
    private namespaceService: K8sNamespacesService,
    private router: Router,
    private route: ActivatedRoute,
    private toolbarService: ToolbarService,
    public dialog: MatDialog,
    private toastr: ToastrService,
    private permissionService: PermissionService,
    private requesterService: RequesterService,
    @Inject(DOCUMENT) public document: any
  ) {}

  ngOnInit(): void {
    this.userType = this.requesterService.get().userInfo?.user_type;
    this.toolbarService.changeData({ title: this.title });
    this.namespaceService.selectedNamespace$.pipe(takeWhile(() => this.isAlive)).subscribe((namespace: string) => {
      this.selectedNamespace = namespace;
    });
    this.getNamespaces();
    this.userPermissions = this.permissionService.userPermissionsSnapshot;
  }

  ngOnDestroy(): void {
    this.isAlive = false;
  }

  reloadList() {
    this.getNamespaces();
  }

  getNamespaces(qp?: any): void {
    this.loadingSpanner = true;
    this.isLoading = true;
    this.namespaceService
      .getNamespaces(qp)
      .pipe(takeWhile(() => this.isAlive))
      .subscribe({
        next: data => {
          if (data?.status === 'success') {
            this.isLoading = false;
            this.namespaceList = data.data.Result || [];
            this.resourceToken = data.data.Resource;
            this.tokenReceiveTime = new Date();
            this.remaining = data.data.Remaining;
            this.loadingSpanner = false;
          } else {
            this.isLoading = false;
          }
        },
        error: err => {
          this.isLoading = false;
          this.toastr.error(err, 'Failed');
        }
      });
  }

  setAuthorizedRoute(): string {
    const set = new Set(this.userPermissions);
    if (set.has('VIEW_NAMESPACE_DEPLOYMENT') || this.userType === 'ADMIN' || this.userType === 'SUPER_ADMIN') {
      return 'deployments';
    } else if (set.has('VIEW_NAMESPACE_POD')) {
      return 'pods';
    } else if (set.has('VIEW_NAMESPACE_REPLICA_SET')) {
      return 'replica-sets';
    } else if (set.has('VIEW_NAMESPACE_STATEFUL_SET')) {
      return 'stateful-sets';
    } else if (set.has('VIEW_NAMESPACE_DAEMON_SET')) {
      return 'daemon-sets';
    } else if (set.has('VIEW_NAMESPACE_CONFIG_MAP')) {
      return 'config-maps';
    } else if (set.has('VIEW_NAMESPACE_SECRET')) {
      return 'secrets';
    } else if (set.has('VIEW_NAMESPACE_SERVICE')) {
      return 'services';
    } else if (set.has('VIEW_NAMESPACE_INGRESS')) {
      return 'ingresses';
    } else if (set.has('VIEW_NAMESPACE_PERSISTENT_VOLUME')) {
      return 'pvcs';
    } else if (set.has('VIEW_NAMESPACE_SERVICE_ACCOUNT')) {
      return 'service-accounts';
    } else if (set.has('VIEW_NAMESPACE_CERTIFICATE')) {
      return 'certificates';
    } else if (set.has('VIEW_NAMESPACE_ROLE')) {
      return 'role';
    } else if (set.has('VIEW_NAMESPACE_ROLE_BINDING')) {
      return 'role-binding';
    } else if (set.has('VIEW_NAMESPACE_JOB')) {
      return 'job';
    } else if (set.has('VIEW_NAMESPACE_CRON_JOB')) {
      return 'cron-job';
    } else if (set.has('VIEW_NAMESPACE_NETWORK_POLICY')) {
      return 'network-policy';
    } else if (set.has('VIEW_NAMESPACE_RESOURCE_QUOTA')) {
      return 'resource-quota';
    } else if (set.has('VIEW_NAMESPACE_GATEWAY')) {
      return 'gateway';
    } else if (set.has('VIEW_NAMESPACE_VIRTUAL_SERVICE')) {
      return 'virtual-service';
    }
  }

  onNamespaceClick(item): void {
    const routeTo = this.setAuthorizedRoute();
    this.namespaceService.changeSelectedNamespace(item);
    if (routeTo && routeTo !== '') {
      this.router.navigate([routeTo], {
        queryParams: {
          ...this.route.snapshot.queryParams,
          namespace: item.metadata.name
        },
        relativeTo: this.route
      });
    } else {
      this.toastr.error('You do not have access to this resource');
    }
  }

  openDetails(name: string) {
    this.dialog.open(K8sNamespacesDetailsComponent, {
      disableClose: true,
      width: '1000px',
      maxWidth: '1000px',
      maxHeight: '100vh',
      data: name
    });
  }

  onCreate(): void {
    const dialog = this.dialog.open(K8sUpdateComponent, {
      minHeight: '300px',
      width: '900px',
      disableClose: true
    });
    dialog.componentInstance.applyManifestFor = 'namespace';
    dialog.afterClosed().subscribe(res => {
      if (res) {
        if (res != null) {
          this.getNamespaces();
        }
      }
    });
  }

  onDelete(item: any): void {
    const dialogRef = this.dialog.open(ConfirmDialogStaticComponent, {
      disableClose: true,
      minWidth: '350px',
      data: {
        message: `Are you sure! want to delete ${item?.metadata?.name}?`,
        icon: '/assets/img/bin.svg'
      }
    });
    dialogRef.afterClosed().subscribe((bool: boolean) => {
      if (bool === true) {
        this.namespaceService.deleteNamespace(item?.metadata?.name).subscribe(
          res => {
            if (res.status === 'success') {
              this.toastr.success('Delete initiated');
              this.getNamespaces();
            }
          },
          err => {
            this.toastr.error('Failed: ', err.error.message);
          }
        );
      }
    });
  }

  onUpdate(item: any): void {
    const dialog = this.dialog.open(K8sUpdateComponent, {
      minHeight: '300px',
      width: '900px',
      disableClose: true
    });
    dialog.componentInstance.isEditMode = true;
    dialog.componentInstance.applyManifestFor = 'namespace';

    const metaTemp: { [key: string]: any } = {};
    metaTemp.name = item.metadata.name;
    metaTemp.namespace = item.metadata.namespace;
    metaTemp.uid = item.metadata.uid;
    if (item.metadata.selfLink) {
      metaTemp.selfLink = item.metadata.selfLink;
    }
    if (item.metadata.labels) {
      metaTemp.labels = item.metadata.labels;
    }
    if (item.metadata.annotations) {
      metaTemp.annotations = item.metadata.annotations;
    }

    const preInputData: { [key: string]: any } = {};
    preInputData.kind = item.kind;
    preInputData.apiVersion = item.apiVersion;
    preInputData.metadata = metaTemp;

    if (item.spec) {
      preInputData.spec = item.spec;
    }

    dialog.componentInstance.preInputData = preInputData;

    dialog.componentInstance.payload = {
      name: item.metadata.name,
      kind: item.kind,
      apiVersion: item.apiVersion,

      namespace: item?.metadata?.namespace
    };

    dialog.afterClosed().subscribe(res => {
      if (res) {
        this.getNamespaces();
      }
    });
  }

  onSearch() {
    if (this.searchBy === 'label') {
      const keyValuePairs = this.searchTerm.split(',');
      const jsonObject = {};
      keyValuePairs.forEach(pair => {
        if (pair.includes(':')) {
          const [key, value] = pair.split(':');
          jsonObject[key] = value;
        } else {
          this.toastr.error('Incorrect format for label search. Please use key:value format');
          return;
        }
      });
      if (JSON.stringify(jsonObject) === '{}') {
        return;
      }
      const jsonString = JSON.stringify(jsonObject);
      const qp = { labels: jsonString };
      this.getNamespaces(qp);
    }
    if (this.searchBy === 'name') {
      const qp = { q: this.searchTerm };
      this.getNamespaces(qp);
    }
  }

  clearSearch() {
    this.getNamespaces();
    this.searchTerm = '';
  }

  handleInputChange() {
    if (this.searchTerm.length === 0) {
      this.getNamespaces();
    }
  }

  ngAfterContentInit(): void {
    const content = this.document.querySelector('.sidenav-content');
    const scroll$ = fromEvent(content, 'scroll').pipe(
      takeWhile(() => this.isAlive),
      throttleTime(10), // only emit every 10 ms
      map((): boolean => {
        return content.offsetHeight + content.scrollTop + 80 >= content.scrollHeight;
      }),
      distinctUntilChanged(), // only emit when scrolling direction changed
      share() // share a single subscription to the underlying sequence in case of multiple subscribers
    );

    scroll$.subscribe((isBottom: boolean) => {
      if (isBottom && this.resourceToken.length !== 0 && !this.loadMoreData) {
        this.loadMore();
      }
    });
  }

  loadMore() {
    this.loadMoreData = true;
    let queryParam = {};
    const currentTime = new Date();
    const diff = (currentTime.getTime() - this.tokenReceiveTime.getTime()) / 60000;
    if (diff > 2) {
      queryParam = { limit: this.data.length + 10 };
    } else {
      queryParam = { continue: this.resourceToken };
    }
    if (this.searchTerm.length > 0) {
      queryParam['q'] = this.searchTerm;
    }
    this.namespaceService
      .getNamespaces(queryParam)
      .pipe(takeWhile(() => this.isAlive))
      .subscribe({
        next: data => {
          this.loadMoreData = false;
          this.remaining = data.data.Remaining;
          this.resourceToken = data.data.Resource || '';
          this.tokenReceiveTime = new Date();
          if (queryParam.hasOwnProperty('limit')) {
            this.namespaceList = data.data?.Result || [];
          } else {
            this.namespaceList = this.namespaceList.concat(data.data.Result) || [];
          }

          this.loadMoreData = false;
        },
        error: err => {
          this.toastr.error('Failed: ', err.error.message);
        }
      });
  }
}
