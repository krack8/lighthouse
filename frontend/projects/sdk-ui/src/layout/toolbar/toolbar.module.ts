import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { MatRippleModule } from '@angular/material/core';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatMenuModule } from '@angular/material/menu';
import { MatInputModule } from '@angular/material/input';
import { MatDialogModule } from '@angular/material/dialog';
import { MatTooltipModule } from '@angular/material/tooltip';
import { FlexLayoutModule } from '@angular/flex-layout';
import { IconModule } from '@visurel/iconify-angular';

import { ContainerModule } from '@sdk-ui/directives';
import { ToolbarComponent } from './toolbar.component';
import { ToolbarUserModule } from './toolbar-user/toolbar-user.module';
import { NavigationModule } from '../navigation/navigation.module';
import { NavigationItemModule } from '../navigation-item/navigation-item.module';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatSelectModule } from '@angular/material/select';
import { FormsModule } from '@angular/forms';

@NgModule({
  declarations: [ToolbarComponent],
  imports: [
    CommonModule,
    FlexLayoutModule,
    MatButtonModule,
    MatIconModule,
    MatMenuModule,
    MatInputModule,
    MatRippleModule,
    ToolbarUserModule,
    IconModule,
    NavigationModule,
    RouterModule,
    NavigationItemModule,
    ContainerModule,
    MatDialogModule,
    MatTooltipModule,
    MatFormFieldModule,
    MatSelectModule,
    FormsModule
  ],
  exports: [ToolbarComponent],

})
export class ToolbarModule {}
