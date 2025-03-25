import { Component, Inject, OnInit } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { ClusterService } from '@cluster/cluster.service';
import icClose from '@iconify/icons-ic/twotone-close';
import { ToastrService } from '@sdk-ui/ui';

@Component({
  selector: 'kc-cluster-rename-dialog',
  templateUrl: './cluster-rename-dialog.component.html',
  styleUrls: ['./cluster-rename-dialog.component.scss']
})
export class ClusterRenameDialogComponent implements OnInit {
  icClose = icClose;
  clusterName: string = '';
  isSubmitting: boolean = false;

  constructor(
    @Inject(MAT_DIALOG_DATA) public data: any,
    public dialogRef: MatDialogRef<ClusterRenameDialogComponent>,
    private _clusterService: ClusterService,
    private toastr: ToastrService

  ) { }

  ngOnInit(): void {
    this.clusterName = this.data?.name || '';
  }

  renameCluster(): void {
    this.isSubmitting = true;
    if (this.clusterName.length <3){
      this.toastr.error('Cluster name cannot be less than 3 characters');
      this.isSubmitting = false;
      return;
    }
    const payload = { name: this.clusterName };
    this._clusterService.renameCluster(this.data.id, payload).subscribe({
      next: (res) => {
        this.toastr.success('Cluster renamed successfully');
        this.dialogRef.close('success');
      },
      error: (err) => {
        this.toastr.error(err);
        this.isSubmitting = false;
      }
    }
    );
  }

}
