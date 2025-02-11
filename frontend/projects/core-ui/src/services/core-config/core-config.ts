import { ICoreConfig, ISsoConfig, Theme } from './core-config.interfaces';

export class CoreConfig implements ICoreConfig {
  id: string = '';
  updateDate!: string;
  updatedBy!: string;
  status!: string;
  // Theme
  webTheme!: Theme;
  // SSO
  ssoEnabled: boolean = false;
  sso: ISsoConfig | null = null;
  // System Setting
  // - Logo
  logoUrl: string = '';
  favicon: string = '';
  name: string = '';
  // - Other Config
  passwordLength: number = 8;
  // domains: string[] = [];
}
