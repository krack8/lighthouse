import { CommonModule } from '@angular/common';
import { Component, Input, OnInit } from '@angular/core';
import { AgoPipe } from '../../../../../../projects/shared-ui/src/pipes/ago.pipe';
import { K8sNamespacesService } from '@k8s/k8s-namespaces/k8s-namespaces.service';
import { ToastrService } from '@sdk-ui/ui';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatIconModule } from '@angular/material/icon';
import icRefresh from '@iconify/icons-ic/twotone-refresh';
import { IconModule } from '@visurel/iconify-angular';

@Component({
  selector: 'kc-k8s-events-template',
  templateUrl: './k8s-events-template.component.html',
  styleUrls: ['./k8s-events-template.component.scss'],
  standalone: true,
  imports: [CommonModule, AgoPipe, MatProgressSpinnerModule, MatIconModule, IconModule]
})
export class K8sEventsTemplateComponent implements OnInit {
  @Input() objectName: any;
  eventsData: any;
  icRefresh = icRefresh;
  isLoading: boolean = true;
  constructor(
    private _namespaceService: K8sNamespacesService,
    private toastr: ToastrService
  ) {}

  ngOnInit(): void {
    this.getEvents();
  }

  getEvents(): void {
    this.isLoading = true;
    this._namespaceService.getEvents(this.objectName).subscribe({
      next: res => {
        this.eventsData = res?.data;
        this.isLoading = false;
      },
      error: err => {
        this.toastr.error('Failed: ', err.error.message);
      }
    });
  }
}
