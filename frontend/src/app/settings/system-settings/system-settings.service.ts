import { Injectable } from '@angular/core';
import * as endpoints from './system-settings.endpoints';
import { Observable } from 'rxjs';
import { HttpService } from '@core-ui/services';

@Injectable()
export class SystemSettingsService {
  constructor(private _httpServie: HttpService) {}

  saveSystemSetting(payload: any): Observable<any> {
    return this._httpServie.put(endpoints.SAVE_SYSTEM_SETTINGS, payload);
  }
  saveWebThemePreference(qp: { webTheme: string }): Observable<any> {
    return this._httpServie.put(endpoints.SAVE_WEB_THEME_PREFERENCE, {}, qp);
  }
}
