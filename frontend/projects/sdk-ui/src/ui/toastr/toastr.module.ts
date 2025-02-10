import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { ToastrComponent } from './toastr.component';
import { ToastrDirective } from './toastr.directive';
import { ToastrService } from './toastr.service';

@NgModule({
  imports: [CommonModule],
  declarations: [ToastrComponent, ToastrDirective],
  providers: [ToastrService]
})
export class ToastrModule {}
