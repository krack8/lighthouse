import { Component, Inject, Input, Optional } from '@angular/core';
import { NgClass, NgIf, NgStyle } from '@angular/common';
import { MatDialogRef, MAT_DIALOG_DATA, MatDialogModule } from '@angular/material/dialog';
import { Observable } from 'rxjs';
import { MatButtonModule } from '@angular/material/button';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { IDialogStaticConfig } from './config-dialog-static.interfaces';
import { ToastrModule, ToastrService } from '@sdk-ui/ui';

/**
 * @description ConfirmDialogStatic dialog no long static, it's now dynamic dialog and cann pass Api observable.
 * @Input actionApi$!: Observable<unknown> | undefined;
 * @Input isSubmitting: boolean = false;
 * @MAT_DIALOG_DATA Interface IDialogStaticConfig (props: type, message, note, noteStyle, successBtnText, cancelBtnText, icon)
 */
@Component({
  standalone: true,
  imports: [NgClass, NgIf, NgStyle, MatDialogModule, MatButtonModule, MatProgressBarModule, MatDialogModule, ToastrModule],
  selector: 'kc-confirm-dialog-static',
  templateUrl: './confirm-dialog-static.component.html',
  styleUrls: ['./confirm-dialog-static.component.scss']
})
export class ConfirmDialogStaticComponent {
  // Success and Api Submit
  @Input() actionApi$!: Observable<unknown> | undefined;
  @Input() isSubmitting: boolean = false;

  // dialog Config
  config: IDialogStaticConfig = {
    type: 'info',

    // Static Props
    message: 'Are you sure!',
    successBtnText: 'Okay',
    cancelBtnText: 'Cancel',
    icon: 'assets/img/icons/airdrop.svg'
  };

  constructor(
    @Inject(MAT_DIALOG_DATA) private configData: Partial<IDialogStaticConfig>,
    private dialogRef: MatDialogRef<ConfirmDialogStaticComponent>,
    private toastrService: ToastrService
  ) {
    const { successBtnText, cancelBtnText, icon, message, type, note, noteStyle, cancelBtnValue } = configData;
    this.config.message = message as string;
    this.config.cancelBtnValue = cancelBtnValue;

    if (successBtnText !== undefined) this.config.successBtnText = successBtnText;
    if (note) this.config.note = note;
    if (noteStyle) this.config.noteStyle = noteStyle;
    if (icon !== undefined) this.config.icon = icon;
    if (cancelBtnText !== undefined) this.config.cancelBtnText = cancelBtnText;

    if (type !== undefined) {
      this.config.type = type;
      if (type === 'delete') {
        if (successBtnText === undefined) this.config.successBtnText = 'Delete';
        if (icon === undefined) this.config.icon = '/assets/img/bin.svg';
      }
    }
  }

  onSuccess() {
    if (this.actionApi$ === undefined) {
      this.dialogRef.close(true);
    } else {
      this.isSubmitting = true;
      this.actionApi$.subscribe({
        next: (res: any) => {
          if (res.status === 'success' || (res.status >= 200 && res.status < 300)) {
            this.toastrService.success(res.message || 'Action completed successfully');
            this.dialogRef.close(res?.data || true);
          } else {
            this.toastrService.error(res.message);
          }
          this.isSubmitting = false;
        },
        error: err => {
          this.toastrService.error(err.error?.message || err.message);
          this.isSubmitting = false;
        }
      });
    }
  }
}
