import { Component, inject, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ClusterService } from '@cluster/cluster.service';
import { MatExpansionModule } from '@angular/material/expansion';
import { CdkClipboardModule } from '@cdk-ui/clipboard';
import { ICluster } from '@cluster/cluster.model';
import { ActivatedRoute } from '@angular/router';
import { PageLayoutModule } from '@sdk-ui/ui';
import { ToolbarService } from '@sdk-ui/services';

@Component({
  selector: 'kc-cluster-init',
  standalone: true,
  imports: [CommonModule, MatExpansionModule, CdkClipboardModule, PageLayoutModule],
  templateUrl: './cluster-init.component.html',
  styleUrls: ['./cluster-init.component.scss']
})
export class ClusterInitComponent implements OnInit {
  private readonly _clusterService = inject(ClusterService);
  private readonly _route = inject(ActivatedRoute);
  private readonly _toolberService = inject(ToolbarService);

  clusterDetails!: ICluster;

  ngOnInit(): void {
    this.clusterDetails = this._route.snapshot.data['clusterDetails'];
    this._toolberService.changeData({ title: this.clusterDetails.name });
  }

  getClusterHelmChart(): void {}
}
