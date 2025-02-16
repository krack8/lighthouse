import { Component, Inject, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { NgIf } from '@angular/common';
import { MatDialogRef, MAT_DIALOG_DATA, MatDialogModule } from '@angular/material/dialog';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { ToastrService, ToastrModule } from '@sdk-ui/ui';
import { FlexLayoutModule } from '@angular/flex-layout';
import { MatButtonModule } from '@angular/material/button';

@Component({
  standalone: true,
  imports: [NgIf, MatButtonModule, MatDialogModule, MatProgressBarModule, ToastrModule, FlexLayoutModule],
  selector: 'kc-secure-delete-dialog',
  templateUrl: './secure-delete-dialog.component.html',
  styleUrls: ['./secure-delete-dialog.component.scss']
})
export class SecureDeleteDialogComponent implements OnInit {
  fConfirm = false;
  isSubmitting!: boolean;
  invalidInp = true;

  constructor(
    @Inject(MAT_DIALOG_DATA) public data: any,
    public dialogRef: MatDialogRef<SecureDeleteDialogComponent>,
    private _router: Router,
    private toastr: ToastrService
  ) {}

  ngOnInit() {}

  confirm() {
    this.fConfirm = true;
  }

  onChange(value: string) {
    this.validate(value);
  }

  validate(name: string) {
    if (name === this.data.name) {
      this.invalidInp = false;
    } else {
      this.invalidInp = true;
    }
  }

  delete(name: string) {
    this.isSubmitting = true;
    if (name !== this.data.name) {
      this.isSubmitting = false;
      this.toastr.error('Name is Invalid', 'Validation Error');
    } else {
      this.data.method.subscribe({
        next: (res: any) => {
          if (res.status === 'success' || res.status === 'OK') {
            this.isSubmitting = false;
            this.toastr.success(res.message || this.data.module + ' deleted successfully', this.data.successTitle || 'Success');

            if (this.data.route) this._router.navigate([this.data.route]);
          }

          if (res.status === 'error') {
            this.toastr.error(res.message || 'Cannot delete' + this.data.module);
          }
          this.dialogRef.close(res.status || "success");
        },
        error: (err: any) => {
          this.isSubmitting = false;
          this.toastr.error(err.error.message || 'something is wrong');
          console.log(err);
          throw err;
        }
      });
    }
  }
}
