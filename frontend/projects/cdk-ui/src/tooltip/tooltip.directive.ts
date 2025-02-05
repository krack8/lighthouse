import { ComponentRef, Directive, ElementRef, HostListener, inject, Input, TemplateRef } from '@angular/core';
import { ConnectedPosition, Overlay, OverlayRef, PositionStrategy } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { CdkTooltipContentComponent } from './tooltip-content.component';

/**
 * @description HTML structure tooltip
 */
@Directive({
  selector: '[cdkTooltip]',
  standalone: true
})
export class CdkTooltipDirective {
  private overlay = inject(Overlay);
  private elementRef = inject(ElementRef);

  @Input('cdkTooltip') template!: TemplateRef<any>;
  @Input() tooltipPosition: 'above' | 'below' | 'start' | 'end' = 'below'; // Default position

  private overlayRef!: OverlayRef;
  private tooltipContentInstance!: ComponentRef<CdkTooltipContentComponent>;
  private selfVisiting = false;

  @HostListener('mouseenter') show() {
    this.selfVisiting = true;

    if (!this.template || this.overlayRef?.hasAttached()) return;

    const positionStrategy = this.getPositionStrategy();
    this.overlayRef = this.overlay.create({ positionStrategy });

    const componentPortal = new ComponentPortal(CdkTooltipContentComponent);
    this.tooltipContentInstance = this.overlayRef.attach(componentPortal);

    // Pass the template to the tooltip content component
    this.tooltipContentInstance.instance.contentTemplate = this.template;

    // Listen for mouseenter and mouseleave on the overlay
    const tooltipElement = this.tooltipContentInstance.location.nativeElement;
    tooltipElement.addEventListener('mouseleave', () => {
      setTimeout(() => {
        if (!this.selfVisiting) this.hide();
      }, 50);
    });
  }

  @HostListener('mouseleave') hide() {
    this.selfVisiting = false;
    // Delay hiding if the tooltip is hovered
    setTimeout(() => {
      if (this.tooltipContentInstance && !this.tooltipContentInstance.instance.isHovered) {
        this.overlayRef.detach();
      }
    }, 100);
  }

  private getPositionStrategy(): PositionStrategy {
    const positions: { [key: string]: ConnectedPosition } = {
      above: {
        originX: 'center',
        originY: 'top',
        overlayX: 'center',
        overlayY: 'bottom'
      },
      below: {
        originX: 'center',
        originY: 'bottom',
        overlayX: 'center',
        overlayY: 'top'
      },
      start: {
        originX: 'start',
        originY: 'center',
        overlayX: 'end',
        overlayY: 'center'
      },
      end: {
        originX: 'end',
        originY: 'center',
        overlayX: 'start',
        overlayY: 'center'
      }
    };

    const selectedPosition = positions[this.tooltipPosition];
    return this.overlay.position().flexibleConnectedTo(this.elementRef).withPositions([selectedPosition]);
  }
}
