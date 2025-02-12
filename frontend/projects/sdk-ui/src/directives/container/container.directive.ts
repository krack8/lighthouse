import { Directive, HostBinding } from '@angular/core';

/**
 * @description container can be fix width as box
 * @deprecated box container styles removed
 */
@Directive({
  selector: '[kcContainer]'
})
export class ContainerDirective {
  @HostBinding('class.container') enabled!: boolean;
}
