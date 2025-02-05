import { AbstractControl, ValidationErrors } from '@angular/forms';

export class InputsValidators {
  static numberInt(control: AbstractControl): ValidationErrors | null {
    const intFormat = /^[0-9]*$/;
    if (!(control.value as string).match(intFormat)) {
      return { numberInt: true };
    }

    return null;
  }
}
