import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SafeStylePipe } from './safe-style.pipe';

@NgModule({
  declarations: [SafeStylePipe],
  imports: [CommonModule],
  exports: [SafeStylePipe]
})
export class SafeStyleModule {}
