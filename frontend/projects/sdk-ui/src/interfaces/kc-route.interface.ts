import { Route } from '@angular/router';

export interface KcRouteData {
  scrollDisabled?: boolean;
  toolbarShadowEnabled?: boolean;
  containerEnabled?: boolean;

  [key: string]: any;
}

export interface KcRoute extends Route {
  data?: KcRouteData;
  children?: KcRoute[];
}

export type KcRoutes = KcRoute[];
