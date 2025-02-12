import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Ng2SearchPipeModule } from 'ng2-search-filter';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatIconModule } from '@angular/material/icon';
import { IconModule } from '@visurel/iconify-angular';
import { MatButtonModule } from '@angular/material/button';
import { SharedModule } from '@shared-ui/shared.module';
import { FormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatSelectModule } from '@angular/material/select';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatSortModule } from '@angular/material/sort';
import { MatMenuModule } from '@angular/material/menu';
import { MatDialogModule } from '@angular/material/dialog';
import { K8sUpdateModule } from '@k8s/k8s-update/k8s-update.module';
import { MatExpansionModule } from '@angular/material/expansion';

import { K8sClusterRoleRoutingModule } from './k8s-cluster-role-routing.module';
import { K8sClusterRoleListComponent } from './k8s-cluster-role-list/k8s-cluster-role-list.component';
import { K8sClusterRoleService } from './k8s-cluster-role.service';
import { K8sClusterRoleDetailsComponent } from './k8s-cluster-role-details/k8s-cluster-role-details.component';
import { MetadataTemplateComponent } from '@k8s/common/components/metadata-template/metadata-template.component';
import { JsonDataViewerTemplateComponent } from '@k8s/common/components/json-data-viewer-template/json-data-viewer-template.component';

@NgModule({
  declarations: [K8sClusterRoleListComponent, K8sClusterRoleDetailsComponent],
  imports: [
    CommonModule,
    K8sClusterRoleRoutingModule,
    FormsModule,
    MatProgressSpinnerModule,
    MatIconModule,
    MatButtonModule,
    MatFormFieldModule,
    MatSelectModule,
    MatTooltipModule,
    MatSortModule,
    MatMenuModule,
    MatDialogModule,
    MatExpansionModule,
    Ng2SearchPipeModule,
    SharedModule,
    IconModule,
    K8sUpdateModule,
    MatIconModule,
    MetadataTemplateComponent,
    JsonDataViewerTemplateComponent
  ],
  providers: [K8sClusterRoleService]
})
export class K8sClusterRoleModule {}
