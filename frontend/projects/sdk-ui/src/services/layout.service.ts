import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class LayoutService {
  private _quickpanelOpenSubject = new BehaviorSubject<boolean>(false);
  quickpanelOpen$ = this._quickpanelOpenSubject.asObservable();

  private _sidenavOpenSubject = new BehaviorSubject<boolean>(false);
  sidenavOpen$ = this._sidenavOpenSubject.asObservable();

  private _sidenavCollapsedSubject = new BehaviorSubject<boolean>(false);
  sidenavCollapsed$ = this._sidenavCollapsedSubject.asObservable();

  private _sidenavCollapsedOpenSubject = new BehaviorSubject<boolean>(false);
  sidenavCollapsedOpen$ = this._sidenavCollapsedOpenSubject.asObservable();

  private _configpanelOpenSubject = new BehaviorSubject<boolean>(false);
  configpanelOpen$ = this._configpanelOpenSubject.asObservable();

  private sideNavExpanded!: boolean;

  constructor(private router: Router) {}

  openQuickpanel() {
    this._quickpanelOpenSubject.next(true);
  }

  closeQuickpanel() {
    this._quickpanelOpenSubject.next(false);
  }

  openSidenav() {
    this._sidenavOpenSubject.next(true);
  }

  closeSidenav() {
    this._sidenavOpenSubject.next(false);
  }

  collapseSidenav() {
    this.sideNavExpanded = false;
    this._sidenavCollapsedSubject.next(true);
  }

  expandSidenav() {
    this.sideNavExpanded = true;
    this._sidenavCollapsedSubject.next(false);
  }

  collapseOpenSidenav() {
    this._sidenavCollapsedOpenSubject.next(true);
  }

  collapseCloseSidenav() {
    this._sidenavCollapsedOpenSubject.next(false);
  }

  openConfigpanel() {
    this._configpanelOpenSubject.next(true);
  }

  closeConfigpanel() {
    this._configpanelOpenSubject.next(false);
  }

  isSideNavExpanded() {
    return this.sideNavExpanded;
  }

  enableRTL() {
    this.router
      .navigate([], {
        queryParams: {
          rtl: 'true'
        }
      })
      .then(() => {
        if (window) {
          window.location.reload();
        }
      });
  }

  disableRTL() {
    this.router
      .navigate([], {
        queryParams: {
          rtl: 'false'
        }
      })
      .then(() => {
        if (window) {
          window.location.reload();
        }
      });
  }
}
