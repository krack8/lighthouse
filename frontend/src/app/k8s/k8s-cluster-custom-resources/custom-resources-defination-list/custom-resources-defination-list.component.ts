import { Component, Inject, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { distinctUntilChanged, map, share, takeWhile, throttleTime } from 'rxjs/operators';
import icSearch from '@iconify/icons-ic/search';
import icInfo from '@iconify/icons-ic/twotone-info';
import icDown from '@iconify/icons-ic/twotone-arrow-drop-down';
import icAdd from '@iconify/icons-ic/twotone-add';
import icCross from '@iconify/icons-ic/twotone-cancel';
import icEdit from '@iconify/icons-ic/twotone-edit';
import icLabel from '@iconify/icons-ic/twotone-label';
import icRefresh from '@iconify/icons-ic/twotone-refresh';
import icDelete from '@iconify/icons-ic/twotone-delete';
import icMoreHoriz from '@iconify/icons-ic/twotone-more-horiz';
import { ToolbarService } from '@sdk-ui/services/toolbar.service';
import icKeyboardBackspace from '@iconify/icons-ic/keyboard-backspace';
import { MatDialog } from '@angular/material/dialog';
import { K8sUpdateComponent } from '@k8s/k8s-update/k8s-update.component';
import { ToastrService } from '@sdk-ui/ui';
import { ConfirmDialogStaticComponent } from '@shared-ui/ui';
import { K8sClusterCustomResourcesService } from '../k8s-cluster-custom-resources.service';
import { CustomResourcesDefinationDetailsComponent } from '../custom-resources-defination-details/custom-resources-defination-details.component';
import { DOCUMENT } from '@angular/common';
import { fromEvent } from 'rxjs';

@Component({
  selector: 'kc-custom-resources-defination-list',
  templateUrl: './custom-resources-defination-list.component.html',
  styleUrls: ['./custom-resources-defination-list.component.scss']
})
export class CustomResourcesDefinationListComponent implements OnInit, OnDestroy {
  icKeyboardBackspace = icKeyboardBackspace;
  icSearch = icSearch;
  icInfo = icInfo;
  icAdd = icAdd;
  icDown = icDown;
  icCross = icCross;
  icEdit = icEdit;
  icLabel = icLabel;
  icDelete = icDelete;
  icRefresh = icRefresh;
  icMoreHoriz = icMoreHoriz;
  isAlive: boolean = true;
  customResources: any[] = [];
  total;
  isLoading!: boolean;
  searchTerm = '';
  searchBy = 'name';
  title: any = 'Custom Resource Defination';
  queryParams: any;
  resourceToken: string = '';
  loadMoreData: boolean = false;
  tokenReceiveTime: Date;
  searchObject: any = {};

  constructor(
    @Inject(DOCUMENT) public document: any,
    private CustomResourcesService: K8sClusterCustomResourcesService,
    private router: Router,
    private route: ActivatedRoute,
    private toolbarService: ToolbarService,
    public dialog: MatDialog,
    private toastr: ToastrService
  ) {}

  ngOnInit(): void {
    this.toolbarService.changeData({ title: this.title });
    this.queryParams = this.route.snapshot.queryParams;
    this.getCrdList();
  }

  ngOnDestroy(): void {
    this.isAlive = false;
  }

  reloadList() {
    this.getCrdList();
  }

  getCrdList(qp?: any): void {
    this.isLoading = true;
    this.CustomResourcesService.getCustomResourcesDefination(qp)
      .pipe(takeWhile(() => this.isAlive))
      .subscribe({
        next: data => {
          if (data?.status === 'success') {
            this.isLoading = false;
            this.resourceToken = data.data.Resource;
            this.tokenReceiveTime = new Date();
            this.total = data.data.Total;
            this.customResources = data.data.CrdForList || [];
          } else {
            this.isLoading = false;
            //this.toastr.error('Failed: ', data.message);
          }
        },
        error: err => {
          this.isLoading = false;
          this.toastr.error('Failed: ', err.error.message);
        }
      });
  }

  onCrdClick(item): void {
    this.router.navigate(['custom-resources'], {
      queryParams: {
        ...this.route.snapshot.queryParams,
        resource: item?.name_plural,
        group: item?.group,
        version: item?.version[0],
        versions: item?.version.join(','),
        kind: item?.kind
      },
      relativeTo: this.route
    });
  }

  openDetails(name: string) {
    this.dialog.open(CustomResourcesDefinationDetailsComponent, {
      disableClose: false,
      width: '1600px',
      maxWidth: '1600px',
      maxHeight: '90vh',
      data: name
    });
  }

  onCreate(): void {
    const dialog = this.dialog.open(K8sUpdateComponent, {
      minHeight: '300px',
      width: '900px',
      disableClose: true
    });
    dialog.componentInstance.applyManifestFor = 'crd';
  }

  onDelete(item: any): void {
    const dialogRef = this.dialog.open(ConfirmDialogStaticComponent, {
      disableClose: true,
      minWidth: '350px',
      data: {
        message: `Are you sure! want to delete ${item?.name}?`,
        icon: '/assets/img/bin.svg'
      }
    });
    dialogRef.afterClosed().subscribe((bool: boolean) => {
      if (bool === true) {
        this.CustomResourcesService.deleteCustomResourcesDefination(item?.name).subscribe(
          res => {
            if (res.status === 'success') {
              this.toastr.success('Delete initiated');
              this.getCrdList();
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
    this.isLoading = true;
    let crdDetails: any = {};
    console.log('item', item);
    const temp = item.name;
    console.log('temp', temp);
    this.CustomResourcesService.getCustomResourceDefinationDetails(item?.name).subscribe(
      res => {
        if (res.status === 'success') {
          this.isLoading = false;
          crdDetails = res.data;
          if (crdDetails) {
            const dialog = this.dialog.open(K8sUpdateComponent, {
              minHeight: '300px',
              width: '900px',
              disableClose: true
            });
            dialog.componentInstance.isEditMode = true;
            dialog.componentInstance.applyManifestFor = 'crd';
            //dialog.componentInstance.queryParams = this.queryParams;

            const metaTemp: { [key: string]: any } = {};
            metaTemp.name = crdDetails.metadata.name;
            metaTemp.namespace = crdDetails.metadata.namespace;
            metaTemp.uid = crdDetails.metadata.uid;

            if (crdDetails.metadata.selfLink) {
              metaTemp.selfLink = crdDetails.metadata.selfLink;
            }
            if (crdDetails.metadata.labels) {
              metaTemp.labels = crdDetails.metadata.labels;
            }
            if (crdDetails.metadata.annotations) {
              metaTemp.annotations = crdDetails.metadata.annotations;
            }

            const preInputData: { [key: string]: any } = {};

            if (crdDetails.apiVersion) {
              preInputData.apiVersion = crdDetails.apiVersion;
            }

            if (crdDetails.apiVersion) {
              preInputData.kind = crdDetails.kind;
            }

            preInputData.metadata = metaTemp;

            if (crdDetails.spec) {
              preInputData.spec = crdDetails.spec;
            }

            dialog.componentInstance.preInputData = preInputData;

            dialog.componentInstance.payload = {
              name: crdDetails.metadata.name,
              kind: crdDetails.kind,
              apiVersion: crdDetails.apiVersion,
              namespace: crdDetails?.metadata?.namespace
            };

            dialog.afterClosed().subscribe(res => {
              if (res) {
                if (res != null) {
                  this.getCrdList();
                }
              }
            });
          }
        } else {
          this.toastr.error('Failed: ', res.message);
        }
      },
      err => {
        this.toastr.error('Failed: ', err.error.message);
      }
    );
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
      queryParam = { limit: this.customResources.length + 10 };
    } else {
      queryParam = { continue: this.resourceToken };
    }
    if (this.searchTerm.length > 0) {
      queryParam['q'] = this.searchTerm;
    }
    queryParam = { ...queryParam, ...this.searchObject };
    this.CustomResourcesService.getCustomResourcesDefination(queryParam)
      .pipe(takeWhile(() => this.isAlive))
      .subscribe({
        next: data => {
          this.loadMoreData = false;
          this.resourceToken = data.data.Resource || '';
          this.tokenReceiveTime = new Date();
          if (queryParam.hasOwnProperty('limit')) {
            this.customResources = data.data.CrdForList || [];
          } else {
            this.customResources = this.customResources.concat(data.data.CrdForList) || [];
          }
          this.loadMoreData = false;
        },
        error: err => {
          this.toastr.error('Failed: ', err.error.message);
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
      this.searchObject = qp;
      this.getCrdList(qp);
    }
    if (this.searchBy === 'name') {
      const qp = { q: this.searchTerm };
      this.searchObject = qp;
      this.getCrdList(qp);
    }
  }

  clearSearch() {
    this.getCrdList();
    this.searchTerm = '';
    this.searchObject = {};
  }
  handleInputChange() {
    if (this.searchTerm.length === 0) {
      this.searchObject = {};
      this.getCrdList();
    }
  }
}
