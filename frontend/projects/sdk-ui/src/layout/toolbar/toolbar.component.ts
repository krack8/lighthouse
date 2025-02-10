import { Component, HostBinding, Inject, Input, OnInit } from '@angular/core';
import { Platform } from '@angular/cdk/platform';
import { Observable } from 'rxjs';
import { stagger80ms, fadeInUp400ms, scaleIn400ms, fadeInRight400ms } from '@sdk-ui/animations';
import { CoreConfigService } from '@core-ui/services';
import { LayoutService, ToolbarService } from '@sdk-ui/services';

@Component({
  selector: 'kc-toolbar',
  templateUrl: './toolbar.component.html',
  styleUrls: ['./toolbar.component.scss'],
  animations: [stagger80ms, fadeInUp400ms, scaleIn400ms, fadeInRight400ms]
})
export class ToolbarComponent implements OnInit {
  @Input() mobileQuery!: boolean;

  @Input()
  @HostBinding('class.shadow-b')
  hasShadow!: boolean;

  isChrome!: boolean;

  coreConfig$ = this.coreConfigService.generalInfo$;
  toolbarData$: Observable<any> = this.toolbarService.currentData;

  constructor(
    private layoutService: LayoutService,
    private toolbarService: ToolbarService,
    private coreConfigService: CoreConfigService,
    private platform: Platform
  ) {}

  ngOnInit() {
    this.isChrome = this.platform.BLINK;
    setTimeout(() => {
      this.closeWarning();
    }, 7000);
  }

  openQuickpanel() {
    this.layoutService.openQuickpanel();
  }

  openSidenav() {
    this.layoutService.openSidenav();
  }
  closeWarning() {
    this.isChrome = true;
  }
}
