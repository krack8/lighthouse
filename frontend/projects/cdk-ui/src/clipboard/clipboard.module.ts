import { NgModule } from '@angular/core';
import { CdkClipboardComponent } from './clipboard.component';
import { CommonModule } from '@angular/common';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { MatRippleModule } from '@angular/material/core';

@NgModule({
  declarations: [CdkClipboardComponent],
  imports: [CommonModule, MatSnackBarModule, MatRippleModule],
  exports: [CdkClipboardComponent]
})
export class CdkClipboardModule {}
