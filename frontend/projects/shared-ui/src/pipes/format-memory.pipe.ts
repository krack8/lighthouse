import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  standalone: true,
  name: 'formatMemory'
})
export class FormatMemoryPipe implements PipeTransform {
  transform(memory: number): string {
    let unit = ' GB';
    memory = memory / 1024;
    if (memory > 1024) {
      unit = ' TB';
      memory = memory / 1024;
    }
    if (memory % 1 === 0) {
      return memory.toFixed(0) + unit;
    }
    return memory.toFixed(2) + unit;
  }
}
