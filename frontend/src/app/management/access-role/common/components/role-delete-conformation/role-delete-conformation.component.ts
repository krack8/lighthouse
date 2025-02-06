import { Component, Inject, OnInit } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { ToastrService } from '@sdk-ui/ui';
import { AccessRoleService } from '@management/access-role/access-role.service';

@Component({
  selector: 'kc-role-delete-conformation',
  templateUrl: './role-delete-conformation.component.html',
  styleUrls: ['./role-delete-conformation.component.scss']
})
export class RoleDeleteConformationComponent implements OnInit {
  isDeleting!: boolean;
  isLoadingData: boolean = true;
  userList: any[] = [];
  fetchUserError!: string;

  constructor(
    @Inject(MAT_DIALOG_DATA) public role: any,
    private dialogRef: MatDialogRef<RoleDeleteConformationComponent>,
    private accessRoleService: AccessRoleService,
    private toastr: ToastrService
  ) {}

  ngOnInit(): void {
    this.getUsersFromRole();
  }

  getUsersFromRole(): void {
    this.accessRoleService.getUsersFromRole(this.role?.id).subscribe({
      next: res => {
        this.isLoadingData = false;
        this.userList = res?.data;
      },
      error: err => {
        this.toastr.error(err?.error?.message);
        this.isLoadingData = false;
        this.fetchUserError = err?.error?.message;
      }
    });
  }

  deleteRole(): void {
    this.isDeleting = true;
    this.accessRoleService.deleteAccessRole(this.role.id).subscribe({
      next: res => {
        this.toastr.success(res?.message);
        this.dialogRef.close(true);
        this.isDeleting = false;
      },
      error: err => {
        this.toastr.error(err?.error?.message);
        this.isDeleting = false;
      }
    });
  }
}
