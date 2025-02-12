import { AbstractControl, ValidationErrors } from '@angular/forms';

export class SpaceValidator {
  static noLeadingSpace(control: AbstractControl): ValidationErrors | null {
    const value = control.value || '';
    const isLeadingSpace = value.startsWith(' ');
    return isLeadingSpace ? { leadingSpace: true } : null;
  }
}
