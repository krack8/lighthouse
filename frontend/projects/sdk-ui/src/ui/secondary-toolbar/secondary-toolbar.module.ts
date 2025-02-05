import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SecondaryToolbarComponent } from './secondary-toolbar.component';
import { FlexLayoutModule } from '@angular/flex-layout';
import { IconModule } from '@visurel/iconify-angular';
import { RouterModule } from '@angular/router';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { ContainerModule } from '@sdk-ui/directives';

@NgModule({
  declarations: [SecondaryToolbarComponent],
  imports: [CommonModule, FlexLayoutModule, IconModule, RouterModule, MatButtonModule, MatIconModule, ContainerModule],
  exports: [SecondaryToolbarComponent]
})
export class SecondaryToolbarModule {}
