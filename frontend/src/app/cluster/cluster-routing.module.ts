import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RoleGuardService } from '@core-ui/guards';
import { ClusterListComponent } from './cluster-list/cluster-list.component';

const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    component: ClusterListComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Clusters',
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
