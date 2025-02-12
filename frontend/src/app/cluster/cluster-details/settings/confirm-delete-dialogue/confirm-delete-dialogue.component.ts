import { Component, Inject, OnInit } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { ToastrService } from '@sdk-ui/ui';
import { fadeInUp400ms } from '@sdk-ui/animations/fade-in-up.animation';
import { Utils } from '@shared-ui/utils';

@Component({
  selector: 'kc-confirm-delete-dialogue',
  templateUrl: './confirm-delete-dialogue.component.html',
  styleUrls: ['./confirm-delete-dialogue.component.scss'],
  animations: [fadeInUp400ms]
})
export class ConfirmDeleteDialogueComponent {
  fConfirm = false;
  isSubmitting = false;
  invalidInp = true;

  constructor(
    @Inject(MAT_DIALOG_DATA) public data,
    public dialogRef: MatDialogRef<ConfirmDeleteDialogueComponent>,
    private _router: Router,
    private toastr: ToastrService
  ) {}

  confirm() {
    this.fConfirm = true;
  }

  onChange(value: string) {
    this.validate(value);
  }

  validate(name) {
    if (name === this.data.name) {
      this.invalidInp = false;
    } else {
      this.invalidInp = true;
    }
  }

  delete(name) {
    this.isSubmitting = true;
    if (name !== this.data.name) {
      this.isSubmitting = false;
      this.toastr.error('Name is Invalid', 'Validation Error');
    } else {
      this.data.method.subscribe({
        next: res => {
          if (res.status === 'success' || res.status === 'OK') {
            this.isSubmitting = false;
            this.toastr.success(res.message || this.data.module + ' deleted successfully', this.data.successTitle || 'Success');
            this._router.navigate([this.data.route]);
          }

          if (res.status === 'error') {
            this.toastr.error(res.message || 'Cannot delete' + this.data.module);
          }
          this.dialogRef.close(res.status);
        },
        error: err => {
          this.isSubmitting = false;
          this.toastr.error(err.error.message || 'something is wrong');
          console.log(err);
          throw err;
        }
      });
    }
  }
}
