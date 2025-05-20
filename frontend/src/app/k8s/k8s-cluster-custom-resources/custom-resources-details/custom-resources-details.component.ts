import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import icArrowBack from '@iconify/icons-ic/arrow-back';
import icEdit from '@iconify/icons-ic/twotone-edit';
import icDelete from '@iconify/icons-ic/twotone-delete';
import { K8sClusterCustomResourcesService } from '../k8s-cluster-custom-resources.service';
import { MatDialog } from '@angular/material/dialog';
import { ConfirmDialogStaticComponent } from '@shared-ui/ui';
import { K8sUpdateComponent } from '@k8s/k8s-update/k8s-update.component';
import { ToastrService } from '@sdk-ui/ui';

@Component({
  selector: 'kc-custom-resources-details',
  templateUrl: './custom-resources-details.component.html',
  styleUrls: ['./custom-resources-details.component.scss']
})
export class CustomResourcesDetailsComponent implements OnInit {
  data: any = {};
  queryParams: any;
  isLoading = false;
  icArrowBack = icArrowBack;
  icEdit = icEdit;
  icDelete = icDelete;

  constructor(
    private route: ActivatedRoute,
    private CustomResourcesService: K8sClusterCustomResourcesService,
    public dialog: MatDialog,
    private toastr: ToastrService,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.queryParams = this.route.snapshot.queryParams;
    console.log('this.queryParams', this.queryParams);
    this.getDetails();
  }

  getDetails(): void {
    this.isLoading = true;

    const qp = {
      resource: this.queryParams?.resource,
      group: this.queryParams?.group,
      version: this.queryParams?.version,
      namespace: this.queryParams?.namespace || ''
    };
    this.CustomResourcesService.getCustomResourceDetails(this.queryParams?.name, qp).subscribe({
      next: data => {
        if (data?.status === 'success') {
          this.data = data.data || [];
          this.isLoading = false;
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

  isInt(value: string): boolean {
    const parsedValue = parseInt(value, 10);
    return !isNaN(parsedValue) && String(parsedValue) === value;
  }

  isObject(value: any): boolean {
    return typeof value === 'object' && value !== null;
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
        this.CustomResourcesService.deleteCustomResources(item?.metadata?.name, this.queryParams).subscribe(
          res => {
            if (res.status === 'success') {
              this.toastr.success('Delete initiated');
              const queryParams = {
                resource: this.queryParams.resource,
                group: this.queryParams.group,
                version: this.queryParams.version,
                versions: this.queryParams.versions,
                kind: this.queryParams.kind
              };
              this.router.navigate(['../'], { relativeTo: this.route, queryParams });
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
    dialog.componentInstance.applyManifestFor = 'custom-resources';
    dialog.componentInstance.queryParams = this.queryParams;

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

    if (item.apiVersion) {
      preInputData.apiVersion = item.apiVersion;
    }

    if (item.apiVersion) {
      preInputData.kind = item.kind;
    }

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
        this.getDetails();
      }
    });
  }

  isObjectEmpty(obj: any): boolean {
    return Object.keys(obj).length === 0;
  }
}
