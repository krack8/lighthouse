import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { UserListComponent } from './user-list/user-list.component';
import { UserDetailsComponent } from './user-details/user-details.component';
import { RoleGuardService } from '@core-ui/guards';

const routes: Routes = [
  {
    path: '',
    component: UserListComponent,
    data: {
      title: 'Management | User',
      permissions: ['VIEW_USER']
    },
    canActivate: [RoleGuardService]
  },
  {
    path: ':id',
    component: UserDetailsComponent,
    data: {
      title: 'Management | User | Details',
      permissions: ['MANAGE_USER']
    },
    canActivate: [RoleGuardService]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class UserRoutingModule {}
