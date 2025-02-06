import { animate, style, transition, trigger } from '@angular/animations';

export const scaleInOutAnimation = trigger('scaleInOut', [
  transition(':enter', [
    style({
      transform: 'scale(0)',
      opacity: 0
    }),
    animate(
      '0.2s cubic-bezier(0.35, 0, 0.25, 1)',
      style({
        transform: 'scale(1)',
        opacity: 1
      })
    )
  ]),
  transition(':leave', [
    style({
      transform: 'scale(1)',
      opacity: 1
    }),
    animate(
      '0.2s cubic-bezier(0.35, 0, 0.25, 1)',
      style({
        transform: 'scale(0)',
        opacity: 0
      })
    )
  ])
]);
