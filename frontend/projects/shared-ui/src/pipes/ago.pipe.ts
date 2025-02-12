import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  standalone: true,
  name: 'ago'
})
export class AgoPipe implements PipeTransform {
  transform(datetime: any, args?: any): string {
    if (!datetime) {
      return 'a long time ago';
    }

    let time = (Date.now() - Date.parse(datetime)) / 1000;

    if (time < 10) {
      return 'just now';
    } else if (time < 60) {
      return 'a moment ago';
    }

    const divider = [60, 60, 24, 30, 12];
    const string = [' second', ' minute', ' hour', ' day', ' month', ' year'];

    let i;
    for (i = 0; Math.floor(time / divider[i]) > 0; i++) {
      time /= divider[i];
    }
    const plural = Math.floor(time) > 1 ? 's' : '';

    return Math.floor(time) + string[i] + plural + ' ago';
  }
}
