import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, Resolve, Router, RouterStateSnapshot } from '@angular/router';
import { ToastrService } from '@sdk-ui/ui';
import { EMPTY, Observable } from 'rxjs';
import { catchError, map } from 'rxjs/operators';
import { ClusterService } from './cluster.service';

@Injectable()
export class ClusterResolver implements Resolve<any> {
  constructor(
    private clusterService: ClusterService,
    private router: Router,
    private toastr: ToastrService
  ) {}

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): Observable<any> {
    return this.clusterService.getCluster(route.params['clusterId']).pipe(
      map(_cluster => {
        if (_cluster.is_active) {
          this.router.navigate(['/clusters', route.params['clusterId'], 'k8s']);
        }
        return _cluster;
      }),
      catchError(err => {
        this.toastr.error(err['message'], 'Not Exist');
        this.router.navigate(['/clusters']);
        return EMPTY;
      })
    );
  }
}
