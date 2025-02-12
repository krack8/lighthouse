import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ClusterListComponent } from './cluster-list/cluster-list.component';
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
      permissions: ['*']
    }
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ClusterRoutingModule {}
