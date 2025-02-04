import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ClusterListComponent } from './cluster-list/cluster-list.component';
import { ClusterDetailsComponent } from './cluster-details/cluster-details.component';
import { OverviewComponent } from './cluster-details/overview/overview.component';
import { SettingsComponent } from './cluster-details/settings/settings.component';
import { LogsComponent } from './cluster-details/logs/logs.component';
import { RoleGuardService } from '@core-ui/guards';
import { ClusterResolver } from './cluster.resolver';

const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    component: ClusterListComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster',
      toolbarTitle: 'Cluster List',
      permissions: ['VIEW_CLUSTER']
    }
  },
  {
    path: 'existing',
    pathMatch: 'full',
    loadComponent: () => import('./existing-cluster-form/existing-cluster-form.component').then(m => m.ExistingClusterFormComponent),
    data: {
      title: 'Existing Cluster',
      toolbarTitle: 'Existing Cluster',
      featureName: 'create_onboard_cluster', // guard info
      permissions: ['CREATE_CLUSTER']
    }
  },
  {
    path: ':clusterId',
    component: ClusterDetailsComponent,
    data: {
      title: 'Cluster | Details',
      toolbarTitle: 'Cluster Details',
      permissions: ['VIEW_CLUSTER']
    },
    resolve: { cluster: ClusterResolver },
    canActivate: [RoleGuardService],
    canActivateChild: [RoleGuardService],
    children: [
      {
        path: '',
        redirectTo: 'overview',
        pathMatch: 'full'
      },
      {
        path: 'logs',
        component: LogsComponent,
        data: {
          title: 'Cluster | Details | Logs',
          toolbarTitle: 'Logs',
          permissions: ['VIEW_CLUSTER']
        }
      },
      {
        path: 'overview',
        component: OverviewComponent,
        data: {
          title: 'Cluster | Details | Overview',
          toolbarTitle: 'Overview',
          permissions: ['VIEW_CLUSTER']
        }
      },
      {
        path: 'settings',
        component: SettingsComponent,
        data: {
          title: 'Cluster | Details | Settings',
          toolbarTitle: 'Settings',
          permissions: ['MANAGE_NODE_GROUP', 'DELETE_CLUSTER']
        }
      }
    ]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ClusterRoutingModule {}
