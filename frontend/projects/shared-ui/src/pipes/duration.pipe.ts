import { Pipe, PipeTransform } from '@angular/core';
import moment from 'moment';

@Pipe({
  name: 'duration',
  standalone: true
})
export class DurationPipe implements PipeTransform {
  transform(duration: any, ...args: any[]): any {
    // console.log(duration, 'DURATION');
    if (duration) {
      let dString = '';
      const d = moment.duration(duration);
      dString += d.years() > 0 ? d.years() + ' Year' + (d.years() > 1 ? 's, ' : ', ') : '';
      dString += d.months() > 0 ? d.months() + ' Month' + (d.months() > 1 ? 's, ' : ', ') : '';
      dString += d.days() > 0 ? d.days() + ' Days' + (d.days() > 1 ? 's, ' : ', ') : '';
      dString += d.hours() > 0 ? d.hours() + ' Hours' + (d.hours() > 1 ? 's, ' : ', ') : '';
      dString += d.minutes() > 0 ? d.minutes() + ' Minute' + (d.minutes() > 1 ? 's, ' : ', ') : '';
      dString += d.seconds() > 0 ? d.seconds() + ' Second' + (d.seconds() > 1 ? 's' : '') : '';
      return dString;
    }
    return 'N/A';
  }
}
