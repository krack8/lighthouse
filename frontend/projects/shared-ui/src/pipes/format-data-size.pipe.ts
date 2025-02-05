import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  standalone: true,
  name: 'formatDataSize'
})
export class FormatDataSizePipe implements PipeTransform {
  transform(size: number, ...list: any[]): string {
    if (size > 0) {
      const unit = 1024;
      const i = Math.floor(Math.log(size) / Math.log(1024));
      return ((size / Math.pow(unit, i)).toFixed(2) as any) * 1 + ' ' + [...list][i];
    } else {
      return '0 B';
    }
  }
}
