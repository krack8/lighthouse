import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  standalone: true,
  name: 'formatEmptyString'
})
export class FormatEmptyStringPipe implements PipeTransform {
  transform(value: string): string {
    if (value === undefined || value === null || value === '') return 'n/a';

    return value;
  }
}
