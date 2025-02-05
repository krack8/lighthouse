import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { RoleGuardService } from '@core-ui/guards';

const routes: Routes = [
  {
    path: 'system',
    loadChildren: () => import('./system-settings/system-settings.module').then(m => m.SystemSettingsModule),
    data: {
      featureName: 'system-setting'
    }
  }
  // {
  //   path: 'region',
  //   loadChildren: () => import('./region/region.module').then(m => m.RegionModule),
  //   data: {
  //     containerEnabled: true
  //   }
  // },
  // {
  //   path: 'environment-options',
  //   loadChildren: () => import('./environment-option/environment-option.module').then(m => m.EnvironmentOptionModule),
  //   data: {
  //     containerEnabled: true
  //   }
  // },
  // {
  //   path: 'kube-clusters',
  //   loadChildren: () => import('./kube-cluster/kube-cluster.module').then(m => m.KubeClusterModule),
  //   data: {
  //     containerEnabled: true
  //   }
  // },
  // {
  //   path: 'package',
  //   loadChildren: () => import('./package/package.module').then(m => m.PackageModule),
  //   data: {
  //     containerEnabled: true
  //   }
  // },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class SettingsRoutingModule {}
