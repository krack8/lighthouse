import { Component, Input, OnDestroy, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { takeWhile } from 'rxjs/operators';
import { trackByRoute } from '@core-ui/utils';
import { CoreConfigService, ICoreConfig } from '@core-ui/services';
import { NavigationService, LayoutService } from '@sdk-ui/services';
import { NavigationItem } from '@sdk-ui/interfaces';
import { Router } from '@angular/router';

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

  config!: ICoreConfig;

  constructor(
    private navigationService: NavigationService,
    private layoutService: LayoutService,
    private coreConfigService: CoreConfigService,
    private router: Router
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

  navigateToHome(): void {
    this.router.navigate(['/']); // Change '/home' to your desired route
  }

  /**
   * trackByRoute function for *ngFor directive in this component to track link routes aswell as child routes.
   */
  trackByRoute<T extends { route?: string | string[]; children?: { route: string | string[] }[] }>(index: number, item: T ): string | string[] | undefined {
    if (item.route) {
      return item.route;
    }
    if (item.children && item.children.length > 0) {
      return item.children.map(child => child.route).join(',');
    }
    return undefined;
  }
}
