import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, CanActivateChild, Router, RouterStateSnapshot, UrlTree } from '@angular/router';
import { RequesterService } from '@core-ui/services';
import { Observable } from 'rxjs';

/**
 * @description
 * This is authentication guard. This guard implement with canActivate and canActivateChild interface.
 *
 * @status injected in appModule
 * @pubicApi
 */
@Injectable({
  providedIn: 'root'
})
export class AuthGuard implements CanActivate, CanActivateChild {
  constructor(
    protected requesterService: RequesterService,
    protected router: Router
  ) {}

  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): boolean | UrlTree | Observable<boolean | UrlTree> | Promise<boolean | UrlTree> {
    return this.checkAuthentication(state.url);
  }

  canActivateChild(
    childRoute: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): boolean | UrlTree | Observable<boolean | UrlTree> | Promise<boolean | UrlTree> {
    return this.checkAuthentication(state.url);
  }

  protected checkAuthentication(stateUrl?: string): boolean {
    const requester = this.requesterService.get();

    // Checking Authorization
    if (!this.requesterService.isAuthenticated()) {
      this.router.navigate(['/auth/login']);
      return false;
    }

    // Checking User Is Verification
    // ? Don't place it requester service.
    if (requester?.userInfo?.is_verified !== true) {
      this.router.navigate(['/email-verification']);
      return false;
    }

    return true;
  }
}
