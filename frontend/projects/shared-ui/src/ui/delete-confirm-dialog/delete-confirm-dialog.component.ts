import { Component, Inject } from '@angular/core';
import { NgIf } from '@angular/common';
import { Observable } from 'rxjs';
import { MatDialogRef, MAT_DIALOG_DATA, MatDialogModule } from '@angular/material/dialog';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { MatButtonModule } from '@angular/material/button';

import { IDeleteConfirmConfig } from './delete-config-dialog.interfaces';
import { ToastrService, ToastrModule } from '@sdk-ui/ui';

@Component({
  standalone: true,
  imports: [NgIf, MatDialogModule, MatButtonModule, MatProgressBarModule, ToastrModule],
  selector: 'kc-confirm-dialog',
  templateUrl: './delete-confirm-dialog.component.html',
  styleUrls: ['./delete-confirm-dialog.component.scss']
})
export class DeleteConfirmDialogComponent {
  icon: string = '/assets/img/bin.svg';
  deleteApi$!: Observable<unknown> | undefined;
  isSubmitting: boolean = false;

  constructor(
    @Inject(MAT_DIALOG_DATA) public config: IDeleteConfirmConfig,
    public dialogRef: MatDialogRef<DeleteConfirmDialogComponent>,
    private toastrService: ToastrService
  ) {}

  onClick() {
    if (this.deleteApi$ === undefined) {
      this.dialogRef.close(true);
    } else {
      this.isSubmitting = true;
      this.deleteApi$.subscribe({
        next: (res: any) => {
          if (res.status === 'success') {
            this.toastrService.success(res.message || 'Deleted successfully');
            this.dialogRef.close(res.data);
          } else {
            this.toastrService.error(res.message);
          }
          this.isSubmitting = false;
        },
        error: err => {
          console.log(err);
          this.toastrService.error(err.error?.message || err.message);
          this.isSubmitting = false;
        }
      });
    }
  }
}
