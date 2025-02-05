import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { Error500Component } from './error-500.component';
import { KcRoutes } from '@sdk-ui/interfaces/kc-route.interface';

const routes: KcRoutes = [
  {
    path: '',
    component: Error500Component,
    data: {
      containerEnabled: true,
      toolbarShadowEnabled: true
    }
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class Error500RoutingModule {}
