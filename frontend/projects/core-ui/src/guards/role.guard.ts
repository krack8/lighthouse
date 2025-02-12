import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, Router, RouterStateSnapshot, UrlTree } from '@angular/router';
import { map, take } from 'rxjs/operators';
import { Observable } from 'rxjs';
import { AdminGuard } from './admin.guard';
import { PermissionService } from '@core-ui/services';
import { RequesterService } from '@core-ui/services';

/**
 * @description
 * This is Admin guard. This guard implement with canActivate, canActivateChild and canLoad interface.
 * This guard is extends auth
 *
 * @status It's not used and not injected yet
 * @development
 */
@Injectable({
  providedIn: 'root'
})
export class RoleGuard extends AdminGuard {
  constructor(
    protected override requesterService: RequesterService,
    protected override router: Router,
    private permissionService: PermissionService
  ) {
    super(requesterService, router);
  }

  override canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): boolean | UrlTree | Observable<boolean | UrlTree> | Promise<boolean | UrlTree> {
    super.canActivate(route, state);
    return this.checkRolePermission(route, state);
  }

  override canActivateChild(
    childRoute: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): boolean | UrlTree | Observable<boolean | UrlTree> | Promise<boolean | UrlTree> {
    super.canActivateChild(childRoute, state);
    return this.checkRolePermission(childRoute, state);
  }

  checkRolePermission(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): boolean | Observable<boolean | UrlTree> {
    const validPermissions: string[] = route.data['permissions'] || [];
    return this.permissionService.getPermissions().pipe(
      take(1),
      map((permissions: string[]) => {
        if (validPermissions.some((perm: string) => permissions.includes(perm))) {
          return true;
        }
        return this.router.createUrlTree(['/403'], {
          /* Removed unsupported properties by Angular migration: skipLocationChange. */ queryParams: {
            url: state.url
          }
        });
      })
    );
  }
}
