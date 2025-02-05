import { Component, inject, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import icClose from '@iconify/icons-ic/close';
import icInfo from '@iconify/icons-ic/info';

import { MAT_DIALOG_DATA, MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatSelectModule } from '@angular/material/select';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { IconModule } from '@visurel/iconify-angular';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { Ng2SearchPipeModule } from 'ng2-search-filter';
import { MatTooltipModule } from '@angular/material/tooltip';
import { RouterModule } from '@angular/router';
import { AccessRoleService } from '@management/access-role/access-role.service';
import { ToastrService } from '@sdk-ui/ui';
import { UserService } from '@management/user/user.service';

@Component({
  selector: 'kc-user-role-update-form',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    MatDialogModule,
    MatFormFieldModule,
    MatSelectModule,
    MatButtonModule,
    MatIconModule,
    MatProgressSpinnerModule,
    MatTooltipModule,
    IconModule,
    Ng2SearchPipeModule,
    RouterModule
  ],
  templateUrl: './user-role-update-form.component.html',
  styleUrls: ['./user-role-update-form.component.scss']
})
export class UserRoleUpdateFormComponent implements OnInit {
  public readonly user: any = inject(MAT_DIALOG_DATA);
  private accessRoleSvc = inject(AccessRoleService);
  private toastr = inject(ToastrService);
  private userService = inject(UserService);
  private dialogRef = inject<MatDialogRef<UserRoleUpdateFormComponent>>(MatDialogRef);

  icClose = icClose;
  icInfo = icInfo;

  isRolesLoading = true;
  roleList = [];
  searchRoleTerm = '';

  roleIds = [];
  isSubmitting = false;

  ngOnInit(): void {
    let roles = [];
    if (this.user?.roles && this.user?.roles.length) {
      roles = this.user?.roles.map((item: any) => item.id);
      this.roleIds = roles;
    }
    this.getRoles();
  }

  getRoles(): void {
    this.accessRoleSvc.getAccessRoles().subscribe({
      next: data => {
        this.roleList = data;
        this.isRolesLoading = false;
      },
      error: err => {
        this.isRolesLoading = false;
        this.toastr.error(err.message || 'Something Wrong on fetch roles');
      }
    });
  }

  updateRoles(): void {
    this.isSubmitting = true;
    this.userService.mcAssignRoles({ username: this.user.username, roleIds: this.roleIds }).subscribe(
      res => {
        this.toastr.success(res?.message || 'User role updated.');
        this.dialogRef.close(res);
      },
      err => {
        this.toastr.error(err.message || 'Something Wrong on fetch roles');
        this.isSubmitting = false;
      }
    );
  }
}
