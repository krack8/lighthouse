import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable, of } from 'rxjs';
import { map, switchMap } from 'rxjs/operators';
import { RequesterService } from './requester.service';
import { HttpService } from './http.service';

@Injectable({
  providedIn: 'root'
})
export class PermissionService {
  private userPermissions = new BehaviorSubject<string[]>([]);
  userPermissions$: Observable<string[]> = this.userPermissions.asObservable();

  constructor(
    private _httpService: HttpService,
    private _requesterService: RequesterService
  ) {}

  getPermissions(force: boolean = false): Observable<string[]> {
    return this.userPermissions$.pipe(
      switchMap((permissions: string[]) => {
        if (!permissions.length || force) {
          return this.fetchUserPermissions();
        }
        return of(permissions);
      })
    );
  }

  get userPermissionsSnapshot(): string[] {
    return this.userPermissions.value;
  }

  loadUserPermissions(permissions: string[]): void {
    this.userPermissions.next(permissions);
  }

  /**
   * @description checking authorites for admin and non-admin permission
   * @param {string | string[]} permission - permission name
   * @return {boolean} boolean
   */
  hasAuthorities(permission: string[] | string): boolean {
    if (this._requesterService.isAdmin) return true;
    // Check permission
    const userPermissions = this.userPermissionsSnapshot;
    if (typeof permission === 'string') return userPermissions.includes(permission);
    return permission.some((perm: string) => userPermissions.includes(perm));
  }

  // Dep
  fetchUserPermissions(): Observable<string[]> {
    return this._httpService.get('/v1/permissions/users').pipe(
      map(res => {
        const _permissions: string[] = ['*'];
        Object.entries(res).map(([_, value]) => {
          if (value instanceof Array && value.length) {
            value.forEach(item => {
              _permissions.push(item.name);
            });
          }
        });
        this.loadUserPermissions(_permissions);
        return _permissions;
      })
    );
  }
}
