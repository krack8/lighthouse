import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, Resolve, Router, RouterStateSnapshot } from '@angular/router';
import { ToastrService } from '@sdk-ui/ui';
import { EMPTY, Observable, of } from 'rxjs';
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
    return this.clusterService.mcGetCluster(route.params['clusterId']).pipe(
      map(res => {
        return res?.data;
      }),
      catchError(err => {
        this.toastr.error(err['message'], 'Not Exist');
        this.router.navigate(['/clusters']);
        return EMPTY;
      })
    );
  }
}
