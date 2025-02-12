import { DOCUMENT } from '@angular/common';
import { Component, Inject, OnDestroy, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ActivatedRoute, Router } from '@angular/router';
import icKeyboardBackspace from '@iconify/icons-ic/keyboard-backspace';
import icSearch from '@iconify/icons-ic/search';
import icAdd from '@iconify/icons-ic/twotone-add';
import icCross from '@iconify/icons-ic/twotone-cancel';
import icDelete from '@iconify/icons-ic/twotone-delete';
import icEdit from '@iconify/icons-ic/twotone-edit';
import icInfo from '@iconify/icons-ic/twotone-info';
import icLabel from '@iconify/icons-ic/twotone-label';
import icMoreHoriz from '@iconify/icons-ic/twotone-more-horiz';
import icRefresh from '@iconify/icons-ic/twotone-refresh';
import { K8sUpdateComponent } from '@k8s/k8s-update/k8s-update.component';
import { ToolbarService } from '@sdk-ui/services/toolbar.service';
import { ToastrService } from '@sdk-ui/ui';
import { ConfirmDialogStaticComponent } from '@shared-ui/ui';
import { fromEvent } from 'rxjs';
import { distinctUntilChanged, map, share, takeWhile, throttleTime } from 'rxjs/operators';
import { K8sPersistentVolumeService } from '../k8s-persistent-volume.service';

@Component({
  selector: 'kc-pv-list',
  templateUrl: './pv-list.component.html',
  styleUrls: ['./pv-list.component.scss']
})
export class PvListComponent implements OnInit, OnDestroy {
  icKeyboardBackspace = icKeyboardBackspace;
  icSearch = icSearch;
  icInfo = icInfo;
  icAdd = icAdd;
  icEdit = icEdit;
  icLabel = icLabel;
  icDelete = icDelete;
  icMoreHoriz = icMoreHoriz;
  icCross = icCross;
  icRefresh = icRefresh;
  searchBy = 'name';
  title: any = 'Persistent Volumes';
  isAlive: boolean = true;
  pvList: any[] = [];
  isLoading!: boolean;
  loadingSpanner: boolean = false;
  searchTerm = '';
  queryParams: any;
  resourceToken: string = '';
  loadMoreData: boolean = false;
  remaining;
  tokenReceiveTime: Date;

  constructor(
    private pvService: K8sPersistentVolumeService,
    private router: Router,
    private route: ActivatedRoute,
    private toolbarService: ToolbarService,
    public dialog: MatDialog,
    private toastr: ToastrService,
    @Inject(DOCUMENT) public document: any
  ) {}

  ngOnInit(): void {
    this.toolbarService.changeData({ title: this.title });
    this.queryParams = this.route.snapshot.queryParams;
    this.getPvList();
  }

  ngOnDestroy(): void {
    this.isAlive = false;
  }

  reloadList() {
    this.getPvList();
  }

  getPvList(queryParam?: any): void {
    this.isLoading = true;
    this.pvService
      .getPersitentVolume(queryParam)
      .pipe(takeWhile(() => this.isAlive))
      .subscribe({
        next: data => {
          if (data.status === 'success') {
            this.isLoading = false;
            this.pvList = data?.data.Result || [];
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

  onPvDetailsClick(item): void {
    // this.namespaceService.changeSelectedNamespace(item);
    this.router.navigate(['persistent-volume-details'], {
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
    dialog.componentInstance.applyManifestFor = 'persistent-volume';
    dialog.afterClosed().subscribe(res => {
      if (res) {
        if (res != null) {
          this.getPvList();
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
        this.pvService.deletePersistentVolume(item?.metadata?.name).subscribe(
          res => {
            if (res.status === 'success') {
              this.toastr.success('Delete initiated');
              setTimeout(() => {
                this.getPvList();
              }, 6000);
              this.getPvList();
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
    dialog.componentInstance.applyManifestFor = 'persistent-volume';

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
        if (res != null) {
          this.getPvList();
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
      this.getPvList(qp);
    }
    if (this.searchBy === 'name') {
      const qp = { q: this.searchTerm };
      this.getPvList(qp);
    }
  }

  clearSearch() {
    this.getPvList();
    this.searchTerm = '';
  }

  handleInputChange() {
    if (this.searchTerm.length === 0) {
      this.getPvList();
    }
  }

  extractCapitalLetters(str: string) {
    return str.replace(/[^A-Z]+/g, '');
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
      queryParam = { limit: this.pvList.length + 10 };
    } else {
      queryParam = { continue: this.resourceToken };
    }
    if (this.searchTerm.length > 0) {
      queryParam['q'] = this.searchTerm;
    }
    this.pvService
      .getPersitentVolume(queryParam)
      .pipe(takeWhile(() => this.isAlive))
      .subscribe({
        next: data => {
          this.loadMoreData = false;
          this.remaining = data.data.Remaining;
          this.resourceToken = data.data.Resource || '';
          this.tokenReceiveTime = new Date();
          if (queryParam.hasOwnProperty('limit')) {
            this.pvList = data.data.Result || [];
          } else {
            this.pvList = this.pvList.concat(data.data.Result) || [];
          }
          this.loadMoreData = false;
        },
        error: err => {
          this.toastr.error('Failed: ', err.error.message);
        }
      });
  }
}
