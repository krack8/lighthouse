import { DOCUMENT } from '@angular/common';
import { Component, Inject, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ActivatedRoute } from '@angular/router';
import icSearch from '@iconify/icons-ic/search';
import icAdd from '@iconify/icons-ic/twotone-add';
import icCross from '@iconify/icons-ic/twotone-cancel';
import icDelete from '@iconify/icons-ic/twotone-delete';
import icEdit from '@iconify/icons-ic/twotone-edit';
import icInfo from '@iconify/icons-ic/twotone-info';
import icLabel from '@iconify/icons-ic/twotone-label';
import icCircle from '@iconify/icons-ic/twotone-lens';
import icMoreVert from '@iconify/icons-ic/twotone-more-vert';
import icRefresh from '@iconify/icons-ic/twotone-refresh';
import { K8sUpdateComponent } from '@k8s/k8s-update/k8s-update.component';
import { fadeInRight400ms } from '@sdk-ui/animations/fade-in-right.animation';
import { ToolbarService } from '@sdk-ui/services/toolbar.service';
import { ToastrService } from '@sdk-ui/ui';
import { ConfirmDialogStaticComponent } from '@shared-ui/ui';
import { fromEvent } from 'rxjs';
import { distinctUntilChanged, map, share, takeWhile, throttleTime } from 'rxjs/operators';
import { K8sNamespacesService } from '../k8s-namespaces.service';

@Component({
  selector: 'kc-k8s-daemon-sets',
  templateUrl: './k8s-daemon-sets.component.html',
  styleUrls: ['./k8s-daemon-sets.component.scss'],
  animations: [fadeInRight400ms]
})
export class K8sDaemonSetsComponent implements OnInit {
  icCircle = icCircle;
  icSearch = icSearch;
  icMoreVert = icMoreVert;
  icInfo = icInfo;
  icAdd = icAdd;
  icDelete = icDelete;
  icRefresh = icRefresh;
  icLabel = icLabel;
  icEdit = icEdit;
  icCross = icCross;
  searchBy = 'name';
  title: any = 'DaemonSets';
  isAlive: boolean = true;
  loadingSpanner: boolean = true;
  statsLoaded: boolean = false;
  data: any;
  stats: any;
  searchTerm: string = '';
  queryParams: any;
  resourceToken: string = '';
  loadMoreData: boolean = false;
  remaining;
  tokenReceiveTime: Date;

  constructor(
    private namespaceService: K8sNamespacesService,
    private route: ActivatedRoute,
    private toolbarService: ToolbarService,
    private toastr: ToastrService,
    private dialog: MatDialog,
    @Inject(DOCUMENT) public document: any
  ) {}

  ngOnInit(): void {
    this.toolbarService.changeData({ title: this.title });
    this.route.queryParams.subscribe(params => {
      if (params) {
        this.getInstanceData();
        this.queryParams = this.route.snapshot.queryParams;
      }
    });
  }

  ngOnDestroy(): void {
    this.isAlive = false;
  }

  reloadList() {
    if (this.searchTerm !== '') this.onSearch();
    else this.getInstanceData();
  }

  getInstanceData(queryParam?: any): void {
    this.getDaemonSetStats(queryParam);
    this.loadingSpanner = true;
    this.namespaceService.getDaemonSets(queryParam).subscribe({
      next: res => {
        if (res?.status === 'error') {
          this.toastr.error(res?.message);
        }
        this.data = res?.data.Result || [];
        this.resourceToken = res.data.Resource;
        this.tokenReceiveTime = new Date();
        this.remaining = res.data.Remaining;
        this.loadingSpanner = false;
      },
      error: err => {
        this.loadingSpanner = false;
        this.toastr.error(err?.message);
      }
    });
  }

  getDaemonSetStats(queryParam?: any): void {
    this.statsLoaded = true;
    this.namespaceService.getDaemonSetsStats(queryParam).subscribe({
      next: res => {
        if (res?.status === 'error') {
          this.toastr.error(res?.message);
        }
        this.stats = res?.data;
        this.statsLoaded = false;
      },
      error: err => {
        this.statsLoaded = false;
        this.toastr.error(err?.message);
      }
    });
  }

  onCreate(): void {
    const dialog = this.dialog.open(K8sUpdateComponent, {
      minHeight: '300px',
      width: '900px',
      disableClose: true
    });
    dialog.componentInstance.applyManifestFor = 'daemonset';
    dialog.afterClosed().subscribe(res => {
      if (res) {
          this.getInstanceData();
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
        this.namespaceService.deleteNamespaceDaemonset(item?.metadata?.name).subscribe(
          res => {
            if (res.status === 'success') {
              this.toastr.success('Delete initiated');
              this.getInstanceData();
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
    dialog.componentInstance.applyManifestFor = 'daemonset';

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

    dialog.afterClosed().subscribe(res => {
      if (res) {
          this.getInstanceData();
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
      this.getInstanceData(qp);
    }
    if (this.searchBy === 'name') {
      const qp = { q: this.searchTerm };
      this.getInstanceData(qp);
    }
  }

  clearSearch() {
    this.getInstanceData();
    this.searchTerm = '';
  }
  handleInputChange() {
    if (this.searchTerm.length === 0) {
      this.getInstanceData();
    }
  }

  isInt(value: string): boolean {
    const parsedValue = parseInt(value, 10);
    return !isNaN(parsedValue) && String(parsedValue) === value;
  }

  isObject(value: any): boolean {
    return typeof value === 'object' && value !== null;
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
      .getDeploymentsV1(queryParam)
      .pipe(takeWhile(() => this.isAlive))
      .subscribe({
        next: data => {
          this.loadMoreData = false;
          this.remaining = data.data.Remaining;
          this.resourceToken = data.data.Resource || '';
          this.tokenReceiveTime = new Date();
          if (queryParam.hasOwnProperty('limit')) {
            this.data = data.data.Result || [];
          } else {
            this.data = this.data.concat(data.data.Result) || [];
          }

          this.loadMoreData = false;
        },
        error: err => {
          this.toastr.error('Failed: ', err.error.message);
        }
      });
  }

  percentageCalculator(used, total) {
    return Math.round((used / total) * 100);
  }
}
