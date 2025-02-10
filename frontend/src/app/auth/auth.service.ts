import { Injectable } from '@angular/core';
import { HttpService } from '@core-ui/services/http.service';
import { RequesterService } from '@core-ui/services/requester.service';
import { Observable } from 'rxjs';
import { map, switchMap } from 'rxjs/operators';
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
              authorities: ['ROLE_' + data.user_type],
              token,
              refreshToken
            };
            console.log(JSON.stringify(data, null, '\t'));
            this.requester.save(user);
            return user;
          })
        );
      })
    );
  }
}
