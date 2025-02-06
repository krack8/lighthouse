import { Pipe, PipeTransform } from '@angular/core';

export type OrderType = 'asc' | 'desc';

@Pipe({
  standalone: true,
  name: 'sortBy'
})
export class SortPipe implements PipeTransform {
  transform(value: any[], order: OrderType = 'asc'): any[] {
    if (!value) return value; // Empty array
    if (value.length <= 1) return value; // array with only one item
    if (order === 'asc') return value.sort();
    return value.sort().reverse();
  }
}
