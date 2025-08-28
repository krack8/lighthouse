import { Component, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { K8sNodesService } from '../k8s-nodes.service';
import icArrowBack from '@iconify/icons-ic/arrow-back';
import { ToastrService } from '@sdk-ui/ui';
import icCircle from '@iconify/icons-ic/twotone-lens';
import { ToolbarService } from '@sdk-ui/services';
import { ApexChart, ApexNonAxisChartSeries, ApexPlotOptions, ChartComponent } from 'ng-apexcharts';

export type ChartOptions = {
  series: ApexNonAxisChartSeries;
  chart: ApexChart;
  labels: string[];
  color: string[];
  plotOptions: ApexPlotOptions;
};

@Component({
  selector: 'kc-node-details',
  templateUrl: './node-details.component.html',
  styleUrls: ['./node-details.component.scss']
})
export class NodeDetailsComponent implements OnInit {

  // metrics
  @ViewChild("chart") chart: ChartComponent;
  public chartOptions: Partial<ChartOptions>;

  cpuUsagePercentage: number =0;
  CpuAllocatablePercentage:  number = 0;
  MemoryUsagePercentage: number = 0;
  MemoryAllocatablePercentage: number = 0;
  podUsagePercentage: number = 0;
  graphStats: any;
  podChart = {
    series: [],
    labels: ["Usage"],
    color: ['#36c678'],
  }

  cpuChart = {
    series: [],
    labels: ["Usage", "Capacity", "Allocatable",],
    color: ['#36c678', '#226add','#5bc4d6'],
  }

  memoryChart = {
    series: [],
    labels: ["Usage", "Capacity", "Allocatable",],
    color: ['#36c678', '#226add','#5bc4d6'],
  }

  details: any;
  metrics: any;
  queryParams: any;
  isLoading: boolean = false;
  icArrowBack = icArrowBack;
  icCircle = icCircle;

  unitMap = {
    'cpu': 'core',
    'memory': 'Gi',
    'pods': '',
    'ephemeral-storage': 'Gi'
  };

  constructor(
    private route: ActivatedRoute,
    private nodeService: K8sNodesService,
    private toastr: ToastrService,
    private toolbarService: ToolbarService
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
    this.queryParams = this.route.snapshot.queryParams;
    this.getDetails();
  }

  getDetails(): void {
    this.isLoading = true;
    this.nodeService.getNodeDetails(this.queryParams.name).subscribe({
      next: data => {
        if (data?.status === 'success') {
          this.details = data.data?.Result;
          if (data?.data?.Metrics && data?.data?.Metrics !== null) {
            this.metrics = data.data?.Metrics;

            this.graphStats = {
              deployed_pod_count: data.data?.deployed_pod_count,
              pod_capacity: Number(this.details?.status?.capacity?.pods),
              node_cpu_allocatable: this.convertCpuValueToBase(this.details?.status?.allocatable?.cpu),
              node_cpu_capacity: Number(this.details?.status?.capacity?.cpu),
              node_cpu_usage: this.convertCpuValueToBase(this.metrics?.usage?.cpu),
              node_memory_allocatable: this.convertKbToGigabyte(this.details?.status?.allocatable?.memory),
              node_memory_capacity: this.convertKbToGigabyte(this.details?.status?.capacity?.memory),
              node_memory_usage: this.convertKbToGigabyte(this.metrics?.usage?.memory)
            }

            this.initChart();
          }
          
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

  initChart() {
    this.podUsagePercentage = this.percentageCalculator(this.graphStats.deployed_pod_count, this.graphStats.pod_capacity);
    this.podChart.series = [this.podUsagePercentage];

    this.cpuUsagePercentage = this.percentageCalculator(this.graphStats.node_cpu_usage, this.graphStats.node_cpu_capacity);
    this.cpuChart.series.push(this.cpuUsagePercentage);
    this.CpuAllocatablePercentage = this.percentageCalculator(this.graphStats.node_cpu_allocatable, this.graphStats.node_cpu_capacity);
    this.cpuChart.series.push(this.CpuAllocatablePercentage);

    this.MemoryUsagePercentage = this.percentageCalculator(this.graphStats.node_memory_usage, this.graphStats.node_memory_capacity);
    this.memoryChart.series.push(this.MemoryUsagePercentage);
    this.MemoryAllocatablePercentage = this.percentageCalculator(this.graphStats.node_memory_allocatable, this.graphStats.node_memory_capacity);
    this.memoryChart.series.push(this.MemoryAllocatablePercentage);

    this.isLoading = false;
  }

  percentageCalculator(used, total) {
    if (used && total) {
      return Math.round((used / total) * 100);
    }
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

  isConditionNegative(condition): boolean {
    const type = condition.type;
    const status = condition.status;
    const types = ['Progressing', 'Available'];
    if (types.includes(type) && status === 'True') {
      return false;
    }
    return true;
  }

  convertUnitToBase(value: any, key?: any): any {
    const unit = value.replace(/\d/g, '');
    const size = value.replace(/\D/g, '');
    if (unit === 'Ki') {
      return Number((parseInt(size) / 1_000_000).toFixed(2)) + ' Gi'; // Convert KiB to GiB
    } else if (unit === 'Mi') {
      return Number((parseInt(size) / 1_000).toFixed(2)) + ' Gi'; // Convert MiB to GiB
    }  else if (unit === 'n') {
      return Number((parseInt(size) / 1_000_000_000).toFixed(2)) + ' core'; // Convert nano core to core
    } else if (unit === 'm') {
      return Number((parseInt(size) / 1_000).toFixed(2)) + ' core'; // Convert milli core to core
    } else if (key && key === 'ephemeral-storage' && unit == '') {
      return Number((parseInt(size) / 1_073_741_824).toFixed(2)) + ' Gi'; // Convert bytes to GiB
    }
    if(unit == '' && key && this.unitMap[key]) {
      return parseFloat(size) + ' ' + this.unitMap[key]; // if no unit is available from data, Return the value with the unit from unitMap 
    }
    return parseFloat(size); // Fallback for other units or plain numbers
  }

  //byte to gigabyte
  convertBytesToGigabytes(bytes: number): number {
    return Number((bytes / (1024 * 1024 * 1024)).toFixed(2));
  }
}
