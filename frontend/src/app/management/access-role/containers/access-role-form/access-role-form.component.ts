import { ChangeDetectorRef, Component, OnInit, ViewChild } from '@angular/core';
import { UntypedFormArray, UntypedFormBuilder, UntypedFormControl, UntypedFormGroup, Validators } from '@angular/forms';
import { MatCheckboxChange } from '@angular/material/checkbox';
import { MatStepper } from '@angular/material/stepper';
import { ActivatedRoute, Router } from '@angular/router';
import { ToastrService } from '@sdk-ui/ui';
import { map } from 'rxjs/operators';
import { fadeInRight400ms } from '@sdk-ui/animations/fade-in-right.animation';
import { fadeInUp400ms } from '@sdk-ui/animations/fade-in-up.animation';
import { IFormCategoryPermission, IFormPermission, IPermissionListObject } from '../../access-role-interface';
import { AccessRoleService } from '../../access-role.service';
import icSearch from '@iconify/icons-ic/search';
import { ToolbarService } from '@sdk-ui/services/toolbar.service';

@Component({
  selector: 'kc-access-role-form',
  templateUrl: './access-role-form.component.html',
  styleUrls: ['./access-role-form.component.scss'],
  animations: [fadeInUp400ms, fadeInRight400ms],
})
export class AccessRoleFormComponent implements OnInit {
  @ViewChild('stepper') private stepper!: MatStepper;
  // @ViewChildren("permissionCheckbox", { read: MatCheckbox }) permissionCheckboxList!: QueryList<MatCheckbox>;
  icSearch = icSearch;

  permissions: IFormCategoryPermission[] = [];

  roleForm: UntypedFormGroup = this.fb.group({
    roleName: ['', Validators.required],
    description: [''],
    listOfPermissionsID: this.fb.array([], Validators.required),
  });
  searchAccess!: string;
  submitting: boolean = false;

  roleId!: string;
  permissionLoading!: boolean;
  roleLoading!: boolean;

  constructor(
    private fb: UntypedFormBuilder,
    private accessRoleSvc: AccessRoleService,
    private route: ActivatedRoute,
    private router: Router,
    private toastr: ToastrService,
    private toolbarService: ToolbarService,
    private cd: ChangeDetectorRef
  ) {}

  ngOnInit() {
    this.toolbarService.changeData({ title: 'Roles' });
    this.roleId = this.route.snapshot.params['id'];
    console.log('this.roleId', this.roleId);
    this.getPermissions();
  }

  get permissionIdsArray(): UntypedFormArray {
    return this.roleForm.get('listOfPermissionsID') as UntypedFormArray;
  }

  getCategoryGroup(categoryName: string): IFormCategoryPermission {
    return this.permissions.find((item) => item.name === categoryName) as IFormCategoryPermission;
  }

  isCategoryContainerPermission(category: string): boolean {
    return this.getCategoryGroup(category).permissions.some((item: any) => item.checked);
  }

  // Form Inputs
  setAll(e: MatCheckboxChange): void {
    const _categoryName: string = e.source.name || '';
    const _completed = e.checked;
    const _categoryPermission = this.getCategoryGroup(_categoryName);
    _categoryPermission.completed = _completed;
    if (_categoryPermission == null) {
      return;
    }
    const permissionIds = this.permissionIdsArray.value;
    _categoryPermission.permissions.forEach((t) => {
      if (_completed) {
        if (!permissionIds.includes(t.id)) {
          this.permissionIdsArray.push(new UntypedFormControl(t.id));
        }
      } else {
        if (t.checked) {
          this.permissionIdsArray.removeAt(this.permissionIdsArray.value.indexOf(t.id));
        }
      }
      return (t.checked = _completed);
    });
    this.cd.detectChanges();
  }

  // Group checkbox trigger
  someChecked(catName: string): boolean {
    const _categoryPermission: IFormCategoryPermission = this.getCategoryGroup(catName);
    const checkedItems = _categoryPermission.permissions.filter((t) => t.checked);
    return checkedItems.length > 0 && !_categoryPermission.completed;
  }

  // Single Permission checkbox trigger
  permissionCheckEvent(e: MatCheckboxChange, catName: string) {
    const _categoryPermission: IFormCategoryPermission = this.getCategoryGroup(catName);
    // Update all checked
    if (e.source.checked) {
      console.log(e);
      this.permissionIdsArray.push(new UntypedFormControl(e.source.value));

      let _catViewName = 'VIEW_' + catName;
      if (catName === 'VPC') {
        _catViewName = 'VIEW_NAMESPACE';
      }
      const __name = e.source.name;
      if (__name !== _catViewName) {
        const __permissions = _categoryPermission.permissions;
        const __permissionNames = __permissions.map((item) => item.name);
        // ------------- handling Non VIEW_CATEGORY ITEM
        // check VIEW_CATEGORY item
        const viewItemIndex = __permissionNames.findIndex((name) => name === _catViewName);
        if (viewItemIndex !== -1 && !_categoryPermission.permissions[viewItemIndex].checked) {
          this.permissionIdsArray.push(new UntypedFormControl(_categoryPermission.permissions[viewItemIndex].id));
          _categoryPermission.permissions[viewItemIndex].checked = true;
        }
        // check VIEW_FEATURE Item
        if (!__name.includes('VIEW_')) {
          const featureViewName = 'VIEW_' + __name.split('_').slice(1).join('_');
          const featureViewItemIndex = __permissionNames.findIndex((name) => name === featureViewName);
          if (featureViewItemIndex !== -1 && !_categoryPermission.permissions[featureViewItemIndex].checked) {
            this.permissionIdsArray.push(
              new UntypedFormControl(_categoryPermission.permissions[featureViewItemIndex].id)
            );
            _categoryPermission.permissions[featureViewItemIndex].checked = true;
          }
        }
      }
    }
    else {
      const permissionIds = this.permissionIdsArray.value;
      permissionIds.forEach((id: string, i: number) => {
        if (id == e.source.value) {
          this.permissionIdsArray.removeAt(i);
          return;
        }
      });
    }
    _categoryPermission.completed =
      _categoryPermission.permissions !== null && _categoryPermission.permissions.every((t) => t.checked);
  }

