import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { K8sStorageClassRoutingModule } from './k8s-storage-class-routing.module';
import { K8sStorageClassListComponent } from './k8s-storage-class-list/k8s-storage-class-list.component';
import { K8sStorageClassDetailsComponent } from './k8s-storage-class-details/k8s-storage-class-details.component';

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
import { K8sStorageClassService } from './k8s-storage-class.service';
import { MetadataTemplateComponent } from '@k8s/common/components/metadata-template/metadata-template.component';

@NgModule({
  declarations: [K8sStorageClassListComponent, K8sStorageClassDetailsComponent],
  imports: [
    CommonModule,
    K8sStorageClassRoutingModule,
    CommonModule,
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
  providers: [K8sStorageClassService]
})
export class K8sStorageClassModule {}
