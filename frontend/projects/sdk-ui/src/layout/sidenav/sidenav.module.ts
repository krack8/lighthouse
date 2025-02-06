import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SidenavComponent } from './sidenav.component';
import { MatToolbarModule } from '@angular/material/toolbar';
import { FlexLayoutModule } from '@angular/flex-layout';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { IconModule } from '@visurel/iconify-angular';
import { SidenavItemModule } from '../sidenav-item/sidenav-item.module';

@NgModule({
  declarations: [SidenavComponent],
  imports: [CommonModule, MatToolbarModule, SidenavItemModule, FlexLayoutModule, MatButtonModule, MatIconModule, IconModule],
  exports: [SidenavComponent]
})
export class SidenavModule {}
