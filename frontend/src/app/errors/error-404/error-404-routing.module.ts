import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { Error404Component } from './error-404.component';
import { KcRoutes } from '@sdk-ui/interfaces/kc-route.interface';
import { RoleGuardService } from '@core-ui/guards';

const routes: KcRoutes = [
  {
    path: '',
    pathMatch: 'full',
    component: Error404Component,
    data: {
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
export class Error404RoutingModule {}
