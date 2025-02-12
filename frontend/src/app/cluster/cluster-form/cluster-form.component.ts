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

import { PageLayoutModule, ToastrService } from '@sdk-ui/ui';
import { ToolbarService } from '@sdk-ui/services';
import { ClusterService } from '@cluster/cluster.service';
import { Router } from '@angular/router';

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
    IconModule,
    PageLayoutModule
  ],
  templateUrl: './cluster-form.component.html',
  styleUrls: ['./cluster-form.component.scss']
})
export class ClusterFormComponent implements OnInit {
  private readonly _fb = inject(FormBuilder);
  private readonly _router = inject(Router);
  private readonly _toolberService = inject(ToolbarService);
  private readonly _clusterService = inject(ClusterService);
  private readonly _toastrService = inject(ToastrService);

  icInfo = icInfo;

  clusterForm = this._fb.group({
    name: ['', [Validators.required, Validators.minLength(4), Validators.pattern(/^[a-z][a-z0-9-]+$/)]],
    namespace: ['klovercloud', [Validators.minLength(4), Validators.pattern(/^[a-z][a-z0-9-]+$/)]]
  });
  isSubmitting = false;

  ngOnInit(): void {
    this._toolberService.changeData({ title: 'Crate Agent Cluster' });
  }

  createCluster(): void {
    this.isSubmitting = true;
    this._clusterService.createCluster(this.clusterForm.value).subscribe(
      cluster => {
        this.isSubmitting = false
        this._router.navigate(['/clusters', cluster.id, 'init'])
      },
      err => {
        this.isSubmitting = false
        this._toastrService.error(err.message)
      }
    );
  }
}
