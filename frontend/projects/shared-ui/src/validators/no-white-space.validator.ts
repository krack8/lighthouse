import { AbstractControl } from '@angular/forms';

export function NoWhiteSpaceValidator(control: AbstractControl) {
  if (/\s/g.test(control.value)) return { whitespace: true };
  return null;
}
