import { Inject, Injectable, Optional } from '@angular/core';
import { BehaviorSubject, Observable, of } from 'rxjs';
import { filter } from 'rxjs/operators';

import { DeepPartial } from '@core-ui/interfaces';
import { mergeDeep } from '@core-ui/utils';
import { APP_ENV } from '@core-ui/constants';
import { IAppEnv } from '@core-ui/interfaces';

import { ICoreConfig, Theme } from './core-config.interfaces';
import { CoreConfig } from './core-config';
import { LOCAL_STORAGE_KEY } from '../requester.service';

@Injectable({ providedIn: 'root' })
export class CoreConfigService {
  private _generalInfoSubject = new BehaviorSubject<ICoreConfig | null>(null);
  readonly generalInfo$: Observable<ICoreConfig | null> = this._generalInfoSubject.asObservable().pipe(filter(config => !!config));

  constructor(@Optional() @Inject(APP_ENV) private _env: IAppEnv) {}

  // Core Config
  updateGeneralInfo(data: DeepPartial<ICoreConfig>): void {
    this._generalInfoSubject.next(mergeDeep(this._generalInfoSubject.value, data, true));
  }
  get generalInfoSnapshot(): ICoreConfig | null {
    return this._generalInfoSubject.value;
  }
  // Theming
  updateTheme(theme: Theme): void {
    this.updateGeneralInfo({ webTheme: theme });
  }

  loadConfigurationData(): Observable<any> {
    const _conf = new CoreConfig();
    const loc = localStorage.getItem(LOCAL_STORAGE_KEY);
    if (loc) {
      const currentUser = JSON.parse(loc as string);
      _conf.webTheme = currentUser?.userInfo?.webThemePreference;
    } else {
      _conf.webTheme = Theme.DARK;
    }
    this.updateGeneralInfo(_conf);
    return of(_conf);
  }
}
