import { Component, OnInit } from '@angular/core';
import { ReactiveFormsModule, UntypedFormBuilder, UntypedFormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { RequesterService } from '@core-ui/services/requester.service';
import { Subject } from 'rxjs';
import { takeUntil } from 'rxjs/operators';
import icVisibility from '@iconify/icons-ic/visibility';
import icVisibilityOff from '@iconify/icons-ic/visibility-off';
import { PasswordValidator, MustMatchValidator } from '@shared-ui/validators';
import { ICoreConfig } from '@core-ui/services/core-config/core-config.interfaces';
import { CoreConfigService } from '@core-ui/services/core-config/core-config.service';
import { CommonModule } from '@angular/common';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatIconModule } from '@angular/material/icon';
import { IconModule } from '@visurel/iconify-angular';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { MatButtonModule } from '@angular/material/button';
import { MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import icClose from '@iconify/icons-ic/close';
import { UserProfileService } from '../user-profile.service';

@Component({
  selector: 'kc-change-password',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatIconModule,
    MatProgressBarModule,
    MatButtonModule,
    MatDialogModule,
    IconModule
  ],
  templateUrl: './change-password.component.html',
  styleUrls: ['./change-password.component.scss']
})
export class ChangePasswordComponent implements OnInit {
  icVisibilityOff = icVisibilityOff;
  icVisibility = icVisibility;
  icClose = icClose;

  resetPassForm: UntypedFormGroup;
  err = false;
  isSubmitting = false;

  userData: any;

  hidePassword: boolean = true;
  hideNewPassword: boolean = true;

  coreConfig!: ICoreConfig;

  constructor(
    private formBuilder: UntypedFormBuilder,
    private snackBar: MatSnackBar,
    private requesterService: RequesterService,
    private coreConfigService: CoreConfigService,
    private dialogRef: MatDialogRef<ChangePasswordComponent>,
    private userProfileService: UserProfileService
  ) {}

  ngOnInit() {
    this.coreConfig = this.coreConfigService.generalInfoSnapshot;
    this.userData = this.requesterService.get();
    this.resetPassForm = this.formBuilder.group(
      {
        password1: ['', [Validators.required, Validators.minLength(this.coreConfig?.passwordLength || 6), PasswordValidator]],
        password2: ['', [Validators.required]],
        password: ['', [Validators.required, Validators.minLength(this.coreConfig?.passwordLength || 6)]]
      },
      { validator: MustMatchValidator('password1', 'password2') }
    );
  }

  get password() {
    return this.resetPassForm.get('password');
  }
  get password1() {
    return this.resetPassForm.get('password1');
  }
  get password2() {
    return this.resetPassForm.get('password2');
  }

  submit() {
    if (this.resetPassForm.invalid) {
      this.resetPassForm.markAllAsTouched();
      return;
    }
    this.isSubmitting = true;
    const id = this.userData?.userInfo?.id;

    const payload = {
      currentPassword: this.password.value,
      newPassword: this.password2.value
    };

    this.userProfileService.resetPassword(id, payload).subscribe(
      res => {
        this.snackBar.open(res.message, 'close', { duration: 5000 });
        this.isSubmitting = false;
        this.dialogRef.close(res);
      },
      error => {
        this.snackBar.open(error?.message, 'close', { duration: 5000 });
        this.isSubmitting = false;
      }
    );
  }
}
