import { Injectable } from '@angular/core';
import { HttpService } from '@core-ui/services/http.service';
import { Observable } from 'rxjs';
import { Utils } from '@shared-ui/utils';
import * as endpoints from './access-role.endpoints';
import { IPermissionListObject, IRole } from './access-role-interface';
import { map } from 'rxjs/operators';

@Injectable()
export class AccessRoleService {
  constructor(private httpService: HttpService) {}

  getAccessRoles(): Observable<IRole[]> {
    return this.httpService.get(endpoints.ACCESS_ROLES).pipe(map(res => res?.roles));
  }

  getAccessRole(roleId: string): Observable<any> {
    return this.httpService.get(Utils.formatString(endpoints.ACCESS_ROLE, roleId)).pipe(map(res => res?.role));
  }

  createAccessRole(payload: any): Observable<any> {
    return this.httpService.post(endpoints.ACCESS_ROLES, payload);
  }

  updateAccessRole(roleId: string, payload: any): Observable<any> {
    return this.httpService.put(Utils.formatString(endpoints.ACCESS_ROLE, roleId), payload);
  }

  deleteAccessRole(roleId: string): Observable<any> {
    return this.httpService.delete(Utils.formatString(endpoints.ACCESS_ROLE, roleId));
  }

  getAccessPermissions(): Observable<IPermissionListObject> {
    return this.httpService.get(endpoints.ACCESS_PERMISSION).pipe(map(res => res?.permissions));
  }

  getUsersFromRole(roleId: string): Observable<any> {
    return this.httpService.get(Utils.formatString(endpoints.ACCESS_ROLE, roleId) + '/users');
  }
}
