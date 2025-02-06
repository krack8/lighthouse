import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import icClose from '@iconify/icons-ic/twotone-close';
import { K8sNamespacesListComponent } from '../k8s-namespaces-list/k8s-namespaces-list.component';
import { K8sNamespacesService } from '../k8s-namespaces.service';

@Component({
  selector: 'kc-k8s-namespaces-details',
  templateUrl: './k8s-namespaces-details.component.html',
  styleUrls: ['./k8s-namespaces-details.component.scss']
})
export class K8sNamespacesDetailsComponent implements OnInit {
  icClose = icClose;
  isLoading = false;
  details: any;

  constructor(
    @Inject(MAT_DIALOG_DATA) public data,
    public dialogRef: MatDialogRef<K8sNamespacesListComponent>,
    private namespaceService: K8sNamespacesService
  ) {}

  ngOnInit(): void {
    this.getDetails();
  }

  getDetails(): void {
    this.isLoading = true;
    this.namespaceService.getNamespacesDetailsV1(this.data).subscribe({
      next: data => {
        this.details = data.data || [];
        this.isLoading = false;
      },
      error: err => {
        console.log(err);
      }
    });
  }
}
