import { Directive } from '@angular/core';

@Directive({
  selector: '[kcPageLayoutContent],kc-page-layout-content',
  host: {
    class: 'kc-page-layout-content'
  }
})
export class PageLayoutContentDirective {
  constructor() {}
}
