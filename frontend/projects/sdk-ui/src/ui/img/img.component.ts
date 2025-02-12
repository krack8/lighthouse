import { Component, Input, OnChanges, SimpleChanges } from '@angular/core';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { Style, StyleService } from '@sdk-ui/services';

@Component({
  selector: 'kc-img',
  template: ` <img [src]="src$ | async" /> `
})
export class KcImg implements OnChanges {
  @Input() src!: string;
  @Input() darkSrc?: string;
  @Input() lightSrc?: string;

  src$!: Observable<string>;

  constructor(private styleService: StyleService) {}

  ngOnChanges(changes: SimpleChanges): void {
    // console.log('change', changes)
    // if (['src', 'darkSrc', 'lightSrc'].some(item => Object.keys(changes).includes(item))) {
    //   console.log()
    // }
    this.src$ = this.styleService.style$.pipe(
      map(style => {
        // console.log('style', style)
        if (style === Style.light || style === Style.lightPink) {
          return this.lightSrc || this.src;
        }
        return this.darkSrc || this.src;
      })
    );
  }
}
