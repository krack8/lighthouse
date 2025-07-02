import { AfterViewInit, ChangeDetectorRef, Component, Inject, OnDestroy, OnInit, Optional, ViewChild } from '@angular/core';
import { DOCUMENT } from '@angular/common';
import { MatSidenav, MatSidenavContainer } from '@angular/material/sidenav';
import { NavigationEnd, Router, Scroll } from '@angular/router';
import { distinctUntilChanged, filter, map, startWith, take, takeUntil } from 'rxjs/operators';
import { Subject, combineLatest } from 'rxjs';
import { MediaObserver } from '@angular/flex-layout';
import { checkRouterChildsData } from '@sdk-ui/utils';
import { PermissionService, RequesterService } from '@core-ui/services';
import { LayoutService, NavigationService } from '@sdk-ui/services';
import { SidenavLink } from '@sdk-ui/interfaces';
import { SdkConfigService } from '@sdk-ui/services';
import { SelectedClusterService } from '@core-ui/services/selected-cluster.service';

const layoutBreakpoint = 'lt-lg';
const USER_ADMIN_ROLES = ['ADMIN', 'SUPER_ADMIN'];

@Component({
  selector: 'kc-layout',
  templateUrl: './layout.component.html',
  styleUrls: ['./layout.component.scss']
})
export class LayoutComponent implements OnInit, AfterViewInit, OnDestroy {
  private _destroy$ = new Subject<void>();
  sidenavCollapsed$ = this.layoutService.sidenavCollapsed$;

  mobileQuery$ = this.mediaObserver.asObservable().pipe(map(() => this.mediaObserver.isActive(layoutBreakpoint)));

  toolbarShadowEnabled$ = this.router.events.pipe(
    filter(event => event instanceof NavigationEnd),
    startWith(null),
    map(() => checkRouterChildsData(this.router.routerState.root.snapshot, (data: any) => data.toolbarShadowEnabled))
  );

  scrollDisabled$ = this.router.events.pipe(
    filter(event => event instanceof NavigationEnd),
    startWith(null),
    map(() => checkRouterChildsData(this.router.routerState.root.snapshot, (data: any) => data.scrollDisabled))
  );

  containerEnabled$ = this.router.events.pipe(
    filter(event => event instanceof NavigationEnd),
    startWith(null),
    map(() => checkRouterChildsData(this.router.routerState.root.snapshot, (data: any) => data.containerEnabled))
  );

  userPermissions: string[] = [];

  @ViewChild('sidenav', { static: true }) sidenav!: MatSidenav;
  @ViewChild(MatSidenavContainer, { static: true })
  sidenavContainer!: MatSidenavContainer;

  constructor(
    private cd: ChangeDetectorRef,
    private mediaObserver: MediaObserver,
    private layoutService: LayoutService,
    private requesterService: RequesterService,
    private permissionSvc: PermissionService,
    private navigationService: NavigationService,
    private router: Router,
    private sdkConfigService: SdkConfigService,
    private selectedClusterService: SelectedClusterService,
    @Inject(DOCUMENT) private document: Document
  ) {
    // Check user is authentication
    if (this.requesterService.isAuthenticated()) {
      const userData = this.requesterService.get();
      if (!USER_ADMIN_ROLES.includes(userData?.userInfo?.user_type)) {
        this.permissionSvc.getPermissions().pipe(take(1)).subscribe();
      }
    }

    // Note: disabledFeature and sidenavList will keep on closure
    const sidenavList = this.sdkConfigService.initializeNavigation;

    // Sidenav Items
    combineLatest([
      this.requesterService.userData$.pipe(takeUntil(this._destroy$)),
      this.permissionSvc.userPermissions$.pipe(takeUntil(this._destroy$)),
      this.selectedClusterService.selectedClusterId$.pipe(takeUntil(this._destroy$), distinctUntilChanged())
    ])
      .pipe(takeUntil(this._destroy$))
      .subscribe(([userData, permissions, clusterId ]) => {
        if (userData === null) {
          if (this.userPermissions?.length) {
            this.userPermissions = [];
          }
          this.navigationService.loadItems([]);
          return;
        }
        // When user is authenticated
        const isAdmin = USER_ADMIN_ROLES.includes(userData?.userInfo?.user_type);
        if (isAdmin) {
          this.navigationService.loadItems(this.updateClusterId(sidenavList));
        } else {
          this.userPermissions = permissions;
          const newSidenavList = sidenavList.map(item => this._checkSidenavLinkPermission(item)).filter(Boolean);
          this.navigationService.loadItems(this.updateClusterId(newSidenavList));
        }
      });
  }

