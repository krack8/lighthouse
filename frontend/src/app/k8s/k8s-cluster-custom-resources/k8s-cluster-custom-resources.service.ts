import { Injectable } from '@angular/core';
import { K8sService } from '@k8s/k8s.service';
import { Observable, of } from 'rxjs';
import { switchMap } from 'rxjs/operators';
import * as endpoints from './k8s-cluster-custom-resources.endpoints';
import { HttpService } from '@core-ui/services';

@Injectable()
export class K8sClusterCustomResourcesService {
  constructor(
    private k8sService: K8sService,
    private http: HttpService
  ) {}

  getCustomResourcesDefination(queryParams?: any): Observable<any> {
    return this.http.get(endpoints.CLUSTER_CUSTOM_RESOURCES_DEFINATION, { cluster_id: this.k8sService.clusterIdSnapshot, ...queryParams });
  }

  getCustomResourceDefinationDetails(name: string): Observable<any> {
    return this.http.get(endpoints.CLUSTER_CUSTOM_RESOURCES_DEFINATION + '/' + name, { cluster_id: this.k8sService.clusterIdSnapshot });
  }

  getCustomResources(params: any): Observable<any> {
    return this.http.get(endpoints.CLUSTER_CUSTOM_RESOURCES, { ...params, cluster_id: this.k8sService.clusterIdSnapshot });
  }

  getCustomResourceDetails(name: string, qp: any): Observable<any> {
    return this.http.get(endpoints.CLUSTER_CUSTOM_RESOURCES + '/' + name, { ...qp, cluster_id: this.k8sService.clusterIdSnapshot });
  }

  /* 
    Create & update
  */
  applyManifest(data: any, manifestFor: any, queryParams?: any): Observable<any> {
    switch (manifestFor) {
      case 'custom-resources':
        return this.http.post(endpoints.CLUSTER_CUSTOM_RESOURCES, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          ...queryParams
        });
      case 'crd':
        return this.http.post(endpoints.CLUSTER_CUSTOM_RESOURCES_DEFINATION, data, {
          cluster_id: this.k8sService.clusterIdSnapshot
        });
      default:
        return null;
    }
  }

  /* 
    Delete
  */
  deleteCustomResources(name: any, queryParams: any): Observable<any> {
    return this.http.delete(endpoints.CLUSTER_CUSTOM_RESOURCES + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      ...queryParams
    });
  }

  deleteCustomResourcesDefination(name: any): Observable<any> {
    return this.http.delete(endpoints.CLUSTER_CUSTOM_RESOURCES_DEFINATION + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot
    });
  }
}
