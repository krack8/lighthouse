import { Component, Inject, OnInit } from '@angular/core';
import { UntypedFormBuilder, UntypedFormGroup, Validators } from '@angular/forms';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import icInfo from '@iconify/icons-ic/info';
import icClose from '@iconify/icons-ic/twotone-close';
import icVisibility from '@iconify/icons-ic/visibility';
import icVisibilityOff from '@iconify/icons-ic/visibility-off';
import { UserService } from '../user.service';
import { MustMatchValidator } from '@shared-ui/validators';
import { ToastrService } from '@sdk-ui/ui';
import { AccessRoleService } from '@management/access-role/access-role.service';
import { PasswordValidator } from '@shared-ui/validators';
import { EmailValidator } from '@shared-ui/validators';
import { SpaceValidator } from '@shared-ui/validators';
import { MatRadioChange } from '@angular/material/radio';
import { ClusterService } from '@cluster/cluster.service';

@Component({
  selector: 'kc-user-form',
  templateUrl: './user-form.component.html',
  styleUrls: ['./user-form.component.scss']
})
export class UserFormComponent implements OnInit {
  icClose = icClose;
  icVisibilityOff = icVisibilityOff;
  icVisibility = icVisibility;
  icInfo = icInfo;

  passwordVisibilityHide: boolean = true;
  enabledPasswordField: boolean = false;
  userForm!: UntypedFormGroup;
  isSubmitting!: boolean;

  roleList: any[] = [];
  isRolesLoading!: boolean;
  searchRoleTerm: string = '';
  clusterList: any[] = [];

  readonly systemRoleUsername: string = 'SYSTEM'; // For Role

  constructor(
    @Inject(MAT_DIALOG_DATA) public data: any,
    public dialogRef: MatDialogRef<UserFormComponent>,
    public toastr: ToastrService,
    private fb: UntypedFormBuilder,
    private _userService: UserService,
    private accessRoleSvc: AccessRoleService,
    private clusterService: ClusterService
  ) {}

  ngOnInit() {
    this.getClustersList();
    this.userForm = this.fb.group(
      {
        first_name: ['', [Validators.required, SpaceValidator.noLeadingSpace]],
        last_name: ['', [Validators.required, SpaceValidator.noLeadingSpace]],
        user_type: ['ADMIN'],
        username: ['', [Validators.required, EmailValidator]],
        password: ['', [Validators.required, Validators.minLength(6), PasswordValidator]],
        passwordConfirm: [''],
        user_is_active: [true]
      },
      {
        validator: MustMatchValidator('password', 'passwordConfirm')
      }
    );
    if (this.data) {
      this.userForm.addControl('user_is_active', this.fb.control(this.data.user_is_active));
      this.userForm.removeControl('password');
      this.userForm.removeControl('passwordConfirm');

      const payload: any = {
        first_name: this.data.first_name,
        last_name: this.data.last_name,
        user_type: this.data.user_type,
        username: this.data.username,
        user_is_active: this.data.user_is_active,
      };

      if(this.data.user_type === 'USER'){
        payload['cluster_ids'] = this.data?.cluster_ids;
        this.userForm.addControl('cluster_ids', this.fb.control([], Validators.required));
      }

      // if (this.data.user_type === 'USER') {
      //   const _roleList = this.data.roles;
      //   let roles = [];
      //   if (_roleList && _roleList.length) {
      //     roles = this.data.roles.map((item: any) => item.id);
      //     payload['role_ids'] = roles;
      //   }
      //   this.userForm.addControl('role_ids', this.fb.control(roles));
      //   this.getRoles();
      // } else {
      //   this.userForm.removeControl('role_ids');
      // }
      this.userForm.patchValue(payload);
      this.userForm.get('username').disable();
    }
  }

  userTypeChange(e: MatRadioChange): void {
    this.userForm.markAsDirty();
    if (e.value === 'ADMIN') {
      this.userForm.removeControl('role_ids');
      this.userForm.removeControl('cluster_ids');
      return;
    }
    this.userForm.addControl('cluster_ids', this.fb.control([], Validators.required));
    // if (!this.roleList?.length) this.getRoles();
  }

  onSubmit(): void {
    this.isSubmitting = true;
    if (this.data) {
      const formData = this.userForm.getRawValue();

      this._userService.mcUpdateUser(this.data.id, formData).subscribe({
        next: _ => {
          this.toastr.success('User Updated.');
          this.dialogRef.close(true);
        },
        error: err => {
          this.toastr.error(err.message || 'Something is wrong!');
          this.isSubmitting = false;
        }
      });
      return;
    }
    // Update User
    const formData = this.userForm.value;
    if (formData.hasOwnProperty('passwordConfirm')) {
      delete formData['passwordConfirm'];
    }
    this._userService.mcCreateUser(formData).subscribe({
      next: _ => {
        this.toastr.success('User Created!');
        this.dialogRef.close(true);
      },
      error: err => {
        this.toastr.error(err.message || 'Something is wrong!');
        this.isSubmitting = false;
      }
    });
  }

  getRoles(): void {
    this.isRolesLoading = true;
    this.accessRoleSvc.getAccessRoles().subscribe({
      next: data => {
        this.roleList = data;
        this.isRolesLoading = false;
        if (!this.data) {
          const systemRole = this.roleList.find(_role => _role.created_by === this.systemRoleUsername);
          if (systemRole.id) {
            this.userForm.get('role_ids').setValue([systemRole.id]);
          }
        }
      },
      error: err => {
        this.isRolesLoading = false;
        this.toastr.error(err.message || 'Something Wrong on fetch roles');
      }
    });
  }

  getClustersList(){
    this.clusterService.getClustersList().subscribe({
      next: data => {
        this.clusterList = data;
        console.log('Cluster List', this.clusterList);
      },
      error: err => {
        this.toastr.error(err.message || 'Something Wrong on fetch clusters');
      }
    });
  }
}
