import { Component, inject, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatIconModule } from '@angular/material/icon';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { MatButtonModule } from '@angular/material/button';
import { IconModule } from '@visurel/iconify-angular';

import icInfo from '@iconify/icons-ic/info';
import icClose from '@iconify/icons-ic/close';

import { ToastrService } from '@sdk-ui/ui';
import { ClusterService } from '@cluster/cluster.service';
import { Router } from '@angular/router';
import { MAT_DIALOG_DATA, MatDialogModule } from '@angular/material/dialog';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatExpansionModule } from '@angular/material/expansion';
import { CdkClipboardModule } from '@cdk-ui/clipboard';

@Component({
  selector: 'kc-cluster-form',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatTooltipModule,
    MatProgressBarModule,
    MatIconModule,
    MatDialogModule,
    MatProgressSpinnerModule,
    MatExpansionModule,
    CdkClipboardModule,
    IconModule
  ],
  templateUrl: './cluster-form.component.html',
  styleUrls: ['./cluster-form.component.scss']
})
export class ClusterFormComponent implements OnInit {
  private readonly _fb = inject(FormBuilder);
  private readonly _router = inject(Router);
  private readonly _clusterService = inject(ClusterService);
  private readonly _toastrService = inject(ToastrService);

  public cluster = inject(MAT_DIALOG_DATA);

  icInfo = icInfo;
  icClose = icClose;

  clusterForm = this._fb.group({
    name: ['', [Validators.required, Validators.minLength(3), Validators.pattern(/^[a-z][a-z0-9-]+$/)]],
    masterClusterId: ['', Validators.required]
  });
  isSubmitting = false;

  isHelmChartLoading = false;
  chartData!: { helm_command: string; helm_repo: string };

  isMasterClusterLoading = false;

  ngOnInit(): void {
    if (this.cluster) {
      this.getClusterHelmChart();
    } else {
      this.getMasterCluster();
    }
  }

  createCluster(): void {
    this.isSubmitting = true;
    this._clusterService.createCluster(this.clusterForm.value).subscribe(
      cluster => {
        this.isSubmitting = false;
        this.cluster = cluster;
        this.cluster['cluster_status'] = 'PENDING';
        this.getClusterHelmChart();
      },
      err => {
        this.isSubmitting = false;
        this._toastrService.error(err.message);
      }
    );
  }

  getMasterCluster(): void {
    this.isMasterClusterLoading = true;
    this._clusterService.getMasterCluster().subscribe(
      _cluster => {
        this.isMasterClusterLoading = false;
        this.clusterForm.get('masterClusterId').patchValue(_cluster.id);
      },
      err => {
        this.isMasterClusterLoading = false;
        this._toastrService.error(err.message);
      }
    );
  }

  // Helm Chart
  getClusterHelmChart(): void {
    this.isHelmChartLoading = true;
    this._clusterService.getHelmChart(this.cluster.id).subscribe(
      data => {
        this.chartData = data;
        this.isHelmChartLoading = false;
      },
      err => {
        this.isHelmChartLoading = false;
        this._toastrService.error(err.message);
      }
    );
  }
  getProcessedHelmCommand(): string {
    if (this.chartData && this.chartData.helm_command)
      return this.chartData.helm_command
        .split('\n')
        .map(line => line.trim().replace(/(\s+)--set/g, '--set'))
        .join('\n');
    return 'N/A';
  }
}
