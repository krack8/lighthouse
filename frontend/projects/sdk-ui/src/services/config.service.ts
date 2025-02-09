import { Injectable } from '@angular/core';
import { IConfig } from '@sdk-ui/interfaces';

let _config = {
  navigation: [],
  creationNavigation: [],
  toolbarUserDropdown: []
} as IConfig;

@Injectable({
  providedIn: 'root'
})
export class SdkConfigService {
  get initializeNavigation() {
    return _config.navigation;
  }

  get initializeCreationNavigation() {
    return _config.creationNavigation;
  }

  get initializeToolbarUserDropdown() {
    return _config.toolbarUserDropdown;
  }

  static injectConfig(__config: IConfig) {
    _config = __config;
  }

  static injectConfigValue(key: keyof IConfig, value: any) {
    _config[key] = value;
  }
}
