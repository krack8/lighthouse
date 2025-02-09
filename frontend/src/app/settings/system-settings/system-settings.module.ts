import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { SystemSettingsRoutingModule } from './system-settings-routing.module';

import { SystemSettingsComponent } from './containers/system-settings/system-settings.component';
import { ThemeInfoComponent } from './common/components/theme-info/theme-info.component';
import { FlexLayoutModule } from '@angular/flex-layout';
import { MatInputModule } from '@angular/material/input';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatDialogModule } from '@angular/material/dialog';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatSelectModule } from '@angular/material/select';
import { MatRadioModule } from '@angular/material/radio';
import { IconModule } from '@visurel/iconify-angular';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { SharedModule } from '@shared-ui/shared.module';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { PageLayoutModule } from '@sdk-ui/ui/page-layout/page-layout.module';
import { ContainerModule } from '@sdk-ui/directives/container/container.module';
import { SystemSettingsService } from './system-settings.service';
import { CdkIconModule } from '@cdk-ui/icon';

@NgModule({
  declarations: [SystemSettingsComponent, ThemeInfoComponent],
  imports: [
    CommonModule,
    SystemSettingsRoutingModule,
    FormsModule,
    ReactiveFormsModule,
    SharedModule,
    PageLayoutModule,
    IconModule,
    ContainerModule,
    FlexLayoutModule,
    MatInputModule,
    MatCheckboxModule,
    MatSelectModule,
    MatIconModule,
    MatButtonModule,
    MatTooltipModule,
    MatDialogModule,
    MatProgressSpinnerModule,
    MatRadioModule,
    MatProgressBarModule,
    CdkIconModule
  ],
  providers: [SystemSettingsService]
})
export class SystemSettingsModule {}
