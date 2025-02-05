import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { RoleGuardService } from '@core-ui/guards';
import { K8sClusterRoleBindingDetailsComponent } from './k8s-cluster-role-binding-details/k8s-cluster-role-binding-details.component';
import { K8sClusterRoleBindingListComponent } from './k8s-cluster-role-binding-list/k8s-cluster-role-binding-list.component';

const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    component: K8sClusterRoleBindingListComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Cluster Role Binding List',
      toolbarTitle: 'Cluster Role Binding',
      breadcrumb: 'cluster-role-binding',
      permissions: ['VIEW_K8S_CLUSTER_ROLE_BINDING']
    }
  },
  {
    path: 'cluster-role-binding-details',
    component: K8sClusterRoleBindingDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Cluster Role Binding | Details',
      toolbarTitle: 'Cluster Role Binding Details',
      breadcrumb: 'cluster-role-binding',
      permissions: ['VIEW_K8S_CLUSTER_ROLE_BINDING']
    }
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class K8sClusterRoleBindingRoutingModule {}
