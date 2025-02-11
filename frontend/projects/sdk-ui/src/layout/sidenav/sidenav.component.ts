import { Component, Input, OnDestroy, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { takeWhile } from 'rxjs/operators';
import { trackByRoute } from '@core-ui/utils';
import { CoreConfigService, ICoreConfig } from '@core-ui/services';
import { NavigationService, LayoutService } from '@sdk-ui/services';
import { NavigationItem } from '@sdk-ui/interfaces';

@Component({
  selector: 'kc-sidenav',
  templateUrl: './sidenav.component.html',
  styleUrls: ['./sidenav.component.scss']
})
export class SidenavComponent implements OnInit, OnDestroy {
  isAlive: boolean = true;

  @Input() collapsed!: boolean;
  collapsedOpen$ = this.layoutService.sidenavCollapsedOpen$;

  items$: Observable<NavigationItem[]> = this.navigationService.Items$;
  trackByRoute = trackByRoute;

  config!: ICoreConfig;

  constructor(
    private navigationService: NavigationService,
    private layoutService: LayoutService,
    private coreConfigService: CoreConfigService
  ) {}

  ngOnInit() {
    this.coreConfigService.generalInfo$.pipe(takeWhile(() => this.isAlive)).subscribe(config => {
      this.config = config as ICoreConfig;
    });
  }

  ngOnDestroy(): void {
    this.isAlive = false;
  }

  onMouseEnter() {
    this.layoutService.collapseOpenSidenav();
  }

  onMouseLeave() {
    this.layoutService.collapseCloseSidenav();
  }

  toggleCollapse() {
    this.collapsed ? this.layoutService.expandSidenav() : this.layoutService.collapseSidenav();
  }
}
