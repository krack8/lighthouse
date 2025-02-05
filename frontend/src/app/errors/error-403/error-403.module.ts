import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { Error403RoutingModule } from './error-403-routing.module';
import { Error403Component } from './error-403.component';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { IconModule } from '@visurel/iconify-angular';

@NgModule({
  declarations: [Error403Component],
  imports: [CommonModule, Error403RoutingModule, MatButtonModule, MatIconModule, IconModule]
})
export class Error403Module {}
