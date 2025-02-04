import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import * as endpoints from './user.endpoints';
import { HttpService } from '@core-ui/services/http.service';
import { Utils } from '@shared-ui/utils';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  constructor(private _httpService: HttpService) {}

  // MultiCluster
  mcGetUsers(): Observable<any> {
    return this._httpService.get(endpoints.MC_GET_USERS);
  }

  mcCreateUser(payload: any): Observable<any> {
    return this._httpService.post(endpoints.MC_GET_USERS, payload);
  }

  mcGetUser(userId: string): Observable<any> {
    return this._httpService.get(Utils.formatString(endpoints.MC_USER, userId));
  }

  mcUpdateUser(userId: string, payload: any): Observable<any> {
    return this._httpService.put(Utils.formatString(endpoints.MC_USER, userId), payload);
  }

  mcDeleteUser(userId: string): Observable<any> {
    return this._httpService.delete(Utils.formatString(endpoints.MC_USER, userId));
  }

  mcAssignRoles(payload: Record<string, any>): Observable<any> {
    return this._httpService.post(endpoints.ASSIGN_ROLE, payload);
  }
}
