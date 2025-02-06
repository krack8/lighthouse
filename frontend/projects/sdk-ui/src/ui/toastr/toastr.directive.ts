import { Directive, Input, HostListener } from '@angular/core';
import { ToastrService } from './toastr.service';

@Directive({ selector: '[closeToastr]' })
export class ToastrDirective {
  @Input() closeToastr!: number;

  constructor(private toastrService: ToastrService) {}

  @HostListener('click') onClose() {
    this.toastrService.dismiss(this.closeToastr);
  }
}
