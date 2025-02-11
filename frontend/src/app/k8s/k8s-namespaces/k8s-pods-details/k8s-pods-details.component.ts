import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ActivatedRoute, Router } from '@angular/router';
import icArrowBack from '@iconify/icons-ic/arrow-back';
import icStopped from '@iconify/icons-ic/do-not-disturb-on';
import icRunning from '@iconify/icons-ic/sharp-check-circle';
import icDelete from '@iconify/icons-ic/twotone-delete';
import icLogs from '@iconify/icons-ic/twotone-description';
import icEdit from '@iconify/icons-ic/twotone-edit';
import icRight from '@iconify/icons-ic/twotone-greater-than';
import icCircle from '@iconify/icons-ic/twotone-lens';
import { JsonDataViewerComponent } from '@k8s/common/components/json-data-viewer/json-data-viewer.component';
import { K8sUpdateComponent } from '@k8s/k8s-update/k8s-update.component';
import { K8sService } from '@k8s/k8s.service';
import { ToolbarService } from '@sdk-ui/services/toolbar.service';
import { ToastrService } from '@sdk-ui/ui';
import { ConfirmDialogStaticComponent } from '@shared-ui/ui';
import { K8sNamespacesService } from '../k8s-namespaces.service';
import { GrafanaDashboardComponent } from './grafana-dashboard/grafana-dashboard.component';
import { K8sPodsContainerLogComponent } from './k8s-pods-container-log/k8s-pods-container-log.component';

@Component({
  selector: 'kc-k8s-pods-details',
  templateUrl: './k8s-pods-details.component.html',
  styleUrls: ['./k8s-pods-details.component.scss']
})
export class K8sPodsDetailsComponent implements OnInit {
  icCircle = icCircle;

  isLoading: boolean = true;
  data: any;
  eventsData: any;
  queryParams: any;
  namespaceInstance: string;
  icEdit = icEdit;
  icLogs = icLogs;
  icRight = icRight;
  icStopped = icStopped;
  icRunning = icRunning;
  icDelete = icDelete;
  icArrowBack = icArrowBack;
  clusterId: string;
  lokiDetails: any;
  title = 'Pods';

  constructor(
    private _namespaceService: K8sNamespacesService,
    private k8sService: K8sService,
    private route: ActivatedRoute,
    private toastr: ToastrService,
    private dialog: MatDialog,
    private router: Router,
    private toolbarService: ToolbarService
  ) {}

  ngOnInit(): void {
    this.toolbarService.changeData({ title: this.title });
    this.queryParams = this.route.snapshot.queryParams;
    this.namespaceInstance = this.route.snapshot.params?.name;
    this.clusterId = this.k8sService.clusterIdSnapshot;
    this.getInstanceData();
  }

  getInstanceData(): void {
    this.isLoading = true;
    this._namespaceService.getPodsDetails(this.namespaceInstance).subscribe({
      next: res => {
        if (res?.data?.Result?.metadata?.name) {
          this.data = res?.data?.Result;
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
        this._namespaceService.deleteNamespacePods(item?.metadata?.name).subscribe(
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

  viewLogs(containerName?: any): void {
    const containers = this.data?.spec?.containers;
    const allContainerNames = [];
    containers.forEach(element => {
      allContainerNames.push(element.name);
    });
    const data = {
      pod: this.data?.metadata?.name,
      namespace: this.data?.metadata?.namespace,
      restart:
        this.data?.status?.containerStatuses?.length && this.data?.status?.containerStatuses[0]?.restartCount
          ? this.data?.status?.containerStatuses[0]?.restartCount
          : 0,
      allContainers: allContainerNames
    };
    if (containerName) {
      data['container'] = containerName;
    }

    this.dialog.open(K8sPodsContainerLogComponent, {
      height: '90%',
      width: '85%',
      disableClose: true,
      data: data
    });
  }

  onUpdate(item: any): void {
    const dialog = this.dialog.open(K8sUpdateComponent, {
      minHeight: '300px',
      width: '900px',
      disableClose: true
    });
    dialog.componentInstance.isEditMode = true;
    dialog.componentInstance.applyManifestFor = 'pods';

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

  get filterPodsStatusMes(): any {
    const podData = this.data;
    const status = podData.status.phase;
    if (status === 'Running') {
      if (podData?.status?.containerStatuses?.length && podData.status.containerStatuses[0]?.state?.terminated) {
        return podData.status.containerStatuses[0].state.terminated.reason;
      }
    }
    if (status === 'Pending') {
      if (podData?.status?.containerStatuses.length && podData.status.containerStatuses[0]?.state?.waiting) {
        return podData.status.containerStatuses[0].state.waiting.reason;
      }
    }
    return status;
  }

  jsonViewer(data: any, title: any): void {
    const dialogData = {
      data,
      title: title
    };
    this.dialog.open(JsonDataViewerComponent, {
      minHeight: '20%',
      minWidth: '20%',
      disableClose: false,
      data: dialogData
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

  btoa(string) {
    return btoa(string);
  }

  isInt(value: string): boolean {
    const parsedValue = parseInt(value, 10);
    return !isNaN(parsedValue) && String(parsedValue) === value;
  }

  isObject(value: any): boolean {
    return typeof value === 'object' && value !== null;
  }

  open(): void {
    const dialogRef = this.dialog.open(GrafanaDashboardComponent, {
      width: '100%',
      maxWidth: '1280px',
      height: '100vh',
      disableClose: false,
      position: { top: '0', right: '0' }
    });
  }
}
