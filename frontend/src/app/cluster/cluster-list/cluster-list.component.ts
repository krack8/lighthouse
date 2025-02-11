import { Component, OnInit } from '@angular/core';
import icAddCircleOutline from '@iconify/icons-ic/twotone-add';
import { ToastrService } from '@sdk-ui/ui';
import { ClusterService } from '../cluster.service';
import { ToolbarService } from '@sdk-ui/services/toolbar.service';

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
    private toolbarService: ToolbarService
  ) {}

  ngOnInit(): void {
    this.toolbarService.changeData({ title: 'Cluster' });
    this.getCluster();
  }

  getCluster(): void {
    this.dataLoading = true;
    this.clusterService.mcGetClusters().subscribe({
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
}
