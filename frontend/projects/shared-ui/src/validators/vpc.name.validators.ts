import { AbstractControl, ValidationErrors } from '@angular/forms';

export class VpcValidator {
  // @ts-ignore
  static name(control: AbstractControl): ValidationErrors | null {
    const nameFormat = /^[A-Za-z0-9 -]+$/;
    if (!(control.value as string)?.match(nameFormat)) {
      return { name: true };
    }
    return null;
  }
}
