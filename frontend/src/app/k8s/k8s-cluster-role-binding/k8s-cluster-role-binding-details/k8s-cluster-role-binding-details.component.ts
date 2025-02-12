import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ActivatedRoute, Router } from '@angular/router';
import icArrowBack from '@iconify/icons-ic/arrow-back';
import icEdit from '@iconify/icons-ic/twotone-edit';
import icDelete from '@iconify/icons-ic/twotone-delete';
import { ToastrService } from '@sdk-ui/ui';
import { ToolbarService } from '@sdk-ui/services/toolbar.service';
import { ConfirmDialogStaticComponent } from '@shared-ui/ui';
import { K8sClusterRoleBindingService } from '../k8s-cluster-role-binding.service';
import { K8sUpdateComponent } from '@k8s/k8s-update/k8s-update.component';

@Component({
  selector: 'kc-k8s-cluster-role-binding-details',
  templateUrl: './k8s-cluster-role-binding-details.component.html',
  styleUrls: ['./k8s-cluster-role-binding-details.component.scss']
})
export class K8sClusterRoleBindingDetailsComponent implements OnInit {
  data: any = {};
  queryParams: any;
  isLoading = false;
  icArrowBack = icArrowBack;
  icEdit = icEdit;
  icDelete = icDelete;

  constructor(
    private route: ActivatedRoute,
    private clusterRoleBindingService: K8sClusterRoleBindingService,
    private toolbarService: ToolbarService,
    private router: Router,
    public dialog: MatDialog,
    private toastr: ToastrService
  ) {}

  ngOnInit(): void {
    this.queryParams = this.route.snapshot.queryParams;
    this.getDetails();
  }

  getDetails(): void {
    this.isLoading = true;
    this.clusterRoleBindingService.getClusterRoleBindingDetails(this.queryParams.name).subscribe({
      next: data => {
        if (data.status === 'success') {
          this.data = data.data || [];
          this.isLoading = false;
        } else {
          this.isLoading = false;
        }
      },
      error: err => {
        this.toastr.error('Failed: ', err.error.message);
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
        this.clusterRoleBindingService.deleteClusterRoleBinding(item?.metadata?.name).subscribe(
          res => {
            if (res.status === 'success') {
              this.router.navigate(['../'], { relativeTo: this.route});
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
    dialog.componentInstance.applyManifestFor = 'cluster-role-binding';

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
    if (item.roleRef) {
      preInputData.roleRef = item.roleRef;
    }
    if (item.subjects) {
      preInputData.subjects = item.subjects;
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
          this.getDetails();
        }
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
}
