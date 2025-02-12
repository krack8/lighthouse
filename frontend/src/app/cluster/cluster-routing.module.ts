import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ClusterListComponent } from './cluster-list/cluster-list.component';
import { RoleGuardService } from '@core-ui/guards';
import { ClusterResolver } from './cluster.resolver';

const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    component: ClusterListComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster',
      toolbarTitle: 'Cluster List',
      permissions: ['*']
    }
  },
  {
    path: 'create',
    loadComponent: () => import('./cluster-form/cluster-form.component').then(m => m.ClusterFormComponent),
    data: {
      title: 'Create Agent Cluster',
      toolbarTitle: 'Create Agent Cluster',
      permissions: ['CREATE_CLUSTER']
    }
  },
  {
    path: ':clusterId/init',
    loadComponent: () => import('./cluster-init/cluster-init.component').then(m => m.ClusterInitComponent),
    resolve: { clusterDetails: ClusterResolver },
    data: {
      title: 'Cluster Initialization',
      toolbarTitle: 'Cluster Initialization',
      permissions: ['*']
    }
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ClusterRoutingModule {}
