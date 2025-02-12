import { ElementRef, Injectable, Injector } from '@angular/core';
import { ConnectionPositionPair, Overlay, OverlayConfig, PositionStrategy } from '@angular/cdk/overlay';
import { ComponentPortal, PortalInjector } from '@angular/cdk/portal';
import { PopoverContent, PopoverRef } from './popover-ref';
import { PopoverComponent } from './popover.component';

export interface PopoverParams<T> {
  width?: string | number;
  height?: string | number;
  origin: ElementRef | HTMLElement;
  position?: ConnectionPositionPair[];
  content: PopoverContent;
  data?: T;
  offsetY?: number;
  offsetX?: number;
}

@Injectable({
  providedIn: 'root'
})
export class PopoverService {
  constructor(
    private overlay: Overlay,
    private injector: Injector
  ) {}

  open<T>({ origin, content, data, width, height, position, offsetX, offsetY }: PopoverParams<T>): PopoverRef<T> {
    const overlayRef = this.overlay.create(this.getOverlayConfig({ origin, width, height, position, offsetX, offsetY }));
    const popoverRef = new PopoverRef<T>(overlayRef, content, data as T);

    const injector = this.createInjector(popoverRef, this.injector);
    overlayRef.attach(new ComponentPortal(PopoverComponent, null, injector));

    return popoverRef;
  }

  createInjector(popoverRef: PopoverRef, injector: Injector) {
    const tokens = new WeakMap([[PopoverRef, popoverRef]]);
    return new PortalInjector(injector, tokens);
  }

  private getOverlayConfig({ origin, width, height, position, offsetX, offsetY }: any): OverlayConfig {
    return new OverlayConfig({
      hasBackdrop: true,
      width,
      height,
      backdropClass: 'popover-backdrop',
      positionStrategy: this.getOverlayPosition({ origin, position, offsetX, offsetY }),
      scrollStrategy: this.overlay.scrollStrategies.reposition()
    });
  }

  private getOverlayPosition({ origin, position, offsetX, offsetY }: any): PositionStrategy {
    const positionStrategy = this.overlay
      .position()
      .flexibleConnectedTo(origin)
      .withPositions(position || this.getPositions())
      .withFlexibleDimensions(true)
      .withDefaultOffsetY(offsetY || 0)
      .withDefaultOffsetX(offsetX || 0)
      .withTransformOriginOn('.popover')
      .withPush(true);

    return positionStrategy;
  }

  private getPositions(): ConnectionPositionPair[] {
    return [
      {
        originX: 'center',
        originY: 'top',
        overlayX: 'center',
        overlayY: 'bottom'
      },
      {
        originX: 'center',
        originY: 'bottom',
        overlayX: 'center',
        overlayY: 'top'
      }
    ];
  }
}
