import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { RoleGuardService } from '@core-ui/guards';
import { PvDetailsComponent } from './pv-details/pv-details.component';
import { PvListComponent } from './pv-list/pv-list.component';

const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    component: PvListComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Persistent Volume List',
      toolbarTitle: 'Persistent-Volume',
      breadcrumb: 'persistent-volumes',
      permissions: ['VIEW_K8S_PERSISTENT_VOLUME']
    }
  },
  {
    path: 'persistent-volume-details',
    component: PvDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | PV | Details',
      toolbarTitle: 'Persitent Volume Details',
      breadcrumb: 'persistent-volumes',
      permissions: ['VIEW_K8S_PERSISTENT_VOLUME']
    }
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class K8sPersistentVolumeRoutingModule {}
