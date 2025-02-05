import { Component, Input, OnInit } from '@angular/core';
import { filter, map, startWith } from 'rxjs/operators';
import { NavigationEnd, Router } from '@angular/router';
import { NavigationService } from '@sdk-ui/services';
import { NavigationItem, NavigationLink } from '@sdk-ui/interfaces';
import { trackByRoute } from '@core-ui/utils';

@Component({
  selector: 'kc-navigation-item',
  templateUrl: './navigation-item.component.html',
  styleUrls: ['./navigation-item.component.scss']
})
export class NavigationItemComponent implements OnInit {
  @Input() item: NavigationItem;

  isActive$ = this.router.events.pipe(
    filter(event => event instanceof NavigationEnd),
    startWith(false),
    map(() => this.hasActiveChilds(this.item))
  );

  isLink = this.navigationService.isLink;
  isDropdown = this.navigationService.isDropdown;
  isSubheading = this.navigationService.isSubheading;
  trackByRoute = trackByRoute;

  constructor(
    private navigationService: NavigationService,
    private router: Router
  ) {}

  ngOnInit() {}

  hasActiveChilds(parent: NavigationItem): boolean {
    if (this.isLink(parent)) {
      return this.router.isActive(parent.route as string, false);
    }

    if (this.isDropdown(parent) || this.isSubheading(parent)) {
      if (!parent?.children?.length) return false;
      return parent?.children?.some(child => {
        if (this.isDropdown(child)) {
          return this.hasActiveChilds(child);
        }
        if (this.isLink(child) && !this.isFunction(child.route)) {
          return this.router.isActive(child.route as string, false);
        }

        return false;
      });
    }

    return false;
  }

  isFunction(prop: NavigationLink['route']) {
    return prop instanceof Function;
  }
}
