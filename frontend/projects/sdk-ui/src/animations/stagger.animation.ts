import { animateChild, query, stagger, transition, trigger } from '@angular/animations';

export function staggerAnimation(timing: number) {
  return trigger('stagger', [
    transition('* => *', [
      // each time the binding value changes
      query('@fadeInUp, @fadeInRight, @scaleIn', stagger(timing, animateChild()), { optional: true })
    ])
  ]);
}

export const stagger80ms = staggerAnimation(80);
export const stagger60ms = staggerAnimation(60);
export const stagger40ms = staggerAnimation(40);
export const stagger20ms = staggerAnimation(20);
