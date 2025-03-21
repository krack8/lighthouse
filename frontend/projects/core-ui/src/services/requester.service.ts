import { Injectable } from '@angular/core';
import { Observable, BehaviorSubject, of } from 'rxjs';
import jwtDecode from 'jwt-decode';
import { map } from 'rxjs/operators';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Router } from '@angular/router';
import { USER_ADMIN_ROLES } from '@core-ui/models';
import { HttpService } from './http.service';

const MC_USER_INFO = '/v1/users/profile';
const MC_REFRESH_TOKEN = '/auth/refresh-token';
const USER_LOGOUT = '/auth/logout';

export const LOCAL_STORAGE_KEY = 'ngx-webstorage|kc-requester';

@Injectable({
  providedIn: 'root'
})
export class RequesterService {
  private userDataSubject = new BehaviorSubject<any>(null);
  userData$: Observable<any> = this.userDataSubject.asObservable();

  timeoutId: any;

  constructor(
    private httpService: HttpService,
    private snackBar: MatSnackBar,
    private router: Router
  ) {
    const loc = localStorage.getItem(LOCAL_STORAGE_KEY);
    if (loc) {
      const currentUser = JSON.parse(loc as string);
      if (currentUser) {
        // ? Checking User Status
        if (currentUser?.userInfo?.user_is_active === false) {
          if (currentUser?.userInfo?.user_type === 'USER') {
            this.snackBar.open('Your account is deactivated. Please contact your admin', 'Close', {
              duration: 10000
            });
          } else {
            this.snackBar.open('Your account is deactivated', 'Close', {
              duration: 10000
            });
          }
          this.clear();
          this.router.navigate(['/auth/login']);
          return;
        }

        this.tokenExpireSetTimeout(currentUser);
        this.userDataSubject.next(currentUser);
      }
    }
  }

  get() {
    return this.userDataSubject.value;
  }

  save(user: any): void {
    localStorage.setItem(LOCAL_STORAGE_KEY, JSON.stringify(user));
    this.loadUserData(user);
    if (user?.refreshToken) {
      this.tokenExpireSetTimeout(user);
    } else {
      if (this.timeoutId) {
        clearTimeout(this.timeoutId);
      }
    }
  }

  clear() {
    localStorage.removeItem(LOCAL_STORAGE_KEY);
    this.loadUserData(null);
    if (this.timeoutId) {
      clearTimeout(this.timeoutId);
    }
  }

  /**
   * @param {any} data - User authentication data
   */
  loadUserData(data: any): void {
    this.userDataSubject.next(data);
  }

  isAuthTokenValid(accessToken: string): boolean {
    return this._getTokenExpireTime(accessToken) * 1000 > Date.now();
  }
  /**
   * @param {string} token - JWT ACCESS Token
   * @returns {any} {username: string, userType: string}
   */
  getUserDataFromToken(token: string): any {
    const decoded: any = jwtDecode(token);
    const data = {
      username: decoded.username
      // userType: decoded.authorities
    };
    return data;
  }

  /**
   * @param {string} token - JWT ACCESS Token
   * @return {number} expire time as millisecond
   */
  _getTokenExpireTime(token: string): number {
    try {
      const decoded: any = jwtDecode(token);
      // default decoded exp format is second
      return decoded.exp;
    } catch (err) {
      console.log('err', err);
    }
    return 0;
  }

  tokenExpireSetTimeout(user: any): void {
    if (this.timeoutId) {
      clearTimeout(this.timeoutId);
    }
    const token = user.token;
    const refreshToken = user.refreshToken;
    if (refreshToken) {
      const expireTime = this._getTokenExpireTime(token);
      // console.log(new Date(expireTime * 1000))
      const duration = expireTime * 1000 - new Date().getTime();
      this.timeoutId = setTimeout(() => {
        this.getNewByRefreshToken().subscribe();
      }, duration);
    }
  }

  get isAdmin(): boolean {
    const userType = this.get()?.userInfo?.user_type;
    return USER_ADMIN_ROLES.some(r => r === userType);
  }

  public isAuthenticated(): boolean {
    return !!this.userDataSubject.value?.token;
  }

  // APIs
  getNewByRefreshToken(): Observable<any> {
    const currentData = this.get();
    if (!currentData?.refreshToken) {
      if (this.timeoutId) {
        clearTimeout(this.timeoutId);
      }
      return of(false);
    }
    const refresh_token = currentData.refreshToken;
    const expireTime = this._getTokenExpireTime(refresh_token);
    if (new Date().getTime() > expireTime * 1000) {
      this.snackBar.open('Your session has been expired! Please Sign In Again.', 'close', {
        duration: 5000,
        panelClass: ['snackbar-dark']
      });
      this.clear();
      this.router.navigate(['/auth/login']);
      return of(false);
    }
    return this.httpService.post(MC_REFRESH_TOKEN, { refresh_token }).pipe(
      map((res: any) => {
        currentData['token'] = res.access_token;
        if (res.refresh_token) currentData['refreshToken'] = res.refresh_token;
        this.save(currentData);
        return res;
      })
    );
  }

  getUserProfile(): Observable<any> {
    return this.httpService.get(MC_USER_INFO);
  }

  logoutUser(): Observable<any> {
    return this.httpService.post(USER_LOGOUT, {});
  }
}
