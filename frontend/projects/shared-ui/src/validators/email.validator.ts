import { AbstractControl, ValidationErrors, ValidatorFn } from '@angular/forms';

/*
Example of valid search email id
  - mysite@ourearth.com
  - my.ownsite@ourearth.org
  - mysite@you.me.net

Example of invalid search email id
  - mysite.ourearth.com [@ is not present]
  - mysite@.com.my [ tld (Top Level domain) can not start with dot "." ]
  - @you.me.net [ No character before @ ]
  - mysite123@gmail.b [ ".b" is not a valid tld ]
  - mysite@.org.org [ tld can not start with dot "." ]
  - .mysite@mysite.org [ an email should not be start with "." ]
  - mysite()*@gmail.com [ here the regular expression only allows character, digit, underscore, and dash ]
  - mysite..1234@yahoo.com [double dots are not allowed]
*/
export const EmailValidator = function (control: AbstractControl): ValidationErrors | null {
  const value = control.value;
  // eslint-disable-next-line no-useless-escape
  const validatorRegex = /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/;

  if (!value) {
    return null;
  }

  if (!validatorRegex.test(value)) {
    return { email: true };
  }
  return null;
};

export const EmailDomainValidator = function (domains: string[] | null, forceValidation?: boolean): ValidatorFn {
  return (control: AbstractControl): ValidationErrors | null => {
    // console.log("control", control)
    const val: string = control.value;

    if (!val) {
      return null;
    }

    // Empty validation main
    if (!domains?.length) {
      if (forceValidation) {
        return { domain: 'The Validation domain was not provided' };
      }
      return null;
    }

    const valSplit: string[] = val.split('@');
    if (!domains.includes(valSplit[valSplit.length - 1])) {
      return { domain: 'Email domain is invalid!' };
    }
    return null;
  };
};
