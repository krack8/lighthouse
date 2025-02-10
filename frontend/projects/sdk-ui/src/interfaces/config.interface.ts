import { Icon } from '@visurel/iconify-angular';
import { CSSVariable } from './css-variable.enum';
import { ToolbarUserMenuItem } from './toolbar-menu-item.interface';

// Navigation
export type SidenavLink = {
  id?: string;
  type: 'link' | 'dropdown' | 'subheading';
  label: string;
  badge?: {
    value: string;
    background: CSSVariable;
    color: CSSVariable;
  };
  route?: string;
  routerLinkActive?: { exact: boolean };
  icon?: Icon;
  children?: SidenavLink[];
  // Permission
  envName?: string | string[];
  permissionName?: string | string[];
};

export type CreateDropDownLink = Pick<SidenavLink, 'label' | 'route' | 'icon' | 'envName' | 'permissionName'> & {
  queryParams?: Record<string, string | number>;
};

export interface IConfig {
  navigation: SidenavLink[];
  creationNavigation: CreateDropDownLink[];
  toolbarUserDropdown: ToolbarUserMenuItem[];
}
