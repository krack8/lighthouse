import { animate, group, style, transition, trigger } from '@angular/animations';

export const popoverAnimation = trigger('transformPopover', [
  transition(':enter', [
    style({
      opacity: 0,
      transform: 'scale(0.6)'
    }),
    group([
      animate(
        '100ms linear',
        style({
          opacity: 1
        })
      ),
      animate(
        '150ms cubic-bezier(0, 0, 0.2, 1)',
        style({
          transform: 'scale(1)'
        })
      )
    ])
  ]),
  transition(':leave', [
    style({
      opacity: 1
    }),
    animate(
      '100ms linear',
      style({
        opacity: 0
      })
    )
  ])
]);
