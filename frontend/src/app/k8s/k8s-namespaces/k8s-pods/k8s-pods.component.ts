import { DOCUMENT } from '@angular/common';
import { Component, Inject, OnInit, ViewChild } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ActivatedRoute, Router } from '@angular/router';
import icSearch from '@iconify/icons-ic/search';
import icAdd from '@iconify/icons-ic/twotone-add';
import icCross from '@iconify/icons-ic/twotone-cancel';
import icDelete from '@iconify/icons-ic/twotone-delete';
import icLogs from '@iconify/icons-ic/twotone-description';
import icEdit from '@iconify/icons-ic/twotone-edit';
import icUpgrade from '@iconify/icons-ic/twotone-file-upload';
import icInfo from '@iconify/icons-ic/twotone-info';
import icLabel from '@iconify/icons-ic/twotone-label';
import icCircle from '@iconify/icons-ic/twotone-lens';
import icMoreVert from '@iconify/icons-ic/twotone-more-vert';
import icRefresh from '@iconify/icons-ic/twotone-refresh';
import { K8sUpdateComponent } from '@k8s/k8s-update/k8s-update.component';
import { K8sService } from '@k8s/k8s.service';
import { fadeInRight400ms } from '@sdk-ui/animations/fade-in-right.animation';
import { ToolbarService } from '@sdk-ui/services/toolbar.service';
import { ToastrService } from '@sdk-ui/ui';
import { ConfirmDialogStaticComponent } from '@shared-ui/ui';
import { fromEvent } from 'rxjs';
import { distinctUntilChanged, map, share, takeWhile, tap, throttleTime } from 'rxjs/operators';
import { K8sNamespacesService } from '../k8s-namespaces.service';
import { K8sPodsContainerLogComponent } from '../k8s-pods-details/k8s-pods-container-log/k8s-pods-container-log.component';
import { ApexChart, ApexNonAxisChartSeries, ApexPlotOptions, ChartComponent } from 'ng-apexcharts';

// export type ChartOptions = {
//   series: ApexNonAxisChartSeries;
//   chart: ApexChart;
//   labels: string[];
//   color: string[];
//   plotOptions: ApexPlotOptions;
//  };

@Component({
  selector: 'kc-k8s-pods',
  templateUrl: './k8s-pods.component.html',
  styleUrls: ['./k8s-pods.component.scss'],
  animations: [fadeInRight400ms]
})
export class K8sPodsComponent implements OnInit {

  //pod metrics
  @ViewChild("chart") chart: ChartComponent;
  public chartOptions: any;

  TotalPods: number = 0;
  RunningPods: number = 0;
  PendingPods: number = 0;
  FailedPods: number = 0;
  TotalCPU: number = 0;
  TotalMemory: number = 0;

  podChart = {
    series: [],
    labels: ['Running', 'Pending', 'Failed'],
    color: ['#36c678', '#ffbb33', '#b1122a'],
  }

  //icons
  icCircle = icCircle;
  icSearch = icSearch;
  icInfo = icInfo;
  icDelete = icDelete;
  icEdit = icEdit;
  icLogs = icLogs;
  icAdd = icAdd;
  icUpgrade = icUpgrade;
  icLabel = icLabel;
  icMoreVert = icMoreVert;
  icRefresh = icRefresh;
  icCross = icCross;

  searchBy = 'name';
  title: any = 'Pods';
  isAlive: boolean = true;
  loadingSpanner: boolean = false;
  data: any;
  searchTerm: string = '';
  queryParams: any;
  stats: any;
  clusterId: string;
  statsLoaded: boolean = false;
  resourceToken: string = '';
  loadMoreData: boolean = false;
  remaining;
  tokenReceiveTime: Date;

  constructor(
    private k8sService: K8sService,
    private namespaceService: K8sNamespacesService,
    private route: ActivatedRoute,
    private toolbarService: ToolbarService,
    private toastr: ToastrService,
    private dialog: MatDialog,
    @Inject(DOCUMENT) public document: any,
    private router: Router
  ) {
    this.chartOptions = {
      //series: [44, 55, 41, 17, 15],
      chart: {
        type: 'donut',
        height: 170,
        width: 200,
        animations: {
          enabled: true,
          easing: 'easeinout',
          speed: 800,
          animateGradually: {
            enabled: true,
            delay: 150
          },
          dynamicAnimation: {
            enabled: true,
            speed: 350
          }
        }
      },
      plotOptions: {
        pie: {
          donut: {
            size: '65%',
            background: 'transparent',

          }
        }
      },
      dataLabels: {
        enabled: false,
        style: {
          fontSize: '14px',
          fontFamily: 'Helvetica, Arial, sans-serif',
          fontWeight: 'bold',
          colors: ['#fff']
        },
        dropShadow: {
          enabled: false,
          top: 1,
          left: 1,
          blur: 1,
          opacity: 0.45
        }
      },
      stroke: {
        show: false, // Hide the border of the donut
      },
      legend: {
        show: false,
      },
      responsive: [{
        breakpoint: 580,
        options: {
          chart: {
            width: 100
          },
          legend: {
            position: 'bottom'
          }
        }
      }],
    };
  }

  ngOnInit(): void {
    this.toolbarService.changeData({ title: this.title });
    this.clusterId = this.k8sService.clusterIdSnapshot;
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
    this.getPodStats(queryParam);
    this.namespaceService
      .getPods(queryParam)
      .pipe(takeWhile(() => this.isAlive))
      .subscribe({
        next: res => {
          this.loadingSpanner = true;
          if (res?.status === 'success') {
            this.resourceToken = res.data.Resource;
            this.tokenReceiveTime = new Date();
            this.remaining = res.data.Remaining;
            this.data = res?.data.Result || [];
            this.loadingSpanner = false;
          } else {
            this.loadingSpanner = false;
          }
        },
        error: err => {
          this.loadingSpanner = false;
          this.toastr.error(err?.message);
        }
      });
  }

