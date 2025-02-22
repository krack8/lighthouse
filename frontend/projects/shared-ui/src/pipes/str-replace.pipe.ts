import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  standalone: true,
  name: 'strReplace'
})
export class StrReplacePipe implements PipeTransform {
  transform(value: string, strToReplace: string, replacementStr: string = ' '): string {
    return value?.replace(new RegExp(strToReplace, 'g'), replacementStr);
  }
}
