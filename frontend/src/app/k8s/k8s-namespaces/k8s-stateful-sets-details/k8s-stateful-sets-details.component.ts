import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ActivatedRoute, Router } from '@angular/router';
import icArrowBack from '@iconify/icons-ic/arrow-back';
import icDelete from '@iconify/icons-ic/twotone-delete';
import icEdit from '@iconify/icons-ic/twotone-edit';
import icRight from '@iconify/icons-ic/twotone-greater-than';
import icCircle from '@iconify/icons-ic/twotone-lens';
import { K8sUpdateComponent } from '@k8s/k8s-update/k8s-update.component';
import { fadeInRight400ms } from '@sdk-ui/animations/fade-in-right.animation';
import { ToolbarService } from '@sdk-ui/services/toolbar.service';
import { ToastrService } from '@sdk-ui/ui';
import { ConfirmDialogStaticComponent } from '@shared-ui/ui';
import { K8sNamespacesService } from '../k8s-namespaces.service';

@Component({
  selector: 'kc-k8s-stateful-sets-details',
  templateUrl: './k8s-stateful-sets-details.component.html',
  styleUrls: ['./k8s-stateful-sets-details.component.scss'],
  animations: [fadeInRight400ms]
})
export class K8sStatefulSetsDetailsComponent implements OnInit {
  icCircle = icCircle;
  isLoading: boolean = true;
  data: any;
  eventsData: any;
  queryParams: any;
  podListData: any = {};
  namespaceInstance: string;
  icEdit = icEdit;
  icDelete = icDelete;
  icArrowBack = icArrowBack;
  icRight = icRight;
  title = 'StatefulSets';
  graphStats: any;
  statsLoaded: boolean = false;

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
    this._namespaceService.getStatefulSetsDetails(this.namespaceInstance).subscribe({
      next: res => {
        if (res?.data?.metadata?.name) {
          this.data = res?.data;
          this.podListData = this.data.metadata.name;
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
        this._namespaceService.deleteNamespaceStatefulset(item?.metadata?.name).subscribe(
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
    dialog.componentInstance.applyManifestFor = 'stateful-set';

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
        if (res != null) {
          this.getInstanceData();
        }
      }
    });
  }

  isTypeObject(param: any): boolean {
    if (typeof param === 'object') {
      return true;
    } else return false;
  }

  isConditionNegative(condition): boolean {
    const type = condition.type;
    const status = condition.status;
    const types = ['PodScheduled', 'Ready', 'Initialized', 'ContainersReady'];
    if (types.includes(type) && status === 'True') {
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
