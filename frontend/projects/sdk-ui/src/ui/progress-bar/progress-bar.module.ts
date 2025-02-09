import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ProgressBarComponent } from './progress-bar.component';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { LoadingBarRouterModule } from '../loading-bar';

@NgModule({
  declarations: [ProgressBarComponent],
  imports: [CommonModule, MatProgressBarModule, LoadingBarRouterModule],
  exports: [ProgressBarComponent]
})
export class ProgressBarModule {}
