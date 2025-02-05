import { ChangeDetectionStrategy, Component, Input, ViewEncapsulation } from '@angular/core';
import { IconType } from './icon.interfaces';

@Component({
  selector: 'cdk-icon',
  encapsulation: ViewEncapsulation.None,
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `
    <ng-container [ngSwitch]="type">
      <ng-container *ngSwitchCase="'check'">
        <img src="assets/img/checked.svg" />
      </ng-container>
      <!-- Regular Icon -->
      <ng-container *ngSwitchDefault>
        <ng-container *ngIf="src; else content"><img class="cdk-icon-img" [src]="src" [alt]="alt || 'icon'" /></ng-container>
        <ng-template #content><ng-content></ng-content></ng-template>
      </ng-container>
    </ng-container>
  `,
  host: {
    class: 'cdk-icon',
    '[class.cdk-icon-regular]': "type === 'regular'",
    '[class.cdk-icon-check]': "type === 'check'",
    // Style
    '[style.background]': 'backgroundColor',
    '[style.width]': 'width',
    '[style.height]': 'height',
    '[style.borderRadius]': 'borderRadius'
  }
})
export class CdkIconComponent {
  @Input() type: IconType = 'regular';
  // Icon Img
  @Input() src?: string;
  @Input() alt?: string;
  // Styles
  @Input() backgroundColor?: string;
  @Input() width?: string;
  @Input() height?: string;
  @Input() borderRadius?: string;
}
