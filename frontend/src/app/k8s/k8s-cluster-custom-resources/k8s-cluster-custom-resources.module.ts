import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatDialogModule } from '@angular/material/dialog';
import { MatExpansionModule } from '@angular/material/expansion';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatMenuModule } from '@angular/material/menu';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatSelectModule } from '@angular/material/select';
import { MatSortModule } from '@angular/material/sort';
import { MatTooltipModule } from '@angular/material/tooltip';
import { K8sUpdateModule } from '@k8s/k8s-update/k8s-update.module';
import { SharedModule } from '@shared-ui/shared.module';
import { IconModule } from '@visurel/iconify-angular';
import { Ng2SearchPipeModule } from 'ng2-search-filter';

import { MetadataTemplateComponent } from '@k8s/common/components/metadata-template/metadata-template.component';
import { CustomResourcesDefinationDetailsComponent } from './custom-resources-defination-details/custom-resources-defination-details.component';
import { CustomResourcesDefinationListComponent } from './custom-resources-defination-list/custom-resources-defination-list.component';
import { CustomResourcesDetailsComponent } from './custom-resources-details/custom-resources-details.component';
import { CustomResourcesListComponent } from './custom-resources-list/custom-resources-list.component';
import { K8sClusterCustomResourcesRoutingModule } from './k8s-cluster-custom-resources-routing.module';
import { K8sClusterCustomResourcesService } from './k8s-cluster-custom-resources.service';

@NgModule({
  declarations: [
    CustomResourcesListComponent,
    CustomResourcesDefinationListComponent,
    CustomResourcesDetailsComponent,
    CustomResourcesDefinationDetailsComponent
  ],
  imports: [
    CommonModule,
    K8sClusterCustomResourcesRoutingModule,
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
    MetadataTemplateComponent
  ],
  providers: [K8sClusterCustomResourcesService]
})
export class K8sClusterCustomResourcesModule {}
