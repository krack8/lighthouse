import { Inject, Injectable, OnDestroy } from '@angular/core';
import { DOCUMENT } from '@angular/common';
import { BehaviorSubject, Subject } from 'rxjs';
import { filter, takeUntil } from 'rxjs/operators';

export enum Style {
  light = 'kc-style-light',
  dark = 'kc-style-dark',
  lightPink = 'kc-style-pink',
  default = dark
}

@Injectable({
  providedIn: 'root'
})
export class StyleService implements OnDestroy {
  private _destroy$ = new Subject<void>();

  currentStyle = Style.default;

  private _styleSubject = new BehaviorSubject<Style | null>(Style.dark);
  style$ = this._styleSubject.asObservable().pipe(filter(style => !!style));

  constructor(@Inject(DOCUMENT) private document: Document) {
    this.style$.pipe(takeUntil(this._destroy$)).subscribe(style => {
      this._updateStyle(style as Style);
    });
  }

  ngOnDestroy(): void {
    this._destroy$.next();
    this._destroy$.complete();
  }

  setStyle(style: Style) {
    this._styleSubject.next(style);
  }

  getStyle() {
    return this.currentStyle;
  }

  private _updateStyle(style: Style) {
    this.currentStyle = style;

    const body = this.document.body;

    Object.values(Style)
      .filter(s => s !== style || (style === Style.lightPink && s !== Style.light))
      .forEach(value => {
        if (body.classList.contains(value)) {
          body.classList.remove(value);
        }
      });
    if (style === Style.lightPink) {
      body.classList.add(Style.light);
    }
    body.classList.add(style);
    // console.log('style', style)
  }
}
