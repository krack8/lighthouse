import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RoleGuardService } from '@core-ui/guards';
import { CustomResourcesDefinationListComponent } from './custom-resources-defination-list/custom-resources-defination-list.component';
import { CustomResourcesDetailsComponent } from './custom-resources-details/custom-resources-details.component';
import { CustomResourcesListComponent } from './custom-resources-list/custom-resources-list.component';

const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    component: CustomResourcesDefinationListComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Cluster Custom Resources Defination List',
      toolbarTitle: 'Cluster Custom Resources Defination',
      breadcrumb: 'custom-resources-defination',
      permissions: ['VIEW_K8S_CUSTOM_RESOURCE_DEFINATION']
    }
  },
  {
    path: 'custom-resources',
    component: CustomResourcesListComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Cluster Custom Resources List',
      toolbarTitle: 'Cluster Custom Resources',
      breadcrumb: 'custom-resources-defination',
      permissions: ['VIEW_K8S_CUSTOM_RESOURCES']
    }
  },
  {
    path: 'custom-resources/custom-resources-details',
    component: CustomResourcesDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Cluster Custom Resources Details',
      toolbarTitle: 'Cluster Custom Resources Details',
      breadcrumb: 'custom-resources-defination',
      permissions: ['VIEW_K8S_CUSTOM_RESOURCES']
    }
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class K8sClusterCustomResourcesRoutingModule {}
