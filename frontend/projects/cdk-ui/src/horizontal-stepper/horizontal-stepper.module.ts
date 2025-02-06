// Imports from @angular
import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { CdkStepperModule } from '@angular/cdk/stepper';
// Component
import { CdkHorizontalStepperComponent } from './horizontal-stepper.component';

@NgModule({
  declarations: [CdkHorizontalStepperComponent],
  imports: [CommonModule, CdkStepperModule],
  exports: [CdkHorizontalStepperComponent]
})
export class CdkHorizontalStepperModule {}
