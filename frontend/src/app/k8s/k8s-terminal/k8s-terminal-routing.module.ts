import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { RoleGuardService } from '@core-ui/guards';
import { K8sTerminalComponent } from './k8s-terminal.component';

const routes: Routes = [
  {
    path: '',
    component: K8sTerminalComponent,
    data: {
      title: 'Terminal',
      permissions: ['*']
    },
    canActivate: [RoleGuardService]
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class K8sTerminalRoutingModule { }
