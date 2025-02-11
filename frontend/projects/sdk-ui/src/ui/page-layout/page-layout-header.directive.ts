import { Directive } from '@angular/core';

@Directive({
  selector: '[kcPageLayoutHeader],kc-page-layout-header',
  host: {
    class: 'kc-page-layout-header'
  }
})
export class PageLayoutHeaderDirective {
  constructor() {}
}
