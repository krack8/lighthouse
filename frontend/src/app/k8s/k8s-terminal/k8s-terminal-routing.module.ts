import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { RoleGuardService } from '@core-ui/guards';
import { PodTerminalComponent } from './pod-terminal/pod-terminal.component';

const routes: Routes = [
  {
    path: '',
    component: PodTerminalComponent,
    data: {
      title: 'Terminal',
      permissions: ['*']
    },
    canActivate: [RoleGuardService]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class K8sTerminalRoutingModule {}
