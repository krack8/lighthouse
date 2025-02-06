import { Injectable } from '@angular/core';
import { AuthGuard } from './auth.guard';
import { ActivatedRouteSnapshot, CanLoad, Route, Router, RouterStateSnapshot, UrlSegment, UrlTree } from '@angular/router';
import { Observable } from 'rxjs';
import { RequesterService } from '@core-ui/services';

/**
 * @description
 * This is Admin guard. This guard implement with canActivate, canActivateChild and canLoad interface.
 * This guard is extends auth
 *
 * @status injected in appModule
 * @pubicApi
 */
@Injectable({
  providedIn: 'root'
})
export class AdminGuard extends AuthGuard implements CanLoad {
  constructor(
    protected override requesterService: RequesterService,
    protected override router: Router
  ) {
    super(requesterService, router);
  }

  // canActivate for current guard
  override canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): boolean | UrlTree | Observable<boolean | UrlTree> | Promise<boolean | UrlTree> {
    super.canActivate(route, state);
    if (this.checkAdmin()) {
      return true;
    }
    return this.router.createUrlTree(['/403'], {
      /* Removed unsupported properties by Angular migration: skipLocationChange. */ queryParams: {
        url: state.url
      }
    });
  }

  // CanActivate for child route guard
  override canActivateChild(
    childRoute: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): boolean | UrlTree | Observable<boolean | UrlTree> | Promise<boolean | UrlTree> {
    super.canActivateChild(childRoute, state);
    if (this.checkAdmin()) {
      return true;
    }
    return this.router.createUrlTree(['/403'], {
      /* Removed unsupported properties by Angular migration: skipLocationChange. */ queryParams: {
        url: state.url
      }
    });
  }

  // Can Load For feature module load
  canLoad(route: Route, segments: UrlSegment[]): boolean | Promise<boolean> | Observable<boolean> {
    if (this.checkAdmin()) {
      return true;
    }
    this.router.navigate(['**'], {
      skipLocationChange: true,
      queryParams: {
        url: route.path
      }
    });
    return false;
  }

  protected checkAdmin(): boolean {
    const requester = this.requesterService.get();
    return ['ADMIN', 'SUPER_ADMIN'].includes(requester?.userInfo?.user_type);
  }
}
