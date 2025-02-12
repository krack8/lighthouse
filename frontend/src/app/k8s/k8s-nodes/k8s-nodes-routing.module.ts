import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { RoleGuardService } from '@core-ui/guards';
import { NodeDetailsComponent } from './node-details/node-details.component';
import { NodeListComponent } from './node-list/node-list.component';

const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    component: NodeListComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Node List',
      toolbarTitle: 'Nodes',
      breadcrumb: 'Nodes',
      permissions: ['VIEW_K8S_NODES']
    }
  },
  {
    path: 'node-details',
    component: NodeDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Nodes | Details',
      toolbarTitle: 'Node Details',
      breadcrumb: 'Nodes',
      permissions: ['VIEW_K8S_NODES']
    }
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class K8sNodesRoutingModule {}
