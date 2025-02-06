import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import icClose from '@iconify/icons-ic/twotone-close';
import { CustomResourcesDefinationListComponent } from '../custom-resources-defination-list/custom-resources-defination-list.component';
import { K8sClusterCustomResourcesService } from '../k8s-cluster-custom-resources.service';
import { ToastrService } from '@sdk-ui/ui';

@Component({
  selector: 'kc-custom-resources-defination-details',
  templateUrl: './custom-resources-defination-details.component.html',
  styleUrls: ['./custom-resources-defination-details.component.scss']
})
export class CustomResourcesDefinationDetailsComponent implements OnInit {
  icClose = icClose;
  isLoading = false;
  details: any;

  constructor(
    @Inject(MAT_DIALOG_DATA) public data,
    public dialogRef: MatDialogRef<CustomResourcesDefinationListComponent>,
    private CustomResourcesService: K8sClusterCustomResourcesService,
    private toastr: ToastrService
  ) {}

  ngOnInit(): void {
    this.getDetails();
  }

  getDetails(): void {
    this.isLoading = true;
    this.CustomResourcesService.getCustomResourceDefinationDetails(this.data).subscribe({
      next: data => {
        if (data?.status === 'success') {
          this.details = data.data || [];
          this.isLoading = false;
        } else {
          this.isLoading = false;
        }
      },
      error: err => {
        this.isLoading = false;
        this.toastr.error(err, 'Failed');
      }
    });
  }

  isInt(value: string): boolean {
    const parsedValue = parseInt(value, 10);
    return !isNaN(parsedValue) && String(parsedValue) === value;
  }

  isObject(value: any): boolean {
    return typeof value === 'object' && value !== null;
  }
}
