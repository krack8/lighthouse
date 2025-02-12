import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { RoleGuardService } from '@core-ui/guards';
import { K8sStorageClassListComponent } from './k8s-storage-class-list/k8s-storage-class-list.component';
import { K8sStorageClassDetailsComponent } from './k8s-storage-class-details/k8s-storage-class-details.component';

const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    component: K8sStorageClassListComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Storage Class List',
      toolbarTitle: 'Storage Class',
      breadcrumb: 'storage-class',
      permissions: ['VIEW_K8S_STORAGE_CLASS']
    }
  },
  {
    path: 'storage-class-details',
    component: K8sStorageClassDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Storage Class | Details',
      toolbarTitle: 'Storage Class Details',
      breadcrumb: 'storage-class',
      permissions: ['VIEW_K8S_STORAGE_CLASS']
    }
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class K8sStorageClassRoutingModule {}
