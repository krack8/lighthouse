// Module
import { NgModule } from '@angular/core';
import { CommonModule, TitleCasePipe } from '@angular/common';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';
import { FlexLayoutModule } from '@angular/flex-layout';
import { IconModule } from '@visurel/iconify-angular';
import { SettingsRoutingModule } from './settings-routing.module';
import { MatButtonModule } from '@angular/material/button';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatDialogModule } from '@angular/material/dialog';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatMenuModule } from '@angular/material/menu';
import { MatPaginatorModule } from '@angular/material/paginator';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatSelectModule } from '@angular/material/select';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { MatSortModule } from '@angular/material/sort';
import { MatTableModule } from '@angular/material/table';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatChipsModule } from '@angular/material/chips';
import { MatListModule } from '@angular/material/list';
import { MatRadioModule } from '@angular/material/radio';
import { ContainerModule } from '@sdk-ui/directives/container/container.module';
import { PageLayoutModule } from '@sdk-ui/ui/page-layout/page-layout.module';
import { SecondaryToolbarModule } from '@sdk-ui/ui/secondary-toolbar/secondary-toolbar.module';
import { SharedModule } from '@shared-ui/shared.module';
import { ToolbarService } from '@sdk-ui/services';
import { Ng2SearchPipeModule } from 'ng2-search-filter';
import { MatExpansionModule } from '@angular/material/expansion';
import { CdkIconModule } from '@cdk-ui/icon';

@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    PageLayoutModule,
    FlexLayoutModule,
    MatInputModule,
    MatPaginatorModule,
    MatTableModule,
    MatSortModule,
    MatCheckboxModule,
    MatIconModule,
    MatButtonModule,
    MatMenuModule,
    MatTooltipModule,
    MatSnackBarModule,
    MatDialogModule,
    MatProgressSpinnerModule,
    MatSelectModule,
    MatChipsModule,
    MatRadioModule,
    MatListModule,
    IconModule,
    SecondaryToolbarModule,
    ContainerModule,
    FormsModule,
    ReactiveFormsModule,
    Ng2SearchPipeModule,
    SharedModule,
    SettingsRoutingModule,
    MatProgressBarModule,
    MatExpansionModule,
    CdkIconModule
  ],
  providers: [TitleCasePipe]
})
export class SettingsModule {
  constructor(private _toolbarService: ToolbarService) {
    this._toolbarService.changeData({ title: 'Settings' });
  }
}
