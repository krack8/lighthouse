import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PageLayoutComponent } from './page-layout.component';
import { PageLayoutHeaderDirective } from './page-layout-header.directive';
import { PageLayoutContentDirective } from './page-layout-content.directive';

@NgModule({
  imports: [CommonModule],
  declarations: [PageLayoutComponent, PageLayoutHeaderDirective, PageLayoutContentDirective],
  exports: [PageLayoutComponent, PageLayoutHeaderDirective, PageLayoutContentDirective]
})
export class PageLayoutModule {}
