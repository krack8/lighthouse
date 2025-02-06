import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { K8sRoutingModule } from './k8s-routing.module';
import { K8sComponent } from './k8s.component';
import { SharedModule } from '@shared-ui/shared.module';
import { IconModule } from '@visurel/iconify-angular';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { PageLayoutModule } from '@sdk-ui/ui/page-layout/page-layout.module';
import { K8sService } from './k8s.service';
import { ClusterService } from '../cluster/cluster.service';
import { K8sResolver } from './k8s.resolver';
import { K8sNamespacesService } from './k8s-namespaces/k8s-namespaces.service';
import { K8sPersistentVolumeService } from './k8s-persistent-volume/k8s-persistent-volume.service';
import { K8sClusterRoleService } from './k8s-cluster-role/k8s-cluster-role.service';
import { K8sClusterRoleBindingService } from './k8s-cluster-role-binding/k8s-cluster-role-binding.service';
import { K8sStorageClassService } from './k8s-storage-class/k8s-storage-class.service';
import { K8sClusterCustomResourcesService } from './k8s-cluster-custom-resources/k8s-cluster-custom-resources.service';
import { K8sUpdateModule } from './k8s-update/k8s-update.module';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatMenuModule } from '@angular/material/menu';
import { FormsModule } from '@angular/forms';
import { Ng2SearchPipeModule } from 'ng2-search-filter';
@NgModule({
  declarations: [K8sComponent],
  imports: [
    CommonModule,
    K8sRoutingModule,
    SharedModule,
    PageLayoutModule,
    IconModule,
    MatButtonModule,
    MatIconModule,
    MatProgressSpinnerModule,
    K8sUpdateModule,
    MatTooltipModule,
    MatMenuModule,
    FormsModule,
    Ng2SearchPipeModule
  ],
  providers: [
    K8sService,
    K8sResolver,
    ClusterService,
    K8sNamespacesService,
    K8sPersistentVolumeService,
    K8sClusterRoleService,
    K8sClusterRoleBindingService,
    K8sStorageClassService,
    K8sClusterCustomResourcesService
  ]
})
export class K8sModule {}
