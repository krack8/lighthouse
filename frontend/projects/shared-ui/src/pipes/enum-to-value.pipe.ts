import { Pipe, PipeTransform } from '@angular/core';

@Pipe({ standalone: true, name: 'enumToValue' })
export class EnumToValuePipe implements PipeTransform {
  transform(value: string): string {
    if (value != null) {
      value = value.replace(/_/g, ' ').toLowerCase();
      value = value.replace(/-/g, ' ').toLowerCase();
    }
    return value;
  }
}
