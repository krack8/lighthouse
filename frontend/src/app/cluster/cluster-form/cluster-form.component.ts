import { CommonModule } from '@angular/common';
import { Component, inject, OnInit } from '@angular/core';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { MatTooltipModule } from '@angular/material/tooltip';
import { IconModule } from '@visurel/iconify-angular';

import icClose from '@iconify/icons-ic/close';
import icInfo from '@iconify/icons-ic/info';

import { MAT_DIALOG_DATA, MatDialog, MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatExpansionModule } from '@angular/material/expansion';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { CdkClipboardModule } from '@cdk-ui/clipboard';
import { ClusterService } from '@cluster/cluster.service';
import { ToastrService } from '@sdk-ui/ui';
import { SecureDeleteDialogComponent } from '@shared-ui/ui';

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
  private readonly _clusterService = inject(ClusterService);
  private readonly _dialog = inject(MatDialog);
  private readonly _toastrService = inject(ToastrService);
  private readonly dialogRef = inject<MatDialogRef<ClusterFormComponent>>(MatDialogRef);

  public cluster = inject(MAT_DIALOG_DATA);

  icInfo = icInfo;
  icClose = icClose;

  clusterForm = this._fb.group({
    name: ['', [Validators.required, Validators.minLength(3), Validators.pattern(/^[a-z][a-z0-9-]*(?: [a-z0-9-]+)*$/)]],
  });
  isSubmitting = false;

  isHelmChartLoading = false;
  chartData!: { helm_command: string; helm_repo: string };

  isMasterClusterLoading = false;
  isCreated: boolean = false;

  ngOnInit(): void {
    if (this.cluster) {
      this.getClusterHelmChart();
    }
  }

  createCluster(): void {
    this.isSubmitting = true;
    this._clusterService.createCluster(this.clusterForm.value).subscribe(
      cluster => {
        this.isSubmitting = false;
        this.isCreated = true;
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

  deleteCluster(): void {
    const deleteDialog = this._dialog.open(SecureDeleteDialogComponent, {
      width: '600px',
      minHeight: '350px',
      data: {
        module: 'CLUSTER',
        route: '/clusters',
        id: this.cluster?.id,
        name: this.cluster?.name,
        method: this._clusterService.deleteCluster(this.cluster?.id),
      },
    });
    deleteDialog.afterClosed().subscribe((status: string) => {
      if (status === 'success') {
        this._toastrService.success('Cluster deleted successfully');
        this.dialogRef.close(true);
      }
    });
  }
}
