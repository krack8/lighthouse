import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {NgTerminalModule} from 'ng-terminal';
import { K8sTerminalRoutingModule } from './k8s-terminal-routing.module';
import { K8sTerminalComponent } from './k8s-terminal.component';


@NgModule({
  declarations: [
    K8sTerminalComponent
  ],
  imports: [
    CommonModule,
    K8sTerminalRoutingModule,
    NgTerminalModule
  ]
})
export class K8sTerminalModule { }
