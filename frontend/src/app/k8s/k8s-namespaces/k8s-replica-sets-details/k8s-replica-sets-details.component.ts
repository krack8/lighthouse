import { Component, OnInit, ViewChild } from '@angular/core';
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
import icCircle from '@iconify/icons-ic/twotone-lens';

@Component({
  selector: 'kc-k8s-replica-sets-details',
  templateUrl: './k8s-replica-sets-details.component.html',
  styleUrls: ['./k8s-replica-sets-details.component.scss']
})
export class K8sReplicaSetsDetailsComponent implements OnInit {
  icCircle = icCircle;
  isLoading: boolean = true;
  data: any;
  eventsData: any;
  queryParams: any;
  namespaceInstance: string;
  icEdit = icEdit;
  icDelete = icDelete;
  icArrowBack = icArrowBack;
  icRight = icRight;
  title = 'ReplicaSets';

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
    this._namespaceService.getReplicaSetsDetails(this.namespaceInstance).subscribe({
      next: res => {
        if (res?.data?.metadata?.name) {
          this.data = res?.data;
          this.isLoading = false;
        }
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
        this._namespaceService.deleteNamespaceReplicasets(item?.metadata?.name).subscribe(
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
    dialog.componentInstance.applyManifestFor = 'replica-sets';

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
          this.getInstanceData();
      }
    });
  }

  isConditionNegative(condition): boolean {
    const type = condition.type;
    const status = condition.status;
    if (['PodScheduled', 'Ready', 'Initialized', 'ContainersReady'].includes(type) && status === 'True') {
      return false;
    }
    if (type === 'Unschedulable' && status === 'False') {
      return false;
    }
    return true;
  }

  isInt(value: string): boolean {
    const parsedValue = parseInt(value, 10);
    return !isNaN(parsedValue) && String(parsedValue) === value;
  }

  isObject(value: any): boolean {
    return typeof value === 'object' && value !== null;
  }
}
