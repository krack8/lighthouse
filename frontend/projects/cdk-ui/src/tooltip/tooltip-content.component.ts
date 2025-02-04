import { Component, HostBinding, HostListener, Input } from '@angular/core';
import { CommonModule } from '@angular/common';
import { trigger, transition, style, animate } from '@angular/animations';

@Component({
  selector: 'cdk-tooltip-content',
  standalone: true,
  imports: [CommonModule],
  template: `<div class="cdk-tooltip-content">
    <ng-container *ngTemplateOutlet="contentTemplate"></ng-container>
  </div>`,
  animations: [
    trigger('tooltipAnimation', [
      transition(':enter', [
        style({ opacity: 0, transform: 'scale(0.9)' }),
        animate('200ms ease-out', style({ opacity: 1, transform: 'scale(1)' }))
      ]),
      transition(':leave', [animate('200ms ease-in', style({ opacity: 0, transform: 'scale(0.9)' }))])
    ])
  ],
  host: {
    class: 'cdk-tooltip-container'
  }
})
export class CdkTooltipContentComponent {
  @Input() contentTemplate!: any;

  isHovered = false;

  @HostBinding('@tooltipAnimation') animation = true;

  @HostListener('mouseenter') onMouseEnter() {
    this.isHovered = true;
  }

  @HostListener('mouseleave') onMouseLeave() {
    this.isHovered = false;
    console.log('leaved', this.isHovered);
  }
}
