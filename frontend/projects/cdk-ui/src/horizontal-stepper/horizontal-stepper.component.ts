import { Component, Input } from '@angular/core';
import { CdkStepper } from '@angular/cdk/stepper';

@Component({
  selector: 'cdk-horizontal-stepper',
  templateUrl: './horizontal-stepper.component.html',
  providers: [{ provide: CdkStepper, useExisting: CdkHorizontalStepperComponent }]
})
export class CdkHorizontalStepperComponent extends CdkStepper {
  @Input() header!: string;

  onClick(index: number): void {
    this.selectedIndex = index;
  }
}
