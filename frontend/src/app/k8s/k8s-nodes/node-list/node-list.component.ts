import { Component, OnDestroy, OnInit, ViewChild } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { ActivatedRoute, Router } from '@angular/router';
import icKeyboardBackspace from '@iconify/icons-ic/keyboard-backspace';
import icSearch from '@iconify/icons-ic/search';
import icAdd from '@iconify/icons-ic/twotone-add';
import icCross from '@iconify/icons-ic/twotone-cancel';
import icLabel from '@iconify/icons-ic/twotone-label';
import icCircle from '@iconify/icons-ic/twotone-lens';
import icMoreHoriz from '@iconify/icons-ic/twotone-more-horiz';
import icRefresh from '@iconify/icons-ic/twotone-refresh';
import icTaints from '@iconify/icons-ic/twotone-shield';
import icUncordon from '@iconify/icons-ic/twotone-timer';
import icCordon from '@iconify/icons-ic/twotone-timer-off';
import { ToolbarService } from '@sdk-ui/services/toolbar.service';
import { ToastrService } from '@sdk-ui/ui';
import { ConfirmDialogComponent } from '@shared-ui/ui';
import { takeWhile } from 'rxjs/operators';
import { K8sNodesService } from '../k8s-nodes.service';
import { NodeTaintDialogComponent } from './node-taint-dialog/node-taint-dialog.component';
import { ApexChart, ApexNonAxisChartSeries, ApexPlotOptions, ChartComponent } from 'ng-apexcharts';

export type ChartOptions = {
 series: ApexNonAxisChartSeries;
 chart: ApexChart;
 labels: string[];
 color: string[];
 plotOptions: ApexPlotOptions;
};

@Component({
 selector: 'kc-node-list',
 templateUrl: './node-list.component.html',
 styleUrls: ['./node-list.component.scss']
})
export class NodeListComponent implements OnInit, OnDestroy {

 //Metrics
 @ViewChild("chart") chart: ChartComponent;
 public chartOptions: Partial<ChartOptions>;

 podChart = {
   series: [],
   labels: ["Usage"],
   color: ['#36c678'],
 }

 cpuChart = {
   series: [],
   labels: ["Usage", "Allocatable",],
   color: ['#36c678','#226add'],
 }

 memoryChart = {
   series: [],
   labels: ["Usage", "Allocatable",],
   color: ['#36c678', '#226add'],
 }

 nodeMetricsMap: Record<string, any>

 graphStats: any;
 cpuUsagePercentage: number =0;
 CpuAllocatablePercentage:  number = 0;
 MemoryUsagePercentage: number = 0;
 MemoryAllocatablePercentage: number = 0;
 podUsagePercentage: number = 0;
 statsLoaded : boolean = false;
  //icons
 toolTip = true;
 icCircle = icCircle;
 icKeyboardBackspace = icKeyboardBackspace;
 icSearch = icSearch;
 icMoreHoriz = icMoreHoriz;
 icLabel = icLabel;
 icTaints = icTaints;
 icAdd = icAdd;
 icUncordon = icUncordon;
 icCordon = icCordon;
 icCross = icCross;
 icRefresh = icRefresh;

 searchBy = 'name';
 isAlive: boolean = true;
 nodeList: any[] = [];
 isLoading!: boolean;
 metricsAvailable: boolean = false;
 searchTerm = '';


 constructor(
   private nodeService: K8sNodesService,
   private router: Router,
   private route: ActivatedRoute,
   private toolbarService: ToolbarService,
   public dialog: MatDialog,
   private snackBar: MatSnackBar,
   private toastr: ToastrService
 ) {
   this.chartOptions = {
     chart: {
       height: 200,
       width: 200,
       type: "radialBar"
     },
     plotOptions: {
       radialBar: {
         track: {
           background: '#3c4a61' // Set the track background color
         },
        
         dataLabels: {
           name: {
             fontSize: "15px"
           },
           value: {
             fontSize: "12px",
             color: "#5bc4d6",
           },
           total: {
             show: false,
             label: "Total",
             formatter: function(w) {
               return "249";
             }
           },
         }
       }
      
     },
   };
 }


 ngOnInit(): void {
   this.toolbarService.changeData({ title: 'Nodes' });
   this.getNodeList();
 }


 ngOnDestroy(): void {
   this.isAlive = false;
 }


 reloadList() {
   this.getNodeList();
 }


 percentageCalculator(used, total) {
   return Math.round((used / total) * 100);
 }


 getNodeList(queryParam?: any): void {
   this.isLoading = true;
   this.nodeService
     .getNodes(queryParam)
     .pipe(takeWhile(() => this.isAlive))
     .subscribe({
       next: data => {
         if (data.status === 'success') {
           this.isLoading = false;
           if (data.data.Metrics.length === 0) {
             this.nodeList = data.data.Result;
           } else {
             this.nodeList = data?.data?.Result || [];
             this.nodeMetricsMap = {};
             data.data?.Metrics.forEach((item) => {
               if (item?.usage) {
                 this.nodeMetricsMap[item.metadata.name] = {
                   cpu: item.usage?.cpu,
                   memory: item.usage?.memory,
                 };
               }
             });
             this.graphStats = data.data.graph_view;
             this.initGraphStats();
             this.metricsAvailable = true;
           }
         } else {
           this.isLoading = false;
         }
       },
       error: err => {
         this.isLoading = false;
         this.toastr.error('Failed', err.message);
       }
     });
 }