  private routeToRoleList(): void {
    this.router.navigate(['/manage/roles']);
  }

  stopPropagation(event: any): void {
    event.stopPropagation();
  }

  onSubmit() {
    this.submitting = true;
    const formData = this.roleForm.value;
    const payload = {
      name: formData.roleName,
      description: formData.description,
      permissions: this.permissionIdsArray.value,
    }
    if (!payload.permissions.length) {
      this.toastr.warn('Min 1 access permission required', 'Permission Required');
      this.submitting = false;
      this.stepper.selectedIndex = 0;
      return;
    }
    if (this.roleId) {
      this.accessRoleSvc.updateAccessRole(this.roleId, payload).subscribe(
        (_) => {
          this.toastr.success(`${payload.name} role updated`);
          this.routeToRoleList();
        },
        (err) => {
          this.submitting = false;
          this.toastr.error(err.message || 'Role update Failed');
        }
      );
      return;
    }
    this.accessRoleSvc.createAccessRole(payload).subscribe(
      (res) => {
        this.toastr.success(`${payload.name} role created`);
        this.routeToRoleList();
      },
      (err) => {
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
        map((data) => {
          console.log('data', data);
          const _permission: IFormCategoryPermission[] = [];
          const resPermissionData: IPermissionListObject = data;
          Object.entries(resPermissionData).filter((cat) => {
            if (cat[1].length) {
              const _permissions: IFormPermission[] = cat[1].map((item) => {
                return { ...item, checked: false };
              });
              const _categoryPermission: IFormCategoryPermission = {
                name: cat[0],
                label: cat[0]?.replace(new RegExp('_', 'g'), ' '),
                completed: false,
                permissions: _permissions,
              };
              _permission.push(_categoryPermission);
            }
          });
          return _permission;
        })
      )
      .subscribe({
        next: (_permissionList) => {
          if (this.roleId) {
            this.getRole();
          }
          this.permissions = _permissionList;
          this.permissionLoading = false;
        },
        error: (err) => {
          this.toastr.error(err.message || 'Get Permissions Failed');
          this.permissionLoading = false;
        },
      });
  }

  getRole(): void {
    this.roleLoading = true;
    this.accessRoleSvc.getAccessRole(this.roleId).subscribe({
      next: (res) => {
        let _data = res;
        let listOfPermission: { [key: string]: any[] } = {};

        const permissionArray = _data.permissions;
        // all permissions are categorized by their category
        let categorizedPermissions: { [key: string]: any[] } = {};
        permissionArray.forEach((permission) => {
          if (!categorizedPermissions[permission.category]) {
            categorizedPermissions[permission.category] = [];
          }
          categorizedPermissions[permission.category].push(permission);
        });

        let CurrentListOfPermission: { [key: string]: any[] } = categorizedPermissions;
        for (let [_category, _permissions] of Object.entries(CurrentListOfPermission)) {
          if (_permissions.length) {
            listOfPermission[_category] = _permissions;
            // Prepare form payload
            const permissionsLen = _permissions.length;
            const _categoryPermission = this.getCategoryGroup(_category);
            if (_categoryPermission?.permissions?.length === permissionsLen) {
              _categoryPermission.completed = true;
            }
            _permissions.forEach((item: any) => {
              this.permissionIdsArray.push(new UntypedFormControl(item.id));
              const perm = _categoryPermission?.permissions?.find((_item) => _item.id === item.id);
              if (perm) {
                perm.checked = true;
              }
            });
          }
        }
        const formPayload = {
          roleName: _data?.name,
          description: _data.description,
        };
        this.roleForm.patchValue(formPayload);
        this.roleLoading = false;
      },
      error: (err) => {
        this.roleLoading = false;
        this.toastr.error(err.message || 'Role Fetch Failed');
        this.routeToRoleList();
      },
    });
  }

  checkDependency(catName: string, permissionName: string, disableCategory?: string): boolean {
    const categoryPermission: IFormCategoryPermission = this.getCategoryGroup(catName);
    if (categoryPermission) {
      const foundPermission = categoryPermission.permissions.find((permission) => permission.name === permissionName);

      if (foundPermission) {
        if (foundPermission.checked === true) {
          return true;
        } else if (foundPermission.checked === false) {
          return false;
        }
      } else {
        this.toastr.error(`Permission ${permissionName} not found in ${catName}.`, 'Dependency Permission Not Found');
        return false;
      }
    } else {
      this.toastr.error(`Category ${catName} not found.`, 'Dependency Category Not Found');
      return false;
    }
  }

  unsetAll(category: string): void {
    const _categoryName: string = category || '';
    const _completed = false;
    const _categoryPermission = this.getCategoryGroup(_categoryName);
    _categoryPermission.completed = _completed;
    if (_categoryPermission == null) {
      return;
    }
    const permissionIds = this.permissionIdsArray.value;
    _categoryPermission.permissions.forEach((t) => {
      if (_completed) {
        if (!permissionIds.includes(t.id)) {
          this.permissionIdsArray.push(new UntypedFormControl(t.id));
        }
      } else {
        if (t.checked) {
          this.permissionIdsArray.removeAt(this.permissionIdsArray.value.indexOf(t.id));
        }
      }
      return (t.checked = _completed);
    });
    this.cd.detectChanges();
  }
}
