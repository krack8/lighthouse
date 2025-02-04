import { Injectable } from '@angular/core';
import {
  Router,
  CanActivate,
  ActivatedRouteSnapshot,
  UrlTree,
  RouterStateSnapshot,
  CanActivateChild,
  CanLoad,
  Route,
  UrlSegment
} from '@angular/router';
import { Observable } from 'rxjs';
import { map, take } from 'rxjs/operators';
import { RequesterService, PermissionService } from '@core-ui/services';
import { USER_ADMIN_ROLES } from '@core-ui/models';

/**
 * @description Route Guard, Stable for CanActivate, CanActivateChild, CanLoad is Oon deployment
 * @implements {CanActivate, CanActivateChild, CanLoad}
 */
@Injectable({
  providedIn: 'root'
})
export class RoleGuardService implements CanActivate, CanActivateChild, CanLoad {
  constructor(
    private requesterService: RequesterService,
    private permissionService: PermissionService,
    private router: Router
  ) {}

  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
    return this._checkAuthorization(route, state);
  }

  canActivateChild(
    childRoute: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): boolean | UrlTree | Observable<boolean | UrlTree> | Promise<boolean | UrlTree> {
    return this._checkAuthorization(childRoute, state);
  }

  // !!! On Development Mode
  canLoad(route: Route, segments: UrlSegment[]): boolean | Observable<boolean> | Promise<boolean> {
    // console.log(route);
    // console.log(segments);
    return this._checkRolePermission(this.requesterService.get()?.userInfo?.user_type || '', route.data?.['permissions'] || []);
  }

  /**
   * @description checking authentication,
   */
  private _checkAuthorization(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
    const requester = this.requesterService.get();
    // Checking Authorization
    if (!this.requesterService.isAuthenticated()) {
      this.router.navigate(['/auth/login']);
      return false;
    }

    // TODO: optimize it when userProfile data fetch
    // Checking User Is Verification
    // Don't place it requester service.
    if (requester?.userInfo?.is_verified !== true) {
      return this.router.navigate(['/email-verification']);
    }

    // Check Permission
    return this._checkRolePermission(requester.userInfo?.user_type || '', route.data['permissions'] || [], state.url);
  }

  /**
   * @description checking role and permissions
   * @param userType: string,
   * @param routePermission: string[]
   * @param stateUrl: string
   * @returns boolean || Observable<boolean>
   */
  private _checkRolePermission(userType: string, routePermission: string[], stateUrl?: string): boolean | Observable<boolean> {
    if (USER_ADMIN_ROLES.some(r => r === userType)) return true;
    // Checking Non Admin Permission
    return this.permissionService.getPermissions().pipe(
      take(1),
      map(permissions => {
        if (routePermission.some((perm: string) => permissions.includes(perm))) return true;
        this.router.navigate(['/403'], {
          skipLocationChange: true,
          queryParams: {
            url: stateUrl
          }
        });
        return false;
      })
    );
  }
}
