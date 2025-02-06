import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { RoleGuardService } from '@core-ui/guards';
import { Error403Component } from './error-403.component';

const routes: Routes = [
  {
    path: '',
    component: Error403Component,
    data: {
      title: 'Forbidden',
      containerEnabled: true,
      toolbarShadowEnabled: true,
      permissions: ['*']
    },
    canActivate: [RoleGuardService]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class Error403RoutingModule {}
