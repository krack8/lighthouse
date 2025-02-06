import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { MatIconModule } from '@angular/material/icon';
import { MatRippleModule } from '@angular/material/core';
import { MatMenuModule } from '@angular/material/menu';
import { MatButtonModule } from '@angular/material/button';
import { MatTooltipModule } from '@angular/material/tooltip';
import { FlexLayoutModule } from '@angular/flex-layout';
import { IconModule } from '@visurel/iconify-angular';
import { ToolbarUserComponent } from './toolbar-user.component';
import { ToolbarUserDropdownComponent } from './toolbar-user-dropdown/toolbar-user-dropdown.component';
import { SafeStyleModule } from '@sdk-ui/pipes';

@NgModule({
  declarations: [ToolbarUserComponent, ToolbarUserDropdownComponent],
  imports: [
    CommonModule,
    FlexLayoutModule,
    MatIconModule,
    MatRippleModule,
    MatMenuModule,
    MatButtonModule,
    SafeStyleModule,
    RouterModule,
    MatTooltipModule,
    IconModule
  ],
  exports: [ToolbarUserComponent]
})
export class ToolbarUserModule {}
