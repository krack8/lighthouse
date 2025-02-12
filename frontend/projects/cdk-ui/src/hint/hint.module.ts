import { NgModule } from '@angular/core';
import { CdkHintComponent } from './hint.component';
import { CommonModule } from '@angular/common';
import { MatTooltipModule } from '@angular/material/tooltip';

@NgModule({
  declarations: [CdkHintComponent],
  imports: [CommonModule, MatTooltipModule],
  exports: [CdkHintComponent]
})
export class CdkHintModule {}
