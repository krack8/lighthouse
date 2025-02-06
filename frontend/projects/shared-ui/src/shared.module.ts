import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatProgressBarModule } from '@angular/material/progress-bar';

import { IconModule } from '@visurel/iconify-angular';

import { ConfirmDialogComponent, SecureDeleteDialogComponent, NothingFoundComponent, ConfirmDialogStaticComponent } from './ui';

import {
  FormatCpuPipe,
  FormatMemoryPipe,
  FormatEmptyStringPipe,
  EnumToValuePipe,
  AgoPipe,
  StrReplacePipe,
  SortPipe,
  FormatDataSizePipe,
  BooleanToTextPipe,
  DurationPipe
} from './pipes';
import { FlexModule, GridModule } from '@angular/flex-layout';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatRippleModule } from '@angular/material/core';
import { MatDialogModule } from '@angular/material/dialog';
import { MatExpansionModule } from '@angular/material/expansion';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatRadioModule } from '@angular/material/radio';
import { MatSelectModule } from '@angular/material/select';
import { MatTooltipModule } from '@angular/material/tooltip';
import { CdkIconModule } from '@cdk-ui/icon';
import { HasAnyAuthorityDirective } from '@core-ui/directives';

@NgModule({
  imports: [
    CommonModule,
    MatDialogModule,
    MatButtonModule,
    MatProgressSpinnerModule,
    MatProgressBarModule,
    MatIconModule,
    MatInputModule,
    MatFormFieldModule,
    MatCheckboxModule,
    MatRippleModule,
    MatTooltipModule,
    MatRadioModule,
    MatSelectModule,
    MatExpansionModule,
    IconModule,
    FlexModule,
    GridModule,
    CdkIconModule,

    // UI
    ConfirmDialogComponent,
    ConfirmDialogStaticComponent,
    SecureDeleteDialogComponent,
    NothingFoundComponent,
    // Pipes
    FormatCpuPipe,
    FormatMemoryPipe,
    FormatEmptyStringPipe,
    EnumToValuePipe,
    AgoPipe,
    StrReplacePipe,
    SortPipe,

    FormatDataSizePipe,
    BooleanToTextPipe,
    DurationPipe,

    HasAnyAuthorityDirective
  ],
  exports: [
    // UI
    ConfirmDialogComponent,
    ConfirmDialogStaticComponent,
    SecureDeleteDialogComponent,
    NothingFoundComponent,
    // Pipes
    FormatCpuPipe,
    FormatMemoryPipe,
    FormatEmptyStringPipe,
    EnumToValuePipe,
    AgoPipe,
    StrReplacePipe,
    SortPipe,

    FormatDataSizePipe,
    BooleanToTextPipe,
    DurationPipe,

    HasAnyAuthorityDirective
  ]
})
export class SharedModule {}
