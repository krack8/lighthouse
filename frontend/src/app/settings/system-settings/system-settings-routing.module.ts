import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { SystemSettingsComponent } from './containers/system-settings/system-settings.component';
import { AdminGuard } from '@core-ui/guards';

const routes: Routes = [
  {
    path: '',
    component: SystemSettingsComponent,
    data: {
      title: 'Settings | System'
    },
    canActivate: [AdminGuard]
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class SystemSettingsRoutingModule {}
