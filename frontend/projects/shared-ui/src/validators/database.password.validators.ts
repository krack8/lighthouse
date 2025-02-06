import { AbstractControl, ValidationErrors } from '@angular/forms';

export const DatabasePasswordValidator = function (control: AbstractControl): ValidationErrors | null {
  const value: string = control.value || '';

  if (!value) {
    return;
  }
  const upperCaseCharacters = /[A-Z]+/g;
  const lowerCaseCharacters = /[a-z]+/g;
  const numberCharacters = /[0-9]+/g;
  const invalidCharacters = /["#$%'()+,-./:;<=>?@[\]^_`{|}~]/;
  const specialCharacters = /[!&*]/;
  function hasWhiteSpace(s) {
    return s.indexOf(' ') >= 0;
  }
  if (upperCaseCharacters.test(value) === false) {
    return { passwordStrength: `Password has to contain Upper case characters` };
  } else if (lowerCaseCharacters.test(value) === false) {
    return { passwordStrength: `Password has to contain lower case characters` };
  } else if (numberCharacters.test(value) === false) {
    return { passwordStrength: `Password has to contain number characters` };
  } else if (invalidCharacters.test(value) === true) {
    return { passwordStrength: `Password can contain only the following special character (!&*)` };
  } else if (specialCharacters.test(value) === false) {
    return { passwordStrength: `Password has to contain special character (!&*)` };
  } else if (hasWhiteSpace(value) === true) {
    return { passwordStrength: `Password Can't Contain Spaces` };
  } else {
    return;
  }
};
