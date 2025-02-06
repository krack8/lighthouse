import { Injectable } from '@angular/core';
import { HttpService } from '@core-ui/services';
import { K8sService } from '@k8s/k8s.service';
import { Observable, of } from 'rxjs';
import { switchMap } from 'rxjs/operators';
import * as endpoints from './k8s-nodes.endpoints';

@Injectable()
export class K8sNodesService {
  constructor(
    private k8sService: K8sService,
    private http: HttpService
  ) {}

  getNodes(queryParam?: any): Observable<any> {
    return this.k8sService.clusterId$.pipe(
      switchMap((clusterId: string) => {
        if (!clusterId) {
          return of({ data: null });
        }
        return this.http.get(endpoints.NODE_LIST, { ...queryParam, cluster_id: clusterId });
      })
    );
  }

  getNodeDetails(name: string): Observable<any> {
    return this.k8sService.clusterId$.pipe(
      switchMap((clusterId: string) => {
        if (!clusterId) {
          return of({ data: null });
        }
        return this.http.get(endpoints.NODE_LIST + '/' + name, { cluster_id: clusterId });
      })
    );
  }

  nodeCordonUncordon(name: string): Observable<any> {
    return this.http.get(endpoints.NODE_CORDON_UNCORDON + name, {
      cluster_id: this.k8sService.clusterIdSnapshot
    });
  }

  taintNode(name: string, data: any): Observable<any> {
    return this.http.post(endpoints.TAINT_NODE + name, data, {
      cluster_id: this.k8sService.clusterIdSnapshot
    });
  }

  untaintNode(name: string, data: any): Observable<any> {
    return this.http.post(endpoints.UNTAINT_NODE + name, data, {
      cluster_id: this.k8sService.clusterIdSnapshot
    });
  }
}
