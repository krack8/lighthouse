import { Injectable } from '@angular/core';
import { HttpRequest, HttpHandler, HttpEvent, HttpInterceptor } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { Router } from '@angular/router';
import { PermissionService, RequesterService } from '@core-ui/services';

@Injectable()
export class ErrorsInterceptor implements HttpInterceptor {
  constructor(
    private requesterSvc: RequesterService,
    private permissionSvc: PermissionService,
    private router: Router
  ) {}

  intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    return next.handle(request).pipe(catchError(res => this.errorHandler(res)));
  }

  private errorHandler(response: any): Observable<any> {
    console.log('errorHandler: ', response);
    const status: number = response.status;
    let message = 'Bad Request'; // Default Message
    switch (true) {
      case status === 401: {
        this.throwLogout();
        break;
      }
      case status === 403: {
        const currentUser = this.requesterSvc.get();
        if (currentUser?.userInfo?.user_type === 'USER') {
          if (response?.error?.path !== '/v1/permissions/users') {
            // TODO: Need to handle 403 error from api error
            // this.router.navigate(["/403"]);
            this.permissionSvc.fetchUserPermissions().subscribe();
            break;
          }
        }
        this.throwLogout();
        break;
      }
      case 500 <= status && status < 600: {
        message = 'Internal Server Error';
        break;
      }
    }

    let error = response.error;
    // eslint-disable-next-line no-prototype-builtins
    while (error?.hasOwnProperty('error')) {
      error = error.error;
    }
    if (error instanceof Blob) {
      message = 'Something is wrong!!!';
      error = null; // Stop Next Checking
    }
    if (typeof error === 'object' && error !== null) {
      const keys = Object.keys(error);
      if (keys.some(item => item === 'message')) {
        message = error['message'];
      }
    } else if (typeof error === 'string') {
      message = error;
    }
    return throwError({ error: { message }, status, message });
  }

  private throwLogout(): void {
    this.requesterSvc.clear();
    this.router.navigate(['/auth/login']);
  }
}
