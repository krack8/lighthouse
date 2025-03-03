import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatDialogModule } from '@angular/material/dialog';
import { MatExpansionModule } from '@angular/material/expansion';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatMenuModule } from '@angular/material/menu';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatSelectModule } from '@angular/material/select';
import { MatSortModule } from '@angular/material/sort';
import { MatTooltipModule } from '@angular/material/tooltip';
import { ExpansionDataViewerTemplateComponent } from '@k8s/common/components/expansion-data-viewer-template/expansion-data-viewer-template.component';
import { MetadataTemplateComponent } from '@k8s/common/components/metadata-template/metadata-template.component';
import { K8sUpdateModule } from '@k8s/k8s-update/k8s-update.module';
import { SharedModule } from '@shared-ui/shared.module';
import { IconModule } from '@visurel/iconify-angular';
import { Ng2SearchPipeModule } from 'ng2-search-filter';
import { K8sNodesRoutingModule } from './k8s-nodes-routing.module';
import { K8sNodesService } from './k8s-nodes.service';
import { NodeDetailsComponent } from './node-details/node-details.component';
import { NodeListComponent } from './node-list/node-list.component';
import { NodeTaintDialogComponent } from './node-list/node-taint-dialog/node-taint-dialog.component';
import { NgApexchartsModule } from 'ng-apexcharts';


@NgModule({
  declarations: [NodeListComponent, NodeDetailsComponent, NodeTaintDialogComponent],
  imports: [
    CommonModule,
    K8sNodesRoutingModule,
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
    ReactiveFormsModule,
    MatInputModule,
    MatProgressBarModule,
    MatCheckboxModule,
    MetadataTemplateComponent,
    ExpansionDataViewerTemplateComponent,
    NgApexchartsModule
  ],
  providers: [K8sNodesService]
})
export class K8sNodesModule {}
