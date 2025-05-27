import icGroupWork from '@iconify/icons-ic/twotone-group-work';
import icDeviceHub from '@iconify/icons-ic/twotone-device-hub';
import icFolderSpecial from '@iconify/icons-ic/twotone-folder-special';
import icApps from '@iconify/icons-ic/twotone-apps'; 
import icStorage from '@iconify/icons-ic/twotone-storage';
import icLan from '@iconify/icons-ic/twotone-lan';
import icSecurity from '@iconify/icons-ic/twotone-security';
import icSettings from '@iconify/icons-ic/twotone-settings';
import icExtension from '@iconify/icons-ic/twotone-extension';
import icContacts from '@iconify/icons-ic/twotone-contacts';

import { SidenavLink } from '@sdk-ui/interfaces';

export const SIDENAV_LIST: SidenavLink[] = [
  {
    type: 'link',
    label: 'Clusters',
    route: '/clusters',
    icon: icGroupWork,
    id: 'clusters'
  },
  {
    type: 'link',
    label: 'Nodes',
    route: `/:clusterId/k8s/node-list`,
    icon: icDeviceHub,
    id: 'nodes',
  },
  {
    type: 'link',
    label: 'Namespaces',
    route: `/:clusterId/k8s/namespaces`,
    icon: icFolderSpecial,
    id: 'namespaces',
  },
  {
    type: 'dropdown',
    label: 'Workloads',
    icon: icApps,
    id: 'Workloads',
    children: [
      {
        type: 'link',
        label: 'Pods',
        route: '/:clusterId/k8s/namespaces/pods',
        queryParams: { namespace: 'default' },
        permissionName: '',
        id: 'pods'
      },
      {
        type: 'link',
        label: 'Deployments',
        route: '/:clusterId/k8s/namespaces/deployments',
        queryParams: { namespace: 'default' },
        permissionName: '',
        id: 'deployments'
      },
      {
        type: 'link',
        label: 'Stateful Sets',
        route: '/:clusterId/k8s/namespaces/stateful-sets',
        queryParams: { namespace: 'default' },
        permissionName: '',
        id: 'statefulSets'
      },
      {
        type: 'link',
        label: 'Daemon Sets',
        route: '/:clusterId/k8s/namespaces/daemon-sets',
        queryParams: { namespace: 'default' },
        permissionName: '',
        id: 'daemonSets'
      },
      {
        type: 'link',
        label: 'Jobs',
        route: '/:clusterId/k8s/namespaces/job',
        queryParams: { namespace: 'default' },
        permissionName: '',
        id: 'jobs'
      },
      {
        type: 'link',
        label: 'Cron Jobs',
        route: '/:clusterId/k8s/namespaces/cron-job',
        queryParams: { namespace: 'default' },
        permissionName: '',
        id: 'cronJobs'
      },
    ]
  },
  {
    type: 'dropdown',
    label: 'Storage',
    icon: icStorage,
    id: 'Workloads',
    children: [
      {
        type: 'link',
        label: 'Persistent Volumes',
        route: '/:clusterId/k8s/persistent-volume',
        permissionName: '',
        id: 'pv'
      },
      {
        type: 'link',
        label: 'PVC',
        route: '/:clusterId/k8s/namespaces/pvcs',
        queryParams: { namespace: 'default' },
        permissionName: '',
        id: 'pvc'
      },
      {
        type: 'link',
        label: 'Storage Class',
        route: '/:clusterId/k8s/storage-class',
        permissionName: '',
        id: 'storageClass'
      }
    ]
  },
  {
    type: 'dropdown',
    label: 'Network',
    icon: icLan,
    id: 'Workloads',
    children: [
      {
        type: 'link',
        label: 'Services',
        route: '/:clusterId/k8s/namespaces/service',
        queryParams: { namespace: 'default' },
        permissionName: '',
        id: 'services'
      },
      {
        type: 'link',
        label: 'Ingresses',
        route: '/:clusterId/k8s/namespaces/ingresses',
        queryParams: { namespace: 'default' },
        permissionName: '',
        id: 'ingresses'
      },
      {
        type: 'link',
        label: 'Network Policies',
        route: '/:clusterId/k8s/namespaces/network-policy',
        queryParams: { namespace: 'default' },
        permissionName: '',
        id: 'networkPolicy'
      },
      {
        type: 'link',
        label: 'Endpoints',
        route: '/:clusterId/k8s/namespaces/endpoints',
        queryParams: { namespace: 'default' },
        permissionName: '',
        id: 'endpoints'
      },
      {
        type: 'link',
        label: 'Endpoint Slices',
        route: '/:clusterId/k8s/namespaces/endpoints-slice',
        queryParams: { namespace: 'default' },
        permissionName: '',
        id: 'endpointSlices'
      }
    ]
  },
  {
    type: 'dropdown',
    label: 'Security',
    icon: icSecurity,
    id: 'Workloads',
    children: [
      {
        type: 'link',
        label: 'Service Accounts',
        route: '/:clusterId/k8s/namespaces/service-accounts',
        queryParams: { namespace: 'default' },
        permissionName: '',
        id: 'serviceAccounts'
      },
      {
        type: 'link',
        label: 'Roles',
        route: '/:clusterId/k8s/namespaces/role',
        queryParams: { namespace: 'default' },
        permissionName: '',
        id: 'role'
      },
      {
        type: 'link',
        label: 'Role Bindings',
        route: '/:clusterId/k8s/namespaces/role-binding',
        queryParams: { namespace: 'default' },
        permissionName: '',
        id: 'roleBinding'
      },
      {
        type: 'link',
        label: 'Cluster Roles',
        route: '/:clusterId/k8s/cluster-role',
        permissionName: '',
        id: 'clusterRole'
      },
      {
        type: 'link',
        label: 'Cluster Role Bindings',
        route: '/:clusterId/k8s/cluster-role-binding',
        permissionName: '',
        id: 'clusterRoleBinding'
      }
    ]
  },
  {
    type: 'dropdown',
    label: 'Configurations',
    icon: icSettings,
    id: 'Workloads',
    children: [
      {
        type: 'link',
        label: 'Config Maps',
        route: '/:clusterId/k8s/namespaces/config-maps',
        queryParams: { namespace: 'default' },
        permissionName: '',
        id: 'configMap'
      },
      {
        type: 'link',
        label: 'Secrets',
        route: '/:clusterId/k8s/namespaces/secrets',
        queryParams: { namespace: 'default' },
        permissionName: '',
        id: 'secret'
      },
      {
        type: 'link',
        label: 'Resource Quotas',
        route: '/:clusterId/k8s/namespaces/resource-quota',
        queryParams: { namespace: 'default' },
        permissionName: '',
        id: 'resourceQuota'
      }
    ]
  },
  {
    type: 'link',
    label: 'Custom Resources',
    route: '/:clusterId/k8s/custom-resources-defination',
    icon: icExtension,
    id: 'custom-resources',
    permissionName: 'VIEW_K8S_CUSTOM_RESOURCE_DEFINATION'
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
        permissionName: 'VIEW_USER',
        id: 'users'
      },
      {
        type: 'link',
        label: 'Roles',
        route: '/manage/roles',
        permissionName: 'VIEW_ROLE',
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