 onNodeDetailsClick(item): void {
   this.router.navigate(['node-details'], {
     queryParams: {
       ...this.route.snapshot.queryParams,
       name: item.metadata.name
     },
     relativeTo: this.route
   });
 }


 convertKbToGigabyte(size: string): number {
   const kbSize = parseInt(size.replace(/\D/g, ''));
   const mbSize = kbSize / 1000000;
   return Number(mbSize.toFixed(2));
 }


 convertCpuValueToBase(value): Number {
   const unit = value.replace(/\d/g, '');
   const cpu = value.replace(/\D/g, '');
   if (unit === 'n') {
     const val = cpu / 1_000_000_000;
     return Number(val.toFixed(2));
   }
   if (unit === 'm') {
     const val = cpu / 1_000;
     return Number(val.toFixed(2));
   }
   if (unit === '') return Number(cpu);
 }


 onSearch() {
   if (this.searchBy === 'label') {
     const keyValuePairs = this.searchTerm.split(',');
     const jsonObject = {};
     keyValuePairs.forEach(pair => {
       const [key, value] = pair.split(':');
       jsonObject[key] = value;
     });
     const jsonString = JSON.stringify(jsonObject);
     const qp = { labels: jsonString };
     this.getNodeList(qp);
   }
   if (this.searchBy === 'name') {
     const qp = { q: this.searchTerm };
     this.getNodeList(qp);
   }
 }
 clearSearch() {
   this.getNodeList();
   this.searchTerm = '';
 }
 handleInputChange() {
   if (this.searchTerm.length === 0) {
     this.getNodeList();
   }
 }


 nodeCordonUncordon(nodeName: string, action: string) {
   const dialogRef = this.dialog.open(ConfirmDialogComponent, {
     width: '400px',
     data: `Are you sure! you want to ${action} ${nodeName}?`
   });


   dialogRef.afterClosed().subscribe(
     (result: any) => {
       if (result) {
         this.isLoading = true;
         this.nodeService.nodeCordonUncordon(nodeName).subscribe(
           res => {
             if (res.status === 'success') {
               this.isLoading = false;
               this.toastr.success('Node ' + action + 'ed successfully');
               this.getNodeList();
             }
           },
           err => {
             this.getNodeList();
             this.isLoading = false;
             this.toastr.error('Node ' + action + ' was unsuccessful');
           }
         );
       }
     },
     err => {
       this.snackBar.open(err.message, 'close', { duration: 2000 });
     }
   );
 }


 openTaintDialog(node: any): void {
   const dialogRef = this.dialog.open(NodeTaintDialogComponent, {
     minHeight: '270px',
     width: '720px',
     disableClose: true,
     data: node
   });


   dialogRef.afterClosed().subscribe(res => {
     if (res) {
       this.getNodeList();
     }
   });
 }


 getStatusDescriptions(item: any) {
   for (let i = item.status.conditions.length - 1; i >= 0; i--) {
     let status = item.status.conditions[i];
     if (status.type === 'Ready') {
       if (status.status === 'True') {
         let description = status.type;
         if (item?.spec?.unschedulable) {
           description += ', Scheduling Disabled';
         }
         return description;
       } else {
         return 'Not Ready';
       }
     }
   }
 }

 initGraphStats(){
   this.podChart.series = [];
   this.memoryChart.series = [];
   this.cpuChart.series = [];
   // Pod Usage
   this.podUsagePercentage = this.percentageCalculator(this.graphStats.deployed_pod_count, this.graphStats.pod_capacity);
   this.podChart.series = [this.podUsagePercentage];
    // CPU Usage and Allocatable
   this.cpuUsagePercentage = this.percentageCalculator(this.graphStats.node_cpu_usage, this.graphStats.node_cpu_capacity);
   this.cpuChart.series.push(this.cpuUsagePercentage);
   this.CpuAllocatablePercentage = this.percentageCalculator(this.graphStats.node_cpu_allocatable, this.graphStats.node_cpu_capacity);
   this.cpuChart.series.push(this.CpuAllocatablePercentage);
    // Memory Usage and Allocatable
   this.MemoryUsagePercentage = this.percentageCalculator(this.graphStats.node_memory_usage, this.graphStats.node_memory_capacity);
   this.memoryChart.series.push(this.MemoryUsagePercentage);
   this.MemoryAllocatablePercentage = this.percentageCalculator(this.graphStats.node_memory_allocatable, this.graphStats.node_memory_capacity);
   this.memoryChart.series.push(this.MemoryAllocatablePercentage);

   this.statsLoaded = true;
 }

 isNodeReady(item: any): boolean {
  if (!item?.status?.conditions) {
    return false; // Return false if conditions are not available
  }

  // Find the "Ready" condition
  const readyCondition = item.status.conditions.find((condition: any) => condition.type === 'Ready');

  // Check if the "Ready" condition exists and its status is "True"
  if (readyCondition?.status === 'True') {
    // Check if the node is schedulable
    return !item?.spec?.unschedulable; // Return true if schedulable, false otherwise
  }

  return false; // Return false if the "Ready" condition is not "True"
}
}




