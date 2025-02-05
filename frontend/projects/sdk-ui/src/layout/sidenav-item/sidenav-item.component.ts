import {
  ChangeDetectionStrategy,
  ChangeDetectorRef,
  Component,
  HostBinding,
  Input,
  OnChanges,
  OnDestroy,
  OnInit,
  SimpleChanges
} from '@angular/core';
import { NavigationEnd, Router } from '@angular/router';
import { filter, takeUntil } from 'rxjs/operators';
import { Subject } from 'rxjs';

import { NavigationService } from '@sdk-ui/services';
import { NavigationDropdown, NavigationItem, NavigationLink } from '@sdk-ui/interfaces';
import { dropdownAnimation } from '@sdk-ui/animations';

@Component({
  selector: 'kc-sidenav-item',
  templateUrl: './sidenav-item.component.html',
  styleUrls: ['./sidenav-item.component.scss'],
  animations: [dropdownAnimation],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class SidenavItemComponent implements OnInit, OnChanges, OnDestroy {
  private _destroy$ = new Subject<void>();
  @Input() item: NavigationItem;
  @Input() level!: number;
  isOpen!: boolean;
  isActive!: boolean;

  isLink = this.navigationService.isLink;
  isDropdown = this.navigationService.isDropdown;
  isSubheading = this.navigationService.isSubheading;

  constructor(
    private router: Router,
    private cd: ChangeDetectorRef,
    private navigationService: NavigationService
  ) {}

  @HostBinding('class')
  get levelClass() {
    return `item-level-${this.level}`;
  }

  ngOnInit() {
    this.router.events
      .pipe(
        filter(event => event instanceof NavigationEnd),
        filter(() => this.isDropdown(this.item)),
        takeUntil(this._destroy$)
      )
      .subscribe(() => this.onRouteChange());

    this.navigationService.openChange$
      .pipe(
        filter(() => this.isDropdown(this.item)),
        takeUntil(this._destroy$)
      )
      .subscribe(item => this.onOpenChange(item as NavigationDropdown));
  }

  ngOnChanges(changes: SimpleChanges): void {
    this._destroy$.next();
    this._destroy$.complete();
    // eslint-disable-next-line no-prototype-builtins
    if (changes && changes.hasOwnProperty('item') && this.isDropdown(this.item)) {
      this.onRouteChange();
    }
  }

  toggleOpen() {
    this.isOpen = !this.isOpen;
    this.navigationService.triggerOpenChange(this.item as NavigationDropdown);
    this.cd.markForCheck();
  }

  onOpenChange(item: NavigationDropdown) {
    if (this.isChildrenOf(this.item as NavigationDropdown, item)) {
      return;
    }

    if (this.hasActiveChilds(this.item as NavigationDropdown)) {
      return;
    }

    if (this.item !== item) {
      this.isOpen = false;
      this.cd.markForCheck();
    }
  }

  onRouteChange() {
    if (this.hasActiveChilds(this.item as NavigationDropdown)) {
      this.isActive = true;
      this.isOpen = true;
      this.navigationService.triggerOpenChange(this.item as NavigationDropdown);
      this.cd.markForCheck();
    } else {
      this.isActive = false;
      this.isOpen = false;
      this.navigationService.triggerOpenChange(this.item as NavigationDropdown);
      this.cd.markForCheck();
    }
  }

  isChildrenOf(parent: NavigationDropdown, item: NavigationDropdown): boolean {
    if (!parent.children?.length) return false;

    if (parent.children.indexOf(item) !== -1) {
      return true;
    }

    return parent.children.filter(child => this.isDropdown(child)).some(child => this.isChildrenOf(child as NavigationDropdown, item));
  }

  hasActiveChilds(parent: NavigationDropdown): boolean {
    if (!parent.children?.length) return false;

    return parent.children.some(child => {
      if (this.isDropdown(child)) {
        return this.hasActiveChilds(child);
      }

      if (this.isLink(child) && !this.isFunction(child.route)) {
        return this.router.isActive(child.route as string, false);
      }
      return false;
    });
  }

  isFunction(prop: NavigationLink['route']) {
    return prop instanceof Function;
  }

  ngOnDestroy(): void {}
}