  getPodStats(queryParam?: any): void {
    this.statsLoaded = true;
    this.namespaceService.getPodStats(queryParam).subscribe({
      next: res => {
        if (res?.status === 'success') {
          this.stats = res?.data;
          //initializing metrics data
          this.TotalPods = res?.data?.Total || 0;
          this.RunningPods = res?.data?.Running || 0;
          this.PendingPods = res?.data?.Pending || 0;
          this.FailedPods = res?.data?.Failed || 0;
          this.TotalCPU = res?.data?.CPU || 0;
          this.TotalMemory = res?.data?.Memory || 0;
          this.podChart.series = [this.RunningPods, this.PendingPods, this.FailedPods];
          this.statsLoaded = false;
        }
        if (res?.status === 'error') {
          this.toastr.error(res?.message);
          this.statsLoaded = false;
        }
      },
      error: err => {
        this.statsLoaded = false;
        this.toastr.error(err?.message);
      }
    });
  }

  calcMemory(containers, _type = 'requests'): any {
    if (containers.length === 1) {
      try {
        if (containers[0].resources[_type].memory !== undefined) {
          return containers[0].resources[_type].memory;
        } else {
          return '-';
        }
      } catch (err) {
        return '-';
      }
    } else if (containers.length > 1) {
      let count = 0;
      containers.map(container => {
        if (container.resources[_type] && container.resources[_type].memory && container.resources[_type].memory.includes('Gi')) {
          const asMi = parseFloat(container.resources[_type].memory) * 1024;
          count = count + asMi;
        } else if (container.resources[_type] && container.resources[_type].memory && container.resources[_type].memory.includes('Mi')) {
          const asNumber = parseFloat(container.resources[_type].memory);
          count = count + asNumber;
        }
      });
      if (count < 1024) {
        return `${count}Mi`;
      } else if (count < 1024 ** 2) {
        return `${(count / 1024).toFixed(2)}Gi`;
      } else if (count < 1024 ** 3) {
        return `${(count / 1024 ** 2).toFixed(2)}Tb`;
      }
      return count || '-';
    }
    return '-';
  }
  calcCpu(containers, _type = 'requests') {
    if (containers.length === 1) {
      try {
        return containers[0].resources[_type].cpu;
      } catch (err) {
        return '-';
      }
    } else if (containers.length > 1) {
      let count = 0;
      containers.map(container => {
        if (container.resources[_type] && container.resources[_type].cpu && container.resources[_type].cpu.includes('m')) {
          const asNumber = parseFloat(container.resources[_type].cpu);
          count = count + asNumber;
        } else if (container.resources[_type] && container.resources[_type].cpu) {
          const asMillie = parseFloat(container.resources[_type].cpu) * 1000;
          count = count + asMillie;
        }
      });
      if (count < 1000) {
        return `${count}m`;
      } else if (count < 1000 ** 2) {
        return `${(count / 1000).toFixed(2)}`;
      }
      return count || '-';
    }
    return '-';
  }

  containerRunning(containers: any[]) {
    let count = 0;
    if (containers) {
      containers.forEach(container => {
        if (container.ready === true) {
          count++;
        }
      });
      return count;
    } else {
      return '0';
    }
  }

  onCreate(): void {
    const dialog = this.dialog.open(K8sUpdateComponent, {
      minHeight: '300px',
      width: '900px',
      disableClose: true
    });
    dialog.componentInstance.applyManifestFor = 'pods';
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
        this.namespaceService.deleteNamespacePods(item?.metadata?.name).subscribe(
          res => {
            if (res.status === 'success') {
              this.toastr.success('Delete initiated');
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

  viewLogs(pod?: any): void {
    const containers = pod?.spec?.containers;
    const allContainerNames = [];
    containers.forEach(element => {
      allContainerNames.push(element.name);
    });
    const data = {
      pod: pod.metadata?.name,
      namespace: pod.metadata?.namespace,
      restart:
        pod.status?.containerStatuses?.length && pod.status?.containerStatuses[0]?.restartCount
          ? pod.status?.containerStatuses[0]?.restartCount
          : 0,
      allContainers: allContainerNames
    };

    this.dialog.open(K8sPodsContainerLogComponent, {
      height: '90%',
      width: '88%',
      disableClose: true,
      data: data
    });
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

  loadMore(retry?: boolean) {
    this.loadMoreData = true;
    const queryParam = { continue: this.resourceToken };

    const namespace$ = this.namespaceService.selectedNamespace$.pipe(tap(() => {}));
    if (this.searchTerm.length > 0) {
      queryParam['q'] = this.searchTerm;
    }
    if (retry) {
      const limit: number = this.data.length + 10;
      queryParam['limit'] = limit;
      delete queryParam['continue'];
    }
    this.namespaceService
      .getPods(queryParam)
      .pipe(takeWhile(() => this.isAlive))
      .subscribe({
        next: res => {
          if (res?.status === 'error') {
            this.toastr.error(res?.message);
          }
          if (retry) this.data = [];
          this.loadMoreData = false;
          this.remaining = res.data.Remaining;
          this.resourceToken = res.data.Resource || '';
          this.data = this.data.concat(res.data.Result) || [];
        },
        error: err => {
          this.loadMore(true);
        }
      });
  }

  btoa(string) {
    return btoa(string);
  }

  getTerminalUrl(pod: any) {
    return this.namespaceService.getTerminalUrl(pod);
  }
}