  ngOnInit() {
    this.mediaObserver
      .asObservable()
      .pipe(
        filter(() => this.mediaObserver.isActive(layoutBreakpoint)),
        takeUntil(this._destroy$)
      )
      .subscribe(() => this.layoutService.expandSidenav());

    this.layoutService.sidenavOpen$.pipe(takeUntil(this._destroy$)).subscribe(open => (open ? this.sidenav.open() : this.sidenav.close()));

    this.router.events
      .pipe(
        filter(event => event instanceof NavigationEnd),
        filter(() => this.mediaObserver.isActive(layoutBreakpoint)),
        takeUntil(this._destroy$)
      )
      .subscribe(() => this.sidenav.close());
  }

  ngAfterViewInit(): void {
    this.router.events
      .pipe(
        filter(e => e instanceof Scroll),
        takeUntil(this._destroy$)
      )
      .subscribe((e: any) => {
        if (e.position) {
          // backward navigation
          this.sidenavContainer.scrollable.scrollTo({
            start: e.position[0],
            top: e.position[1]
          });
        } else if (e.anchor) {
          // anchor navigation

          const scroll = (anchor: HTMLElement) =>
            this.sidenavContainer.scrollable.scrollTo({
              behavior: 'smooth',
              top: anchor.offsetTop,
              left: anchor.offsetLeft
            });

          let anchorElem = this.document.getElementById(e.anchor);

          if (anchorElem) {
            scroll(anchorElem);
          } else {
            setTimeout(() => {
              anchorElem = this.document.getElementById(e.anchor);
              scroll(anchorElem as HTMLElement);
            }, 100);
          }
        } else {
          // forward navigation
          this.sidenavContainer.scrollable.scrollTo({
            top: 0,
            start: 0
          });
        }
      });
  }

  ngOnDestroy(): void {
    this._destroy$.next();
    this._destroy$.complete();
  }

  /**
   * @description SidenavLink filter helper base on environment disables value
   * @return { SidenavLink | null } SidenavLink | null
   */
  private _sidenavLinkEnvFilter(item: SidenavLink, disabledFeatures: string[]): SidenavLink | null {
    if (!item.envName) {
      if (item.children?.length) {
        const children = item.children.map(child => this._sidenavLinkEnvFilter(child, disabledFeatures)).filter(Boolean);

        if (!children.length) return null;

        item['children'] = children as SidenavLink[];
      }
      return item;
    }

    if (typeof item.envName === 'string') return !disabledFeatures.includes(item.envName) ? item : null;
    return !item.envName.some(env => disabledFeatures.includes(env)) ? item : null;
  }

  /**
   * @description SidenavLink filter helper base on role base permission
   * @return { SidenavLink | null } SidenavLink | null
   */
  private _checkSidenavLinkPermission(item: SidenavLink): SidenavLink | null {
    // No Permission
    if (!item.permissionName) {
      if (item.children?.length) {
        const children = item.children.map(child => this._checkSidenavLinkPermission(child)).filter(Boolean);
        if (!children.length) return null;
        item['children'] = children as SidenavLink[];
      }
      return item;
    }

    if (typeof item.permissionName === 'string') {
      if (item.permissionName === 'ROLE_ADMIN') return null;
      return this.userPermissions.includes(item.permissionName) ? item : null;
    }
    return item.permissionName.some((_perm: string) => this.userPermissions.includes(_perm)) ? item : null;
  }

    updateClusterId(items: SidenavLink[]): SidenavLink[] {
      const replaceClusterId = (navItem: SidenavLink, replace: string = this.selectedClusterService.selectedClusterId) => {
        if (navItem.route ) {
          const routeParts = navItem.route.split('/');
          if (routeParts.length > 2 && routeParts[2] === 'k8s') {
            navItem.route = navItem.route.replace(routeParts[1], replace);
          }
        }``
        if (navItem.children) {
          navItem.children.forEach(child => replaceClusterId(child, replace));
        }
      };
      items.forEach(item => {
        replaceClusterId(item);
      });
      return items;
    }
}
