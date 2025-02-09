import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { UserProfileDetailsComponent } from './user-profile-details/user-profile-details.component';
import { RoleGuardService } from '@core-ui/guards';

const routes: Routes = [
  {
    path: '',
    component: UserProfileDetailsComponent,
    data: {
      title: 'My Profile',
      toolbarTitle: 'My Profile',
      permissions: ['*']
    },
    canActivate: [RoleGuardService]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class UserProfileRoutingModule {}
