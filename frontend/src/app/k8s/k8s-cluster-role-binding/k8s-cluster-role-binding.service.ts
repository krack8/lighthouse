import { Injectable } from '@angular/core';
import { HttpService } from '@core-ui/services';
import { K8sService } from '@k8s/k8s.service';
import { Observable, of } from 'rxjs';
import { switchMap } from 'rxjs/operators';
import * as endpoints from './k8s-cluster-role-binding.endpoints';

@Injectable()
export class K8sClusterRoleBindingService {
  constructor(
    private k8sService: K8sService,
    private http: HttpService
  ) {}

  getClusterRoleBinding(queryParam?: any): Observable<any> {
    return this.k8sService.clusterId$.pipe(
      switchMap((clusterId: string) => {
        if (!clusterId) {
          return of({ data: null });
        }
        return this.http.get(endpoints.CLUSTER_ROLE_BINDING, { ...queryParam, cluster_id: clusterId });
      })
    );
  }

  getClusterRoleBindingDetails(name: string): Observable<any> {
    return this.k8sService.clusterId$.pipe(
      switchMap((clusterId: string) => {
        if (!clusterId) {
          return of({ data: null });
        }
        return this.http.get(endpoints.CLUSTER_ROLE_BINDING + '/' + name, { cluster_id: clusterId });
      })
    );
  }
  /* 
    Create & update
  */
  applyManifest(data: any, manifestFor: any): Observable<any> {
    switch (manifestFor) {
      case 'cluster-role-binding':
        return this.http.post(endpoints.CLUSTER_ROLE_BINDING, data, { cluster_id: this.k8sService.clusterIdSnapshot });
      default:
        return null;
    }
  }

  /* 
    Delete
  */
  deleteClusterRoleBinding(name: any): Observable<any> {
    return this.http.delete(endpoints.CLUSTER_ROLE_BINDING + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot
    });
  }
}
