import { Injectable } from '@angular/core';
import { HttpService } from '@core-ui/services';
import { K8sService } from '@k8s/k8s.service';
import { Observable, of } from 'rxjs';
import { switchMap } from 'rxjs/operators';
import * as endpoints from './k8s-cluster-role.endpoints';

@Injectable()
export class K8sClusterRoleService {
  constructor(
    private k8sService: K8sService,
    private http: HttpService
  ) {}

  getClusterRole(queryParam?: any): Observable<any> {
    return this.k8sService.clusterId$.pipe(
      switchMap((clusterId: string) => {
        if (!clusterId) {
          return of({ data: null });
        }
        return this.http.get(endpoints.CLUSTER_ROLE, { ...queryParam, cluster_id: clusterId });
      })
    );
  }

  getClusterRoleDetails(name: string): Observable<any> {
    return this.k8sService.clusterId$.pipe(
      switchMap((clusterId: string) => {
        if (!clusterId) {
          return of({ data: null });
        }
        return this.http.get(endpoints.CLUSTER_ROLE + '/' + name, { cluster_id: clusterId });
      })
    );
  }
  /* 
    Create & update
  */
  applyManifest(data: any, manifestFor: any): Observable<any> {
    switch (manifestFor) {
      case 'cluster-role':
        return this.http.post(endpoints.CLUSTER_ROLE, data, { cluster_id: this.k8sService.clusterIdSnapshot });
      default:
        return null;
    }
  }

  /* 
    Delete
  */
  deleteClusterRole(name: any): Observable<any> {
    return this.http.delete(endpoints.CLUSTER_ROLE + '/' + name, { cluster_id: this.k8sService.clusterIdSnapshot });
  }
}
