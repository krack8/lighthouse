import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NgTerminalModule } from 'ng-terminal';
import { K8sTerminalRoutingModule } from './k8s-terminal-routing.module';
import { PodTerminalComponent } from './pod-terminal/pod-terminal.component';

@NgModule({
  declarations: [PodTerminalComponent],
  imports: [CommonModule, K8sTerminalRoutingModule, NgTerminalModule]
})
export class K8sTerminalModule {}
