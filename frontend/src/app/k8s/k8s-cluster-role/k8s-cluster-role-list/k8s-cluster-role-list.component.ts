import { Component, Inject, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { distinctUntilChanged, map, share, takeWhile, throttleTime } from 'rxjs/operators';
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
import { K8sClusterRoleService } from '../k8s-cluster-role.service';
import { K8sUpdateComponent } from '@k8s/k8s-update/k8s-update.component';
import { ToastrService } from '@sdk-ui/ui';
import { ConfirmDialogStaticComponent } from '@shared-ui/ui';
import { DOCUMENT } from '@angular/common';
import { fromEvent } from 'rxjs';

@Component({
  selector: 'kc-k8s-cluster-role-list',
  templateUrl: './k8s-cluster-role-list.component.html',
  styleUrls: ['./k8s-cluster-role-list.component.scss']
})
export class K8sClusterRoleListComponent implements OnInit, OnDestroy {
  icKeyboardBackspace = icKeyboardBackspace;
  icSearch = icSearch;
  icInfo = icInfo;
  icAdd = icAdd;
  icEdit = icEdit;
  icDelete = icDelete;
  icLabel = icLabel;
  icMoreHoriz = icMoreHoriz;
  icRefresh = icRefresh;
  icCross = icCross;
  searchBy = 'name';
  title: any = 'Cluster Role';
  isAlive: boolean = true;
  clusterRole: any[] = [];
  isLoading!: boolean;
  searchTerm = '';
  queryParams: any;
  resourceToken: string = '';
  loadMoreData: boolean = false;
  remaining;
  tokenReceiveTime: Date;

  constructor(
    @Inject(DOCUMENT) public document: any,
    private clusterRoleBindingService: K8sClusterRoleService,
    private router: Router,
    private route: ActivatedRoute,
    private toolbarService: ToolbarService,
    public dialog: MatDialog,
    private toastr: ToastrService
  ) {}

  ngOnInit(): void {
    this.toolbarService.changeData({ title: this.title });
    this.queryParams = this.route.snapshot.queryParams;
    this.getClusterRoleList();
  }

  ngOnDestroy(): void {
    this.isAlive = false;
  }

  reloadList() {
    this.getClusterRoleList();
  }

  getClusterRoleList(queryParam?: any): void {
    this.isLoading = true;
    this.clusterRoleBindingService
      .getClusterRole(queryParam)
      .pipe(takeWhile(() => this.isAlive))
      .subscribe({
        next: data => {
          if (data.status === 'success') {
            this.isLoading = false;
            this.resourceToken = data.data.Resource;
            this.tokenReceiveTime = new Date();
            this.remaining = data.data.Remaining;
            this.clusterRole = data.data.Result || [];
          } else {
            this.isLoading = false;
          }
        },
        error: err => {
          this.isLoading = false;
          this.toastr.error('Failed: ', err.error.message);
        }
      });
  }

  onPvDetailsClick(item): void {
    // this.namespaceService.changeSelectedNamespace(item);
    this.router.navigate(['cluster-role-details'], {
      queryParams: {
        ...this.route.snapshot.queryParams,
        name: item.metadata.name
      },
      relativeTo: this.route
    });
  }

  onCreate(): void {
    const dialog = this.dialog.open(K8sUpdateComponent, {
      minHeight: '300px',
      width: '900px',
      disableClose: true
    });
    dialog.componentInstance.applyManifestFor = 'cluster-role';
    dialog.afterClosed().subscribe(res => {
      if (res) {
        if (res != null) {
          this.getClusterRoleList();
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
        this.clusterRoleBindingService.deleteClusterRole(item?.metadata?.name).subscribe(
          res => {
            if (res.status === 'success') {
              this.toastr.success('Delete initiated');
              setTimeout(() => {
                this.getClusterRoleList();
              }, 6000);
              this.getClusterRoleList();
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
    dialog.componentInstance.applyManifestFor = 'cluster-role';

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
    if (item.rules) {
      preInputData.rules = item.rules;
    }
    if (item.aggregationRule) {
      preInputData.aggregationRule = item.aggregationRule;
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
        if (res != null) {
          this.getClusterRoleList();
        }
      }
    });
  }

  extractCapitalLetters(str: string) {
    return str.replace(/[^A-Z]+/g, '');
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
      this.getClusterRoleList(qp);
    }
    if (this.searchBy === 'name') {
      const qp = { q: this.searchTerm };
      this.getClusterRoleList(qp);
    }
  }

  clearSearch() {
    this.getClusterRoleList();
    this.searchTerm = '';
  }
  handleInputChange() {
    if (this.searchTerm.length === 0) {
      this.getClusterRoleList();
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
      queryParam = { limit: this.clusterRole.length + 10 };
    } else {
      queryParam = { continue: this.resourceToken };
    }
    if (this.searchTerm.length > 0) {
      queryParam['q'] = this.searchTerm;
    }
    this.clusterRoleBindingService
      .getClusterRole(queryParam)
      .pipe(takeWhile(() => this.isAlive))
      .subscribe({
        next: data => {
          this.loadMoreData = false;
          this.remaining = data.data.Remaining;
          this.resourceToken = data.data.Resource || '';
          this.tokenReceiveTime = new Date();
          if (queryParam.hasOwnProperty('limit')) {
            this.clusterRole = data.data.Result || [];
          } else {
            this.clusterRole = this.clusterRole.concat(data.data.Result) || [];
          }
          this.loadMoreData = false;
        },
        error: err => {
          this.toastr.error('Failed: ', err.error.message);
        }
      });
  }
}
