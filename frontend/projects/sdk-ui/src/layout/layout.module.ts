import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LayoutComponent } from './layout.component';
import { RouterModule } from '@angular/router';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatButtonModule } from '@angular/material/button';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { SidenavModule } from './sidenav/sidenav.module';
import { ToolbarModule } from './toolbar/toolbar.module';
import { ProgressBarModule } from '@sdk-ui/ui';
import { ClusterService } from '@cluster/cluster.service';
import { K8sService } from '@k8s/k8s.service';

@NgModule({
  declarations: [LayoutComponent],
  imports: [
    CommonModule,
    RouterModule,
    MatSidenavModule,
    MatButtonModule,
    MatToolbarModule,
    MatIconModule,
    SidenavModule,
    ToolbarModule,
    ProgressBarModule
  ],
  providers: [ClusterService, K8sService],
})
export class LayoutModule {}
