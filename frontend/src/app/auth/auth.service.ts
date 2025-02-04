import { Injectable } from '@angular/core';
import { HttpService } from '@core-ui/services/http.service';
import { RequesterService } from '@core-ui/services/requester.service';
import { Observable, of } from 'rxjs';
import { catchError, map, switchMap } from 'rxjs/operators';
import * as endpoints from './auth.endpoints';
import * as interfaces from './auth.interface';

@Injectable()
export class AuthService {
  constructor(
    private httpService: HttpService,
    private requester: RequesterService
  ) {}

  /**
   * @description Login with fetch user profile data
   */
  mcLogin(payload: interfaces.LoginPayload): Observable<any> {
    return this.httpService.post(endpoints.MC_LOGIN, payload).pipe(
      switchMap((res: any) => {
        const token = res.access_token;
        const refreshToken = res.refresh_token;
        this.requester.save({ token, refreshToken });

        return this.requester.getUserProfile().pipe(
          map(data => {
            const user = {
              userInfo: data,
              authorities: ['ROLE_' + data.userType],
              token,
              refreshToken
            };
            this.requester.save(user);
            return user;
          }),
          catchError(err => {
            const user = {
              userInfo: {
                first_name: 'Md Sajal',
                last_name: 'Mia',
                username: 'admin@default.com',
                created_at: '2024-12-22T09:55:07.796986426Z',
                updated_at: '2024-12-22T09:55:07.796986426Z',
                user_type: 'ADMIN',
                roles: [],
                clusterIdList: [],
                user_is_active: true,
                is_verified: true,
                forgot_password_token: null,
                phone: null
              },
              authorities: ['ROLE_ADMIN'],
              token,
              refreshToken
            };
            this.requester.save(user);
            return of(user);
          })
        );
      })
    );
  }
}
