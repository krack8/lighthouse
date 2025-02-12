import { ChangeDetectorRef, Component, Inject, OnInit } from '@angular/core';
import { ReactiveFormsModule, UntypedFormBuilder, UntypedFormGroup, Validators } from '@angular/forms';
import { MAT_DIALOG_DATA, MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { ICoreConfig } from '@core-ui/services/core-config/core-config.interfaces';
import { CoreConfigService } from '@core-ui/services/core-config/core-config.service';
import { MustMatchValidator } from '@shared-ui/validators';
import { PasswordValidator } from '@shared-ui/validators';
import icInfo from '@iconify/icons-ic/info';
import icClose from '@iconify/icons-ic/twotone-close';
import icVisibility from '@iconify/icons-ic/visibility';
import icVisibilityOff from '@iconify/icons-ic/visibility-off';
import { UserService } from '../user.service';
import { ToastrService } from '@sdk-ui/ui';
import { CommonModule } from '@angular/common';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatIconModule } from '@angular/material/icon';
import { IconModule } from '@visurel/iconify-angular';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatButtonModule } from '@angular/material/button';
import { RequesterService } from '@core-ui/services';

@Component({
  selector: 'kc-update-password',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatIconModule,
    MatDialogModule,
    MatProgressSpinnerModule,
    MatButtonModule,
    IconModule
  ],
  templateUrl: './update-password.component.html',
  styleUrls: ['./update-password.component.scss']
})
export class UpdatePasswordComponent implements OnInit {
  icClose = icClose;
  icVisibilityOff = icVisibilityOff;
  icVisibility = icVisibility;
  icInfo = icInfo;
  userForm!: UntypedFormGroup;
  coreConfig!: ICoreConfig;
  passwordVisibilityHide: boolean = true;
  isSubmitting!: boolean;
  inputType = 'password';
  visible = false;
  userData: any;
  isAdmin: boolean = false;

  constructor(
    @Inject(MAT_DIALOG_DATA) public data: any,
    private fb: UntypedFormBuilder,
    private coreConfigService: CoreConfigService,
    private _userService: UserService,
    public toastr: ToastrService,
    private cd: ChangeDetectorRef,
    public requester: RequesterService,
    public dialogRef: MatDialogRef<UpdatePasswordComponent>
  ) {}

  ngOnInit(): void {
    this.userData = this.requester.get();
    this.coreConfig = this.coreConfigService.generalInfoSnapshot;
    this.userForm = this.fb.group(
      {
        currentPassword: ['', [Validators.required, Validators.minLength(this.coreConfig?.passwordLength || 8), PasswordValidator]],
        newPassword: ['', [Validators.required, Validators.minLength(this.coreConfig?.passwordLength || 8), PasswordValidator]],
        passwordConfirm: ['', Validators.required]
      },
      {
        validator: MustMatchValidator('newPassword', 'passwordConfirm')
      }
    );
    if (this.userData?.userInfo?.user_type === 'ADMIN') {
      this.userForm.removeControl('currentPassword');
    }
  }
  toggleVisibility() {
    if (this.visible) {
      this.inputType = 'password';
      this.visible = false;
      this.cd.detectChanges();
    } else {
      this.inputType = 'text';
      this.visible = true;
      this.cd.detectChanges();
    }
  }

  onSubmit() {
    this.isSubmitting = true;
    const formData = this.userForm.getRawValue();
    const payload = {
      currentPassword: formData?.currentPassword,
      newPassword: formData?.newPassword
    }
    this._userService.mcResetUserPassword(this.data?.id, payload).subscribe({
      next: _ => {
        this.toastr.success('Password updated successfully!');
        this.dialogRef.close(true);
      },
      error: err => {
        this.toastr.error(err.message || 'Something is wrong!');
        this.isSubmitting = false;
      }
    });
  }
}
