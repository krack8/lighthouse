import { Component, OnInit } from '@angular/core';
import icAddCircleOutline from '@iconify/icons-ic/twotone-add';
import { ToastrService } from '@sdk-ui/ui';
import { ClusterService } from '../cluster.service';
import { ToolbarService } from '@sdk-ui/services/toolbar.service';
import { ICluster } from '@cluster/cluster.model';
import { Router } from '@angular/router';
import { MatDialog } from '@angular/material/dialog';

enum View {
  GRID = 'grid',
  LIST = 'list'
}

@Component({
  selector: 'kc-cluster-list',
  templateUrl: './cluster-list.component.html',
  styleUrls: ['./cluster-list.component.scss']
})
export class ClusterListComponent implements OnInit {
  icAddCircleOutline = icAddCircleOutline;
  clusterList: any = [];
  dataLoading!: boolean;
  serverError: boolean = false;

  viewStyle: View = View.LIST;

  constructor(
    private clusterService: ClusterService,
    private toastrService: ToastrService,
    private toolbarService: ToolbarService,
    private router: Router,
    private _dialog: MatDialog
  ) {}

  ngOnInit(): void {
    this.toolbarService.changeData({ title: 'Cluster' });
    this.getCluster();
  }

  getCluster(): void {
    this.dataLoading = true;
    this.clusterService.getClusters().subscribe({
      next: data => {
        this.clusterList = data || [];
        this.dataLoading = false;
        this.serverError = false;
      },
      error: err => {
        this.dataLoading = false;
        this.serverError = true;
        this.toastrService.error(err.message, 'ERROR');
      }
    });
  }

  changeView(): void {
    if (this.viewStyle === View.LIST) {
      this.viewStyle = View.GRID;
    } else {
      this.viewStyle = View.LIST;
    }
  }

  routeToDetails(cluster?: ICluster): void {
    if (cluster && cluster.is_active) {
      this.router.navigate(['/clusters', cluster?.id, 'k8s']);
      return;
    }

    import('../cluster-form/cluster-form.component').then(m => {
    const dialog = this._dialog.open(m.ClusterFormComponent, {
        width: '800px',
        data: cluster
      });
      dialog.afterClosed().subscribe((result) => {
        if (result) this.getCluster();
      })
    });
  }
}
