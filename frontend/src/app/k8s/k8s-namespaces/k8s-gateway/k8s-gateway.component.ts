import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import icSearch from '@iconify/icons-ic/search';
import { ToastrService } from '@sdk-ui/ui';
import { ToolbarService } from '@sdk-ui/services/toolbar.service';
import { K8sNamespacesService } from '../k8s-namespaces.service';
import icEdit from '@iconify/icons-ic/twotone-edit';
import icCross from '@iconify/icons-ic/twotone-cancel';
import icRefresh from '@iconify/icons-ic/twotone-refresh';
import icLabel from '@iconify/icons-ic/twotone-label';
import icDelete from '@iconify/icons-ic/twotone-delete';
import icInfo from '@iconify/icons-ic/twotone-info';
import icAdd from '@iconify/icons-ic/twotone-add';
import icMoreVert from '@iconify/icons-ic/twotone-more-vert';
import icUpgrade from '@iconify/icons-ic/twotone-file-upload';
import { MatDialog } from '@angular/material/dialog';
import { K8sUpdateComponent } from '@k8s/k8s-update/k8s-update.component';
import { ConfirmDialogStaticComponent } from '@shared-ui/ui';

@Component({
  selector: 'kc-k8s-gateway',
  templateUrl: './k8s-gateway.component.html',
  styleUrls: ['./k8s-gateway.component.scss']
})
export class K8sGatewayComponent implements OnInit {
  icSearch = icSearch;
  icMoreVert = icMoreVert;
  icInfo = icInfo;
  icDelete = icDelete;
  icAdd = icAdd;
  icRefresh = icRefresh;
  icEdit = icEdit;
  icLabel = icLabel;
  icUpgrade = icUpgrade;
  icCross = icCross;
  searchBy = 'name';
  title: any = 'Gateway';
  isAlive: boolean = true;
  loadingSpanner: boolean = false;
  data: any[] = [];
  searchTerm: string = '';
  queryParams: any;
  kubeClusterId = '';
  namespace = '';

  constructor(
    private namespaceService: K8sNamespacesService,
    private route: ActivatedRoute,
    private toolbarService: ToolbarService,
    private toastr: ToastrService,
    private dialog: MatDialog
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
    this.namespaceService.getGateway(queryParam).subscribe({
      next: res => {
        if (res?.status === 'error') {
          this.toastr.error(res?.message);
        }
        this.data = res?.data || [];
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
    dialog.componentInstance.applyManifestFor = 'gateway';
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
        this.namespaceService.deleteNamespaceGateway(item?.metadata?.name).subscribe(
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
    dialog.componentInstance.applyManifestFor = 'gateway';

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
}
