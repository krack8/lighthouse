import { Component, HostBinding, Inject, Input, OnInit } from '@angular/core';
import { Platform } from '@angular/cdk/platform';
import { Observable } from 'rxjs';
import { stagger80ms, fadeInUp400ms, scaleIn400ms, fadeInRight400ms } from '@sdk-ui/animations';
import { CoreConfigService } from '@core-ui/services';
import { LayoutService, ToolbarService } from '@sdk-ui/services';
import { ClusterService } from '@cluster/cluster.service';
import { ActivatedRoute, Router } from '@angular/router';
import { K8sService } from '@k8s/k8s.service';
import { SelectedClusterService } from '@core-ui/services/selected-cluster.service';

@Component({
  selector: 'kc-toolbar',
  templateUrl: './toolbar.component.html',
  styleUrls: ['./toolbar.component.scss'],
  animations: [stagger80ms, fadeInUp400ms, scaleIn400ms, fadeInRight400ms]
})
export class ToolbarComponent implements OnInit {
  @Input() mobileQuery!: boolean;

  @Input()
  @HostBinding('class.shadow-b')
  hasShadow!: boolean;

  isChrome!: boolean;

  coreConfig$ = this.coreConfigService.generalInfo$;
  toolbarData$: Observable<any> = this.toolbarService.currentData;

  clusterChangeToggleVisibilty: boolean = false;
  clusterList = [];
  selectedCluster: string;

  constructor(
    private layoutService: LayoutService,
    private toolbarService: ToolbarService,
    private coreConfigService: CoreConfigService,
    private platform: Platform,
    private clusterService: ClusterService,
    private router: Router,
    private k8sService: K8sService,
  ) {}

  ngOnInit() {
   this.clusterService.getAllClusterList();
   this.k8sService.clusterId$.subscribe((clusterId: string) => {
      if (clusterId) {
        this.selectedCluster = clusterId;
        this.clusterChangeToggleVisibilty = true;
      } else {
        this.clusterChangeToggleVisibilty = false;
      }
      }); 
    this.getClusterList();
    this.isChrome = this.platform.BLINK;
    setTimeout(() => {
      this.closeWarning();
    }, 7000);
  }

  openQuickpanel() {
    this.layoutService.openQuickpanel();
  }

  openSidenav() {
    this.layoutService.openSidenav();
  }
  closeWarning() {
    this.isChrome = true;
  }

  getClusterList() {
    this.clusterService.clusterList$.subscribe({
      next: data => {
        this.clusterList = data || [];
      }
    });
  }

  onChangeCluster() {
    let urlSegments = window.location.href.split('/');
    urlSegments[3] = this.selectedCluster;
    const newUrl = '/' + urlSegments[3] + '/' + urlSegments[4] + '/' + urlSegments[5]  
    this.k8sService.changeClusterId(this.selectedCluster);
    this.router.navigate([newUrl], {} );
  }
}
