import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { RoleGuardService } from '@core-ui/guards';
import { K8sClusterRoleDetailsComponent } from './k8s-cluster-role-details/k8s-cluster-role-details.component';
import { K8sClusterRoleListComponent } from './k8s-cluster-role-list/k8s-cluster-role-list.component';

const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    component: K8sClusterRoleListComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Cluster Role List',
      toolbarTitle: 'Cluster Role',
      breadcrumb: 'cluster-role',
      permissions: ['VIEW_K8S_CLUSTER_ROLE']
    }
  },
  {
    path: 'cluster-role-details',
    component: K8sClusterRoleDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Cluster Role | Details',
      toolbarTitle: 'Cluster Role Details',
      breadcrumb: 'cluster-role',
      permissions: ['VIEW_K8S_CLUSTER_ROLE']
    }
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class K8sClusterRoleRoutingModule {}
