import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { K8sUpdateComponent } from './k8s-update.component';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatTabsModule } from '@angular/material/tabs';
import { AceEditorModule } from '@klovercloud/ace-editor/ace-editor.module';
import { MatDialogModule } from '@angular/material/dialog';
import { K8sNamespacesService } from '@k8s/k8s-namespaces/k8s-namespaces.service';
import { K8sPersistentVolumeService } from '@k8s/k8s-persistent-volume/k8s-persistent-volume.service';
import { K8sClusterRoleBindingService } from '@k8s/k8s-cluster-role-binding/k8s-cluster-role-binding.service';
import { K8sClusterRoleService } from '@k8s/k8s-cluster-role/k8s-cluster-role.service';
import { K8sStorageClassService } from '@k8s/k8s-storage-class/k8s-storage-class.service';
import { K8sClusterCustomResourcesService } from '@k8s/k8s-cluster-custom-resources/k8s-cluster-custom-resources.service';
import { K8sNodesService } from '@k8s/k8s-nodes/k8s-nodes.service';
import { MatProgressBarModule } from '@angular/material/progress-bar';

@NgModule({
  declarations: [K8sUpdateComponent],
  imports: [
    CommonModule,
    MatCardModule,
    MatTabsModule,
    MatButtonModule,
    MatCardModule,
    AceEditorModule,
    MatDialogModule,
    MatProgressBarModule
  ],
  exports: [K8sUpdateComponent],
  providers: [
    K8sNamespacesService,
    K8sPersistentVolumeService,
    K8sClusterRoleBindingService,
    K8sClusterRoleService,
    K8sStorageClassService,
    K8sClusterCustomResourcesService,
    K8sNodesService
  ]
})
export class K8sUpdateModule {}
