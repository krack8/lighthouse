import { ChangeDetectionStrategy, Component, ElementRef, Inject, Input, NgZone, Optional, ViewContainerRef } from '@angular/core';
import {
  MAT_TOOLTIP_DEFAULT_OPTIONS,
  MAT_TOOLTIP_SCROLL_STRATEGY,
  MatTooltip,
  MatTooltipDefaultOptions,
  TooltipPosition
} from '@angular/material/tooltip';
import { AriaDescriber, FocusMonitor } from '@angular/cdk/a11y';
import { Directionality } from '@angular/cdk/bidi';
import { Overlay, ScrollDispatcher } from '@angular/cdk/overlay';
import { Platform } from '@angular/cdk/platform';
import { DOCUMENT } from '@angular/common';

@Component({
  selector: 'cdk-hint',
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `<svg
    class="flex-shrink-0 inline"
    viewBox="0 0 20 20"
    xmlns="http://www.w3.org/2000/svg"
    [attr.width]="size"
    [attr.height]="size"
    [attr.fill]="color"
  >
    <path
      fill-rule="evenodd"
      d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
      clip-rule="evenodd"
    ></path>
  </svg>`,
  host: {
    class: 'cdk-hint'
  },
  exportAs: 'cdkHint'
})
export class CdkHintComponent extends MatTooltip {
  @Input() get hint() {
    return this.message;
  }
  set hint(value: string) {
    this.message = value;
  }

  @Input() get hintPosition() {
    return this.position;
  }

  set hintPosition(value: TooltipPosition) {
    if (value) {
      this.position = value;
      this.setHintOverlayClass(value);
    }
  }

  @Input() size: string = '1.2rem';
  @Input() color: string = 'currentColor';

  constructor(
    _overlay: Overlay,
    _elementRef: ElementRef,
    _scrollDispatcher: ScrollDispatcher,
    _viewContainerRef: ViewContainerRef,
    _ngZone: NgZone,
    _platform: Platform,
    _ariaDescriber: AriaDescriber,
    _focusMonitor: FocusMonitor,
    @Inject(MAT_TOOLTIP_SCROLL_STRATEGY) _scrollStrategy: any,
    @Optional() _dir: Directionality,
    @Optional()
    @Inject(MAT_TOOLTIP_DEFAULT_OPTIONS)
    _defaultOptions: MatTooltipDefaultOptions,
    @Inject(DOCUMENT) _document: any
  ) {
    super(
      _overlay,
      _elementRef,
      _scrollDispatcher,
      _viewContainerRef,
      _ngZone,
      _platform,
      _ariaDescriber,
      _focusMonitor,
      _scrollStrategy,
      _dir,
      _defaultOptions,
      _document
    );
    this.position = 'after';
    this.setHintOverlayClass(this.position);
  }

  setHintOverlayClass(position: TooltipPosition): void {
    this.tooltipClass = `cdk-hint-tooltip cdk-hint-${position}`;
  }
}
