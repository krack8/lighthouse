import { Route, Routes } from '@angular/router';
import { LayoutComponent } from '@sdk-ui/layout';
import { AdminGuard } from '@core-ui/guards';

const childrenRoutes: Route[] = [
  {
    path: '',
    redirectTo: 'clusters',
    pathMatch: 'full'
  },
  {
    path: 'clusters',
    loadChildren: () => import('./cluster/cluster.module').then(m => m.ClusterModule),
    data: {
      containerEnabled: true,
      featureName: 'cluster'
    }
  },
  {
    path: ':clusterId/k8s',
    loadChildren: () => import('./k8s/k8s.module').then(m => m.K8sModule),
    data: {
      featureName: 'k8s'
    }
  },
  // {
  //   path: 'helm/apps',
  //   loadChildren: () => import('./helm-apps/helm-apps.module').then( m => m.HelmAppsModule),
  //   data: {
  //     featureName: "helm-apps",
  //     containerEnabled: true,

  //   }
  // },
  {
    path: 'manage',
    children: [
      {
        path: 'users',
        loadChildren: () => import('./management/user/user.module').then(m => m.UserModule),
        data: {
          containerEnabled: true
        }
      },
      {
        path: 'roles',
        loadChildren: () => import('./management/access-role/access-role.module').then(m => m.AccessRoleModule),
        data: {
          containerEnabled: true
        }
      }
    ]
  },
  {
    path: 'settings',
    loadChildren: () => import('./settings/settings.module').then(m => m.SettingsModule),
    data: {
      containerEnabled: true
    }
  },
  {
    path: 'user-profile',
    loadChildren: () => import('./user-profile/user-profile.module').then(m => m.UserProfileModule),
    data: {
      containerEnabled: true
    }
  },
  {
    path: '403',
    loadChildren: () => import('./errors/error-403/error-403.module').then(m => m.Error403Module)
  },
  {
    path: '**',
    loadChildren: () => import('./errors/error-404/error-404.module').then(m => m.Error404Module)
  }
];

export const appRoutes: Routes = [
  {
    path: 'auth',
    loadChildren: () => import('./auth/auth.module').then(m => m.AuthModule)
  },
  {
    path: 'k8s/terminal',
    loadChildren: () => import('./k8s/k8s-terminal/k8s-terminal.module').then((m) => m.K8sTerminalModule),
  },
  {
    path: '',
    component: LayoutComponent,
    children: childrenRoutes
  }
];
