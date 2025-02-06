import icAccountCircle from '@iconify/icons-ic/twotone-account-circle';
import icLock from '@iconify/icons-ic/twotone-lock';
import { ToolbarUserMenuItem } from '@sdk-ui/interfaces';

export const USER_DROPDOWN_LIST: ToolbarUserMenuItem[] = [
  {
    id: '1',
    icon: icAccountCircle,
    label: 'My Profile',
    description: 'Personal Information',
    colorClass: 'text-teal-500',
    route: '/user-profile'
  },
  {
    id: '2',
    icon: icLock,
    label: 'Change Password',
    description: 'Change Your Password From Here',
    colorClass: 'text-primary-500',
    route: '/settings/reset-password'
  }
  // {
  //   id: '5',
  //   icon: icSettings,
  //   label: 'Settings',
  //   description: 'Manage other Settings',
  //   colorClass: 'text-purple-500',
  //   route: '/settings/general',
  //   envDisables: 'manage-account-setting',
  //   permissions: 'MANAGE_ACCOUNT_SETTINGS'
  // }
];
