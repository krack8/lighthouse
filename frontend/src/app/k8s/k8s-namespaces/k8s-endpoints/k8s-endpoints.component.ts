import { Component, Inject, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import icSearch from '@iconify/icons-ic/search';
import { ToastrService } from '@sdk-ui/ui';
import { ToolbarService } from '@sdk-ui/services/toolbar.service';
import { K8sNamespacesService } from '../k8s-namespaces.service';
import icEdit from '@iconify/icons-ic/twotone-edit';
import icDelete from '@iconify/icons-ic/twotone-delete';
import icLabel from '@iconify/icons-ic/twotone-label';
import icRefresh from '@iconify/icons-ic/twotone-refresh';
import icUpgrade from '@iconify/icons-ic/twotone-file-upload';
import icMoreVert from '@iconify/icons-ic/twotone-more-vert';
import icCross from '@iconify/icons-ic/twotone-cancel';
import icInfo from '@iconify/icons-ic/twotone-info';
import icAdd from '@iconify/icons-ic/twotone-add';
import { MatDialog } from '@angular/material/dialog';
import { K8sUpdateComponent } from '@k8s/k8s-update/k8s-update.component';
import { ConfirmDialogStaticComponent } from '@shared-ui/ui';
import { DOCUMENT } from '@angular/common';
import { fromEvent } from 'rxjs';
import { distinctUntilChanged, map, share, takeWhile, throttleTime } from 'rxjs/operators';

@Component({
  selector: 'kc-k8s-endpoints',
  templateUrl: './k8s-endpoints.component.html',
  styleUrls: ['./k8s-endpoints.component.scss']
})
export class K8sEndpointsComponent implements OnInit {
  icSearch = icSearch;
  icInfo = icInfo;
  icDelete = icDelete;
  icEdit = icEdit;
  icAdd = icAdd;
  icLabel = icLabel;
  icUpgrade = icUpgrade;
  icMoreVert = icMoreVert;
  icCross = icCross;
  icRefresh = icRefresh;
  searchBy = 'name';
  title: any = 'Endpoints';
  isAlive: boolean = true;
  loadingSpanner: boolean = false;
  data: any[] = [];
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
    this.loadingSpanner = true;
    this.namespaceService.getEndpoints(queryParam).subscribe({
      next: res => {
        if (res?.status === 'error') {
          this.toastr.error(res?.message);
        }
        this.resourceToken = res.data.Resource;
        this.tokenReceiveTime = new Date();
        this.remaining = res.data.Remaining;
        this.data = res?.data.Result || [];
        this.loadingSpanner = false;
      },
      error: err => {
        this.loadingSpanner = false;
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
    dialog.componentInstance.applyManifestFor = 'endpoints';
    dialog.afterClosed().subscribe(res => {
      if (res) {
        if (res != null) {
          this.getInstanceData();
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
        this.namespaceService.deleteNamespaceEndpoint(item?.metadata?.name).subscribe(
          res => {
            if (res.status === 'success') {
              this.toastr.success('Delete initiated');
              setTimeout(() => {
                console.log('called after 3 sec');
                this.getInstanceData();
              }, 6000);
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
    dialog.componentInstance.applyManifestFor = 'endpoints';
    const preInputData: { [key: string]: any } = item;
    if (item.status) {
      delete preInputData.status;
    }
    if (preInputData.metadata.managedFields) {
      delete preInputData.metadata.managedFields;
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
          this.getInstanceData();
        }
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
      .getEndpoints(queryParam)
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
}
