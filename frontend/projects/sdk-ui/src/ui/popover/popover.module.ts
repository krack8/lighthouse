import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PopoverComponent } from './popover.component';
import { OverlayModule } from '@angular/cdk/overlay';

@NgModule({
  declarations: [PopoverComponent],
  imports: [CommonModule, OverlayModule]
})
export class PopoverModule {}
