import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NavigationComponent } from './navigation.component';
import { FlexLayoutModule } from '@angular/flex-layout';
import { MatRippleModule } from '@angular/material/core';
import { MatMenuModule } from '@angular/material/menu';
import { MatIconModule } from '@angular/material/icon';
import { IconModule } from '@visurel/iconify-angular';
import { RouterModule } from '@angular/router';
import { ContainerModule } from '@sdk-ui/directives';
import { NavigationItemModule } from '../navigation-item/navigation-item.module';

@NgModule({
  declarations: [NavigationComponent],
  imports: [
    CommonModule,
    FlexLayoutModule,
    MatRippleModule,
    MatMenuModule,
    MatIconModule,
    IconModule,
    RouterModule,
    NavigationItemModule,
    ContainerModule
  ],
  exports: [NavigationComponent]
})
export class NavigationModule {}
