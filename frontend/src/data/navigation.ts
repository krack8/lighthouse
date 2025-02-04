import icLayers from '@iconify/icons-ic/twotone-layers';
import icDashboard from '@iconify/icons-ic/twotone-dashboard';
import icAssigment from '@iconify/icons-ic/twotone-assignment';
import icBubbleChart from '@iconify/icons-ic/twotone-bubble-chart';
import icSettings from '@iconify/icons-ic/twotone-settings';
// import icShoppingBusket from '@iconify/icons-ic/twotone-shopping-basket';
import icFlare from '@iconify/icons-ic/twotone-flare';
import icDns from '@iconify/icons-ic/twotone-dns';
import icControlCamera from '@iconify/icons-ic/twotone-control-camera';
// import icMarketplace from '@iconify/icons-ic/twotone-store';
import icQueue from '@iconify/icons-ic/twotone-queue';
import icSupport from '@iconify/icons-ic/twotone-contact-support';
import icContacts from '@iconify/icons-ic/twotone-contacts';
import icCloudQueue from '@iconify/icons-ic/cloud-queue';
import icGroupWork from '@iconify/icons-ic/twotone-group-work';
import icPayment from '@iconify/icons-ic/twotone-payment';
// import accountTree from '@iconify/icons-ic/twotone-account-tree';
// import icRollBack from '@iconify/icons-ic/twotone-rotate-90-degrees-ccw';
// import icArrow from '@iconify/icons-ic/twotone-arrow-right';
// import icOpenInNew from '@iconify/icons-ic/twotone-open-in-new';
import icApps from '@iconify/icons-ic/twotone-apps';
import icEndpoint from '@iconify/icons-ic/twotone-explicit';
import { SidenavLink } from '@sdk-ui/interfaces';

export const SIDENAV_LIST: SidenavLink[] = [
  {
    type: 'link',
    label: 'Clusters',
    route: '/clusters',
    icon: icGroupWork,
    id: 'clusters',
    envName: 'cluster',
    permissionName: 'VIEW_CLUSTER'
  },
  // {
  //   type: 'link',
  //   label: 'Helm Apps',
  //   route: 'helm/apps',
  //   icon: icApps,
  //   id: 'helm-apps',
  //   envName: 'helm-apps',
  //   permissionName: ['VIEW_APP_CATALOG', 'VIEW_HELM_APPLICATION'],
  // },
  {
    type: 'dropdown',
    label: 'Management',
    icon: icContacts,
    id: 'management',
    children: [
      {
        type: 'link',
        label: 'Users',
        route: '/manage/users',
        permissionName: 'ROLE_ADMIN',
        id: 'users'
      },
      {
        type: 'link',
        label: 'Roles',
        route: '/manage/roles',
        permissionName: 'ROLE_ADMIN',
        id: 'roles'
      }
    ]
  }

  // {
  //   type: 'dropdown',
  //   label: 'Settings',
  //   icon: icSettings,
  //   id: 'settings',
  //   children: [
  //     {
  //       type: 'link',
  //       label: 'System',
  //       route: '/settings/system',
  //       envName: 'system-setting',
  //       permissionName: 'ROLE_ADMIN',
  //       id: 'system'
  //     },
  //   ],
  // },
];
