import { AbstractControl, UntypedFormArray, ValidationErrors } from '@angular/forms';
import { isBlank, isPresent } from '../utils/utils';

export class ArrayValidators {
  /**
   * @description
   * Validator that duplicate control of array.
   *
   * @param caseSensitive?: boolean
   *
   * @usageNotes
   *
   * ### Validate that the field is duplicate
   *
   * ```typescript
   * const array = new FormArray([new FormControl('')], ArrayValidators.controlDuplicateValue.control);
   *
   * console.log(array.errors); // {duplicate: true}
   * console.log(array.controls[0].errors); // {duplicate: true}
   * ```
   *
   * @returns An error map with the `duplicate` property
   * if the validation check fails, otherwise `null`.
   *
   */
  static controlDuplicateValue(caseSensitive: boolean = true) {
    return (formArray: UntypedFormArray): ValidationErrors | null => {
      const values: string[] = formArray.value;
      let errorDuplicate!: boolean;

      if (formArray.controls.length > 1) {
        let isExists!: boolean;
        let newErrors: any;
        formArray.controls.forEach((control: AbstractControl, index: number) => {
          isExists = false;
          if (!isBlank(control.value)) {
            for (let i = 0; i < values.length; i++) {
              if (!isBlank(values[i]) && values[i] === control.value) {
                if (i !== index) {
                  isExists = true;
                  break;
                }
              }
            }

            if (isExists) {
              if (isBlank(control.errors)) {
                newErrors = { duplicate: true };
              } else {
                newErrors = Object.assign(control.errors as any, { duplicate: true });
              }
              if (!errorDuplicate) {
                errorDuplicate = true;
              }
            } else {
              newErrors = control.errors;

              if (isPresent(newErrors)) {
                // delete duplicate error
                delete newErrors['duplicate'];
                if (isBlank(newErrors)) {
                  // {} to undefined/null
                  newErrors = null;
                }
              }
            }
            control.setErrors(newErrors);
          }
        });
      }
      // Clean errors
      if (errorDuplicate) {
        return { duplicate: true };
      }
      return null;
    };
  }

  /**
   * @description
   * Validator that duplicate specific control array of group .
   *
   * @param fieldName: string
   * @param caseSensitive?: boolean
   *
   * @usageNotes
   *
   * ### Validate that the field is duplicate
   *
   * ```typescript
   * const array = new FormArray([new FormGroup({key: ['']})], ArrayValidators.groupControlDuplicateValue(key));
   *
   * console.log(array.errors); // {duplicate: true}
   * console.log(array.controls[0].get('key').errors); // {duplicate: true}
   * ```
   *
   * @returns An error map with the `duplicate` property
   * if the validation check fails, otherwise `null`.
   *
   *
   */
  static groupControlDuplicateValue(fieldName: string, caseSensitive: boolean = true) {
    return (formArray: UntypedFormArray): ValidationErrors | null => {
      const values: any[] = formArray.value;
      let errorDuplicate!: boolean;

      if (formArray.controls.length > 1) {
        let isExists!: boolean;
        let newErrors: any;
        formArray.controls.forEach((formGroup, index: number) => {
          isExists = false;
          const control = formGroup.get(fieldName) as AbstractControl;
          if (!isBlank(control.value)) {
            for (let i = 0; i < values.length; i++) {
              if (!isBlank(values[i][fieldName]) && values[i][fieldName] === control.value) {
                if (i !== index) {
                  isExists = true;
                  break;
                }
              }
            }
            if (isExists) {
              if (isBlank(control.errors)) {
                newErrors = { duplicate: true };
              } else {
                newErrors = Object.assign(control.errors as any, { duplicate: true });
                console.log('not blank', newErrors);
              }

              if (!errorDuplicate) {
                errorDuplicate = true;
              }
            } else {
              newErrors = control.errors;
              if (isPresent(newErrors)) {
                // delete duplicate error
                delete newErrors['duplicate'];
                if (isBlank(newErrors)) {
                  // {} to undefined/null
                  newErrors = null;
                }
              }
            }
            (formGroup.get(fieldName) as AbstractControl).setErrors(newErrors);
          }
        });
      }
      // Clean errors
      if (errorDuplicate) {
        return { duplicate: true };
      }
      return null;
    };
  }
}
