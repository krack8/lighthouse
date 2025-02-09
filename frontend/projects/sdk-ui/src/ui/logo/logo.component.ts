import { Component, Input, ViewEncapsulation } from '@angular/core';
import { Observable } from 'rxjs';
import { CoreConfigService, ICoreConfig } from '@core-ui/services';

@Component({
  selector: 'kc-logo',
  encapsulation: ViewEncapsulation.None,
  host: {
    class: 'kc-logo'
  },
  template: `
    <ng-container *ngIf="config$ | async as config; else placeholderLoading">
      <img
        [ngStyle]="{ width: logoWidth }"
        [src]="config?.logoUrl || (config?.webTheme?.includes('DARK') ? 'assets/images/logo-dark.svg' : 'assets/images/logo-light.svg')"
      />
      <!-- Power By -->
      <div class="power_by" *ngIf="showPowerBy && (config?.logoUrl || config?.favicon || config?.name)">
        <span class="power_by_label">Powered By: </span>
        <div class="power_by_content">
          <img class="power_by_logo" src="assets/images/ic_klovercloud_logo.png" alt="" />
          <span class="power_by_logo_text"> <span style="color:#5BC4D6">Klover</span><span style="color:#4164A9">Cloud</span> </span>
        </div>
      </div>
    </ng-container>
    <!-- Placeholder Loading -->
    <ng-template #placeholderLoading>Loading</ng-template>
  `,
  styles: [
    `
      .kc-logo {
        display: flex;
        flex-direction: column;
        align-items: center;
      }
      .power_by {
        font-size: 14px;
        padding-top: 10px;
        font-weight: 600;
        text-align: center;
        gap: 0.5rem;
        display: flex;
        align-items: center;
        justify-content: center;
      }
      .power_by_logo {
        margin-right: 0.25rem;
        width: 14px;
      }
      .power_by_logo_text {
        font-weight: 700;
      }
      .power_by_content {
        display: flex;
        align-items: center;
      }
    `
  ]
})
export class LogoComponent {
  @Input() logoWidth: string = '10rem';
  @Input() showPowerBy: boolean = false;

  config$!: Observable<ICoreConfig | null>;

  constructor(private coreConfig: CoreConfigService) {
    this.config$ = coreConfig.generalInfo$;
  }
}
