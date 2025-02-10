export enum Theme {
  DARK = 'DARK',
  LIGHT = 'LIGHT',
  LIGHT_PINK = 'LIGHT_PINK'
}

export interface ISsoConfig {
  type: string;
  authorizeURI: string;
  clientID: string;
  clientSecret: string;
  publicKey: string;
}

export interface ICoreConfig {
  id: string;
  updateDate: string;
  updatedBy: string;
  status: string;
  // SSO
  ssoEnabled: boolean;
  sso: ISsoConfig | null;
  //? System Setting
  // - Theme
  webTheme: Theme;
  // - Logo
  logoUrl: string | null;
  favicon: string | null;
  name: string | null;
  // - Other Config
  passwordLength: number;
  domains?: string[] | null;
}
