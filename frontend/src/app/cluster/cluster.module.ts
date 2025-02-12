import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ClusterRoutingModule } from './cluster-routing.module';
import { ClusterListComponent } from './cluster-list/cluster-list.component';
import { FlexLayoutModule } from '@angular/flex-layout';
import { MatButtonModule } from '@angular/material/button';
import { MatRippleModule } from '@angular/material/core';
import { MatIconModule } from '@angular/material/icon';
import { MatCardModule } from '@angular/material/card';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { CdkStepperModule } from '@angular/cdk/stepper';
import { MatTabsModule } from '@angular/material/tabs';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatSelectModule } from '@angular/material/select';
import { MatInputModule } from '@angular/material/input';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatExpansionModule } from '@angular/material/expansion';
import { MatDialogModule } from '@angular/material/dialog';
import { MatRadioModule } from '@angular/material/radio';
import { ClusterIntroComponent } from './cluster-intro/cluster-intro.component';
import { MatButtonToggleModule } from '@angular/material/button-toggle';
import { Ng2SearchPipeModule } from 'ng2-search-filter';
import { ClusterService } from './cluster.service';
import { SharedModule } from '@shared-ui/shared.module';
import { PageLayoutModule } from '@sdk-ui/ui/page-layout/page-layout.module';
import { IconModule } from '@visurel/iconify-angular';
import { SafeHtmlModule } from '@sdk-ui/pipes/safe-html/safe-html.module';
import { ClusterResolver } from './cluster.resolver';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { CdkHorizontalStepperModule } from '@cdk-ui/horizontal-stepper';
import { CdkClipboardModule } from '@cdk-ui/clipboard';
import { CdkIconModule } from '@cdk-ui/icon';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { OwlDateTimeModule, OwlNativeDateTimeModule } from 'ng-pick-datetime-ex';
import { CdkHintModule } from '@cdk-ui/hint';

const matModules = [
  MatButtonModule,
  MatIconModule,
  MatRippleModule,
  MatCardModule,
  FlexLayoutModule,
  MatRippleModule,
  MatTooltipModule,
  MatProgressSpinnerModule,
  MatTabsModule,
  MatProgressBarModule,
  MatFormFieldModule,
  MatSelectModule,
  MatInputModule,
  MatDialogModule,
  MatRadioModule
];

@NgModule({
  declarations: [
    ClusterListComponent,
    ClusterIntroComponent,
  ],
  imports: [
    CommonModule,
    ClusterRoutingModule,
    ...matModules,
    CdkHorizontalStepperModule,
    CdkStepperModule,
    ReactiveFormsModule,
    MatDatepickerModule,
    OwlDateTimeModule,
    OwlNativeDateTimeModule,
    MatExpansionModule,
    SharedModule,
    FormsModule,
    MatButtonToggleModule,
    Ng2SearchPipeModule,
    PageLayoutModule,
    IconModule,
    SafeHtmlModule,
    MatCheckboxModule,
    CdkClipboardModule,
    MatIconModule,
    CdkIconModule,
    CdkHintModule
  ],
  providers: [ClusterService, ClusterResolver]
})
export class ClusterModule {}
