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

import { K8sPersistentVolumeRoutingModule } from './k8s-persistent-volume-routing.module';
import { PvListComponent } from './pv-list/pv-list.component';
import { K8sPersistentVolumeService } from './k8s-persistent-volume.service';
import { PvDetailsComponent } from './pv-details/pv-details.component';
import { MetadataTemplateComponent } from '@k8s/common/components/metadata-template/metadata-template.component';

@NgModule({
  declarations: [PvListComponent, PvDetailsComponent],
  imports: [
    CommonModule,
    K8sPersistentVolumeRoutingModule,
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
  providers: [K8sPersistentVolumeService]
})
export class K8sPersistentVolumeModule {}
