import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { ToastrService } from '@sdk-ui/ui';
import { AccessRoleService } from '../../access-role.service';

import icInfo from '@iconify/icons-ic/info';
import icAdd from '@iconify/icons-ic/add-circle-outline';
import icSearch from '@iconify/icons-ic/search';
import icCategory from '@iconify/icons-ic/twotone-category';
import { ToolbarService } from '@sdk-ui/services/toolbar.service';
import { RoleDeleteConformationComponent } from '@management/access-role/common/components/role-delete-conformation/role-delete-conformation.component';
import { trackById } from '@core-ui/utils';
import { IRole } from '../../access-role-interface';

@Component({
  selector: 'kc-access-role-list',
  templateUrl: './access-role-list.component.html',
  styleUrls: ['./access-role-list.component.scss']
})
export class AccessRoleListComponent implements OnInit {
  icAdd = icAdd;
  icInfo = icInfo;
  icSearch = icSearch;
  icCategory = icCategory;
  isLoading!: boolean;
  roleList: IRole[] = [];
  searchAccess: string = '';

  trackById = trackById;

  constructor(
    private accessRoleSvc: AccessRoleService,
    private toastr: ToastrService,
    private dialog: MatDialog,
    private toolbarService: ToolbarService
  ) {}

  ngOnInit(): void {
    this.toolbarService.changeData({ title: 'Roles' });
    this.getRoles();
  }

  getRoles(): void {
    this.isLoading = true;
    this.accessRoleSvc.getAccessRoles().subscribe({
      next: _roleList => {
        this.isLoading = false;
        this.roleList = _roleList;
      },
      error: err => {
        this.isLoading = false;
        this.toastr.error(err.message || 'Something wrong on role list');
      }
    });
  }

  deleteRole(role: any): void {
    const dialog = this.dialog.open(RoleDeleteConformationComponent, {
      data: role,
      maxWidth: '450px',
      width: '450px',
      minHeight: '290px'
    });
    dialog.afterClosed().subscribe(res => {
      if (res === true) {
        this.getRoles();
      }
    });
  }
}
