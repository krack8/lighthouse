import { Icon } from '@visurel/iconify-angular';

export interface ToolbarUserMenuItem {
  id: string;
  icon: Icon;
  label: string;
  description: string;
  colorClass: string;
  route: string;
  envDisables?: string;
  permissions?: string;
}
