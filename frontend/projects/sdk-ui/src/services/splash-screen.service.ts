import { Inject, Injectable } from '@angular/core';
import { DOCUMENT } from '@angular/common';
import { animate, AnimationBuilder, style } from '@angular/animations';

@Injectable({
  providedIn: 'root'
})
export class SplashScreenService {
  splashScreenElem!: HTMLElement | null;

  constructor(
    @Inject(DOCUMENT) private document: Document,
    private animationBuilder: AnimationBuilder
  ) {
    this.splashScreenElem = this.document.body.querySelector('#kc-splash-screen');
  }

  hide() {
    const player = this.animationBuilder
      .build([
        style({
          opacity: 1
        }),
        animate(
          '400ms cubic-bezier(0.25, 0.8, 0.25, 1)',
          style({
            opacity: 0
          })
        )
      ])
      .create(this.splashScreenElem);

    player.onDone(() => this.splashScreenElem?.remove());
    player.play();
  }
}
