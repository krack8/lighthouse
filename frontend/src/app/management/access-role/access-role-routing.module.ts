import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AccessRoleFormComponent } from './containers/access-role-form/access-role-form.component';
import { AccessRoleListComponent } from './containers/access-role-list/access-role-list.component';
import { AuthGuard } from '@core-ui/guards';

const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    component: AccessRoleListComponent,
    canActivate: [AuthGuard],
    data: {
      title: 'Roles',
      toolbarTitle: 'Roles',
      permissions: ['VIEW_ROLE']
    }
  },
  {
    path: 'create',
    component: AccessRoleFormComponent,
    canActivate: [AuthGuard],
    data: {
      title: 'Create Role',
      toolbarTitle: 'Roles',
      permissions: ['MANAGE_ROLE']
    }
  },
  {
    path: ':id/edit',
    component: AccessRoleFormComponent,
    canActivate: [AuthGuard],
    data: {
      title: 'Update Role',
      toolbarTitle: 'Roles',
      permissions: ['MANAGE_ROLE']
    }
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class AccessRoleRoutingModule {}
