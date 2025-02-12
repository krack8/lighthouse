import { AbstractControl, ValidationErrors } from '@angular/forms';

export const PasswordValidator = function (control: AbstractControl): ValidationErrors | null {
  const value: string = control.value || '';

  if (!value) {
    return null;
  }

  const upperCaseCharacters = /[A-Z]+/g;
  if (upperCaseCharacters.test(value) === false) {
    return { passwordStrength: `Password has to contain Uppercase characters` };
  }

  const lowerCaseCharacters = /[a-z]+/g;
  if (lowerCaseCharacters.test(value) === false) {
    return { passwordStrength: `Password has to contain lowercase characters` };
  }

  const numberCharacters = /[0-9]+/g;
  if (numberCharacters.test(value) === false) {
    return { passwordStrength: `Password has to contain number characters` };
  }

  const specialCharacters = /[!&*@^#]/;
  if (specialCharacters.test(value) === false) {
    return { passwordStrength: `Password has to contain special character (!&*@^#)` };
  }

  const invalidCharacters = /["#$%'()+,-./:;<=>?[\]_`{|}~]/; // Added to validator: @^#
  if (invalidCharacters.test(value) === true) {
    return { passwordStrength: `Password can contain only the following special character (!&*@^#)` };
  }

  function hasWhiteSpace(s: any) {
    return s.indexOf(' ') >= 0;
  }
  if (hasWhiteSpace(value) === true) {
    return { passwordStrength: `Password Can't Contain Spaces` };
  }
  return null;
};
