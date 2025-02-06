/*
 * Definition: validate duplicate value control from FormArray
 *
 * Usage: this.fb.array([], [CustomUniqueValidators('appVpcId'), CustomUniqueValidators('mountPath')])
 *
 */

import { AbstractControl, UntypedFormArray, UntypedFormGroup, ValidationErrors, ValidatorFn } from '@angular/forms';
import { isPresent, isBlank } from '../utils';

export const CustomUniqueValidators = (field: string, caseSensitive: boolean = true) => {
  return (formArray: UntypedFormArray): ValidationErrors | null => {
    const controls: AbstractControl[] = formArray.controls.filter((formGroup: AbstractControl) => {
      return isPresent((formGroup.get(field) as AbstractControl).value);
    });
    const uniqueObj: any = { uniqueBy: true };
    let find: boolean = false;

    if (controls.length > 1) {
      for (let i: number = 0; i < controls.length; i++) {
        const formGroup: UntypedFormGroup | any = controls[i];
        const mainControl: AbstractControl = formGroup.get(field);

        const val: string = mainControl.value;

        const mainValue: string = caseSensitive ? val.toLowerCase() : val;
        controls.forEach((group, index: number) => {
          if (i === index) {
            // Same group
            return;
          }

          const currControl: any = group.get(field);
          const tempValue: string = currControl.value;
          const currValue: string = caseSensitive ? tempValue.toLowerCase() : tempValue;
          let newErrors: any;

          if (mainValue === currValue) {
            if (isBlank(currControl.errors)) {
              newErrors = uniqueObj;
            } else {
              newErrors = Object.assign(currControl.errors, uniqueObj);
            }

            find = true;
          } else {
            newErrors = currControl.errors;

            if (isPresent(newErrors)) {
              // delete uniqueBy error
              delete newErrors['uniqueBy'];
              if (isBlank(newErrors)) {
                // {} to undefined/null
                newErrors = null;
              }
            }
          }

          // Add specific errors based on condition
          currControl.setErrors(newErrors);
        });
      }

      if (find) {
        // Set errors to whole formArray
        return uniqueObj;
      }
    }

    // Clean errors
    return null;
  };
};
