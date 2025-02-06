import { inject, Injectable } from '@angular/core';
import { HttpService } from '@core-ui/services';
import { Utils } from '@shared-ui/utils';
import { Observable } from 'rxjs';
import * as endpoints from './user-profile.endpoints';

@Injectable()
export class UserProfileService {
  private http = inject(HttpService);

  resetPassword(userId: string, payload: Record<string, string>): Observable<any> {
    return this.http.post(Utils.formatString(endpoints.RESET_PASSWORD, userId), payload);
  }
}
