import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ContainerDirective } from './container.directive';

@NgModule({
  declarations: [ContainerDirective],
  imports: [CommonModule],
  exports: [ContainerDirective]
})
export class ContainerModule {}
