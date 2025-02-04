import { animate, style, transition, trigger } from '@angular/animations';

export function scaleFadeInAnimation(duration: number) {
  return trigger('scaleFadeIn', [
    transition(':enter', [
      style({
        transform: 'scale(0.8)',
        opacity: 0
      }),
      animate(
        `${duration}ms cubic-bezier(0.35, 0, 0.25, 1)`,
        style({
          transform: 'scale(1)',
          opacity: 1
        })
      )
    ])
  ]);
}

export const scaleFadeIn400ms = scaleFadeInAnimation(400);
