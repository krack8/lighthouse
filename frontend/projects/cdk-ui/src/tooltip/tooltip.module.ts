import { NgModule } from '@angular/core';
import { CdkTooltipDirective } from './tooltip.directive';
import { CdkTooltipContentComponent } from './tooltip-content.component';

@NgModule({
  imports: [CdkTooltipDirective, CdkTooltipContentComponent],
  exports: [CdkTooltipDirective, CdkTooltipContentComponent]
})
export class CdkTooltipModule {}
