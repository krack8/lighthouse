import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, Resolve, Router, RouterStateSnapshot } from '@angular/router';
import { ClusterService } from '@cluster/cluster.service';
import { SelectedClusterService } from '@core-ui/services/selected-cluster.service';
import { ToastrService } from '@sdk-ui/ui';
import { Observable, EMPTY } from 'rxjs';
import { catchError, map } from 'rxjs/operators';

@Injectable()
export class K8sResolver implements Resolve<any> {
  constructor(
    private clusterService: ClusterService,
    private router: Router,
    private toastr: ToastrService,
    private selectedClusterService: SelectedClusterService
  ) {}

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): Observable<any> {
    console.log('selected cluster id: ', this.selectedClusterService.defaultClusterId);
    return this.clusterService.getCluster(route.params['clusterId'] === ':clusterId'? this.selectedClusterService.defaultClusterId : route.params['clusterId']).pipe(
      catchError(err => {
        this.toastr.error(err['message'], 'Not Exist');
        this.router.navigate(['/clusters']);
        return EMPTY;
      })
    );
  }
}
