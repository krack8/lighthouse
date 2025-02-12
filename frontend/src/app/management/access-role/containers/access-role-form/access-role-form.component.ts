import { Component, OnInit, ViewChild } from '@angular/core';
import { UntypedFormArray, UntypedFormBuilder, UntypedFormGroup, Validators } from '@angular/forms';
import { MatStepper } from '@angular/material/stepper';
import { ActivatedRoute, Router } from '@angular/router';
import { ToastrService } from '@sdk-ui/ui';
import { map } from 'rxjs/operators';
import { fadeInRight400ms } from '@sdk-ui/animations/fade-in-right.animation';
import { fadeInUp400ms } from '@sdk-ui/animations/fade-in-up.animation';
import { IFormPermission, IPermission } from '../../access-role-interface';
import { AccessRoleService } from '../../access-role.service';
import icSearch from '@iconify/icons-ic/search';
import { ToolbarService } from '@sdk-ui/services/toolbar.service';
import { StrReplacePipe } from '@shared-ui/pipes';

@Component({
  selector: 'kc-access-role-form',
  templateUrl: './access-role-form.component.html',
  styleUrls: ['./access-role-form.component.scss'],
  animations: [fadeInUp400ms, fadeInRight400ms],
  providers: [StrReplacePipe]
})
export class AccessRoleFormComponent implements OnInit {
  @ViewChild('stepper') private stepper!: MatStepper;
  icSearch = icSearch;

  roleForm: UntypedFormGroup = this.fb.group({
    name: ['', Validators.required],
    description: ['']
  });

  allComplete = false;
  submitting: boolean = false;
  roleId!: string;
  roleLoading!: boolean;

  searchAccess = '';
  permissionLoading!: boolean;
  permissions: IFormPermission[] = [];

  constructor(
    private fb: UntypedFormBuilder,
    private accessRoleSvc: AccessRoleService,
    private route: ActivatedRoute,
    private router: Router,
    private toastr: ToastrService,
    private toolbarService: ToolbarService,
    private strReplacePipe: StrReplacePipe
  ) {}

  ngOnInit() {
    this.toolbarService.changeData({ title: 'Roles' });
    this.roleId = this.route.snapshot.params['id'];
    this.getPermissions();
  }

  next(): void {
    if (this.permissions.some(item => item.checked)) {
      this.stepper.selectedIndex = 1;
      return;
    }
    this.toastr.warn('Min 1 access permission required', 'Permission Required');
  }

  get permissionIdsArray(): UntypedFormArray {
    return this.roleForm.get('permissions') as UntypedFormArray;
  }

  // Form Inputs
  setAll(checked: boolean): void {
    this.allComplete = checked;
    this.permissions.forEach(t => (t.checked = checked));
  }

  someComplete() {
    return this.permissions.filter(t => t.checked).length > 0 && !this.allComplete;
  }

  updateAllComplete() {
    this.allComplete = this.permissions.every(t => t.checked);
  }

  private routeToRoleList(): void {
    this.router.navigate(['/manage/roles']);
  }

  onSubmit() {
    this.submitting = true;
    const payload = this.roleForm.value;

    const _permissionIds = [];
    this.permissions.forEach(item => {
      if (item.checked) _permissionIds.push(item.id);
    });
    if (!_permissionIds.length) {
      this.toastr.warn('Min 1 access permission required', 'Permission Required');
      this.submitting = false;
      this.stepper.selectedIndex = 0;
      return;
    }
    payload['permissions'] = _permissionIds;

    if (this.roleId) {
      this.accessRoleSvc.updateAccessRole(this.roleId, payload).subscribe(
        _ => {
          this.toastr.success(`${payload.name} role updated`);
          this.routeToRoleList();
        },
        err => {
          this.submitting = false;
          this.toastr.error(err.message || 'Role update Failed');
        }
      );
      return;
    }
    this.accessRoleSvc.createAccessRole(payload).subscribe(
      res => {
        this.toastr.success(`${payload.name} role created`);
        this.routeToRoleList();
      },
      err => {
        this.submitting = false;
        this.toastr.error(err.message || 'Create Role Failed');
      }
    );
  }

  // API
  getPermissions() {
    this.permissionLoading = true;
    this.accessRoleSvc
      .getAccessPermissions()
      .pipe(
        map((permissionList: IPermission[]) => {
          return permissionList.map(permission => ({
            ...permission,
            label: this.strReplacePipe.transform(permission.name, '_'),
            checked: false
          }));
        })
      )
      .subscribe({
        next: _permissionList => {
          this.permissions = _permissionList;
          this.permissionLoading = false;

          if (this.roleId) this.getRole();
        },
        error: err => {
          this.toastr.error(err.message || 'Get Permissions Failed');
          this.permissionLoading = false;
        }
      });
  }

  getRole(): void {
    this.roleLoading = true;
    this.accessRoleSvc.getAccessRole(this.roleId).subscribe({
      next: _role => {
        if (_role.permissions?.length > 0) {
          const currentPermissionMap = new Map<string, Boolean>();
          _role.permissions.forEach(_item => currentPermissionMap.set(_item.id, true));

          this.allComplete = currentPermissionMap?.size === this.permissions.length;

          // Override existing array
          for (let i = 0; i < this.permissions.length; i++) {
            if (currentPermissionMap.get(this.permissions[i].id)) this.permissions[i].checked = true;
          }
        }

        const formPayload = {
          name: _role.name,
          description: _role.description
        };
        this.roleForm.patchValue(formPayload);
        this.roleLoading = false;
      },
      error: err => {
        this.roleLoading = false;
        this.toastr.error(err.message || 'Role Fetch Failed');
        this.routeToRoleList();
      }
    });
  }
}
