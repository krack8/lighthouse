import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { Error500RoutingModule } from './error-500-routing.module';
import { Error500Component } from './error-500.component';
import { FlexLayoutModule } from '@angular/flex-layout';
import { IconModule } from '@visurel/iconify-angular';

@NgModule({
  declarations: [Error500Component],
  imports: [CommonModule, Error500RoutingModule, FlexLayoutModule, IconModule]
})
export class Error500Module {}
