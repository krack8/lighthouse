import { animate, style, transition, trigger } from '@angular/animations';

export function fadeInRightAnimation(duration: number) {
  return trigger('fadeInRight', [
    transition(':enter', [
      style({
        transform: 'translateX(-20px)',
        opacity: 0
      }),
      animate(
        `${duration}ms cubic-bezier(0.35, 0, 0.25, 1)`,
        style({
          transform: 'translateX(0)',
          opacity: 1
        })
      )
    ])
  ]);
}

export const fadeInRight400ms = fadeInRightAnimation(400);
