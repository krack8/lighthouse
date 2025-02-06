import { Component, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { K8sNodesService } from '../k8s-nodes.service';
import icArrowBack from '@iconify/icons-ic/arrow-back';
import { ToastrService } from '@sdk-ui/ui';
import icCircle from '@iconify/icons-ic/twotone-lens';
import { ToolbarService } from '@sdk-ui/services';

@Component({
  selector: 'kc-node-details',
  templateUrl: './node-details.component.html',
  styleUrls: ['./node-details.component.scss']
})
export class NodeDetailsComponent implements OnInit {
  details: any;
  metrics: any;
  queryParams: any;
  isLoading: boolean = false;
  icArrowBack = icArrowBack;
  icCircle = icCircle;

  constructor(
    private route: ActivatedRoute,
    private nodeService: K8sNodesService,
    private toastr: ToastrService,
    private toolbarService: ToolbarService
  ) {}

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
          this.metrics = data.data?.Metrics;
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
}
