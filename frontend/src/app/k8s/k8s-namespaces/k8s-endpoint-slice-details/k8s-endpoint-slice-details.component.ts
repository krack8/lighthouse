import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { ToolbarService } from '@sdk-ui/services/toolbar.service';
import { K8sNamespacesService } from '../k8s-namespaces.service';
import icArrowBack from '@iconify/icons-ic/arrow-back';
import icRight from '@iconify/icons-ic/twotone-greater-than';
import icEdit from '@iconify/icons-ic/twotone-edit';
import icDelete from '@iconify/icons-ic/twotone-delete';
import { ConfirmDialogStaticComponent } from '@shared-ui/ui';
import { MatDialog } from '@angular/material/dialog';
import { ToastrService } from '@sdk-ui/ui';
import { K8sUpdateComponent } from '@k8s/k8s-update/k8s-update.component';

@Component({
  selector: 'kc-k8s-endpoint-slice-details',
  templateUrl: './k8s-endpoint-slice-details.component.html',
  styleUrls: ['./k8s-endpoint-slice-details.component.scss']
})
export class K8sEndpointSliceDetailsComponent implements OnInit {
  isLoading: boolean = true;
  data: any;
  queryParams: any;
  namespaceInstance: string;
  icArrowBack = icArrowBack;
  icRight = icRight;
  icEdit = icEdit;
  icDelete = icDelete;
  title = 'Endpoint Slice';

  constructor(
    private _namespaceService: K8sNamespacesService,
    private route: ActivatedRoute,
    private toolbarService: ToolbarService,
    private toastr: ToastrService,
    private dialog: MatDialog,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.queryParams = this.route.snapshot.queryParams;
    this.namespaceInstance = this.route.snapshot.params?.name;
    this.toolbarService.changeData({ title: this.title });
    this.getInstanceData();
  }

  getInstanceData(): void {
    this.isLoading = true;
    this._namespaceService.getEndpointSliceDetails(this.namespaceInstance).subscribe({
      next: res => {
        if (res?.data?.metadata?.name) {
          this.data = res?.data;
        }
        this.isLoading = false;
      },
      error: err => {
        this.isLoading = false;
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
        this._namespaceService.deleteNamespaceEndpointSlice(item?.metadata?.name).subscribe(
          res => {
            if (res.status === 'success') {
              const queryParams = this.queryParams;
              this.router.navigate(['../'], { relativeTo: this.route, queryParams });
              this.toastr.success('Delete initiated');
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
    dialog.componentInstance.applyManifestFor = 'endpoint-slice';

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
    if (item.data) {
      preInputData.data = item.data;
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
}
