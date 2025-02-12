import { Component } from '@angular/core';
import { Toastr } from './toastr.interface';
import { toasts } from './toastr.data';
import { ToastrService } from './toastr.service';

@Component({
  selector: 'kc-toastr',
  templateUrl: './toastr.component.html',
  styleUrls: ['./toastr.component.scss']
})
export class ToastrComponent {
  toasts: Toastr[] = toasts;

  constructor(private toastrService: ToastrService) {}

  onDismiss(index: number): void {
    this.toastrService.dismiss(index);
  }
}
