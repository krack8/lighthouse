import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  standalone: true,
  name: 'formatCpu'
})
export class FormatCpuPipe implements PipeTransform {
  transform(cpu: number): string {
    if (cpu === undefined || cpu === 0 || cpu === null) {
      return '0';
    }
    cpu = cpu / 1000;
    if (cpu % 1 === 0) {
      return cpu.toFixed(0);
    }
    return cpu.toFixed(2);
  }
}
