import { Icon } from '@visurel/iconify-angular';
import { CSSVariable } from './css-variable.enum';

export type NavigationItem = any | any | any;

export interface NavigationLink {
  type: 'link';
  route: string | any;
  queryParams?: Record<string, string | number>;
  fragment?: string;
  label: string;
  icon?: Icon;
  id: string;
  isVisible?: boolean;
  routerLinkActive?: { exact: boolean };
  badge?: {
    value: string;
    background: CSSVariable;
    color: CSSVariable;
  };
}

export interface NavigationDropdown {
  type: 'dropdown';
  label: string;
  icon?: Icon;
  id: string;
  isVisible?: boolean;
  children?: Array<NavigationLink | NavigationDropdown>;
  badge?: {
    value: string;
    background: CSSVariable;
    color: CSSVariable;
  };
}

export interface NavigationSubheading {
  type: 'subheading';
  label: string;
  children: Array<NavigationLink | NavigationDropdown>;
}
