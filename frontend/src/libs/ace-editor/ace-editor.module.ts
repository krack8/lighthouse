import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AceEditorComponent } from './ace-editor.component';

@NgModule({
  declarations: [AceEditorComponent],
  imports: [CommonModule],
  exports: [AceEditorComponent]
})
export class AceEditorModule {}
