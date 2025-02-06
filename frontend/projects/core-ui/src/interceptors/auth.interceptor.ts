import { Inject, Injectable, Optional } from '@angular/core';
import { HttpRequest, HttpHandler, HttpEvent, HttpInterceptor } from '@angular/common/http';
import { Observable } from 'rxjs';
import { filter, switchMap, take } from 'rxjs/operators';
import { APP_ENV } from '@core-ui/constants';
import { IAppEnv } from '@core-ui/interfaces';
import { RequesterService } from '@core-ui/services';

const MC_REFRESH_TOKEN = '/v1/auth/refresh-token';

@Injectable()
export class AuthInterceptor implements HttpInterceptor {
  isRefreshing: boolean = false; // Prevent for multiple refresh token request

  constructor(
    private requester: RequesterService,
    @Optional() @Inject(APP_ENV) private _env: IAppEnv
  ) {}

  intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    const data = this.requester.get();
    if (data?.token) {
      // INFO: APIs to ignore token request
      if (request.url.includes(MC_REFRESH_TOKEN)) {
        return next.handle(request);
      }

      // INFO: Authentication APIs request
      const accessToken = data.token;
      if (this.requester.isAuthTokenValid(accessToken)) {
        return next.handle(this.getAuthorizeRequest(request, accessToken));
      }
      if (!this.isRefreshing) {
        // Request For Refresh Token
        this.isRefreshing = true;
        return this.requester.getNewByRefreshToken().pipe(
          take(1),
          switchMap((res: any) => {
            this.isRefreshing = false;
            return next.handle(this.getAuthorizeRequest(request, res['data']?.accessToken));
          })
        );
      }

      // Pending For New Token
      return this.requester.userData$.pipe(
        filter(data => data?.token && typeof data?.token === 'string'),
        take(1),
        switchMap(data => {
          return next.handle(this.getAuthorizeRequest(request, data?.token));
        })
      );
    }
    return next.handle(request);
  }

  /**
   * @definition Middle man to pass token on request
   * @param request HttpRequest<any>
   * @param accessToken accessToken<string>
   * @returns HttpRequest<any>
   */
  private getAuthorizeRequest(request: HttpRequest<any>, accessToken: string): HttpRequest<any> {
    return request.clone({
      headers: request.headers.append('Authorization', `Bearer ${accessToken}`)
    });
  }
}
