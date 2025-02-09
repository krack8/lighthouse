import { Component, HostBinding, Input, ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'kc-page-layout',
  template: '<ng-content></ng-content>',
  host: {
    class: 'kc-page-layout'
  },
  encapsulation: ViewEncapsulation.None,
  styleUrls: ['./page-layout.component.scss']
})
export class PageLayoutComponent {
  @Input() mode: 'card' | 'simple' = 'simple';

  @HostBinding('class.kc-page-layout-card')
  get isCard() {
    return this.mode === 'card';
  }

  @HostBinding('class.kc-page-layout-simple')
  get isSimple() {
    return this.mode === 'simple';
  }
}
