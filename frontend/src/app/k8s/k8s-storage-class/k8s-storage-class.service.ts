import { Injectable } from '@angular/core';
import { HttpService } from '@core-ui/services';
import { K8sService } from '@k8s/k8s.service';
import { Observable, of } from 'rxjs';
import { switchMap } from 'rxjs/operators';
import * as endpoints from './k8s-storage-class.endpoints';

@Injectable()
export class K8sStorageClassService {
  constructor(
    private k8sService: K8sService,
    private http: HttpService
  ) {}

  getStorageClass(queryParam?: any): Observable<any> {
    return this.k8sService.clusterId$.pipe(
      switchMap((clusterId: string) => {
        if (!clusterId) {
          return of({ data: null });
        }
        return this.http.get(endpoints.STORAGE_CLASS_LIST, { ...queryParam, cluster_id: clusterId });
      })
    );
  }

  getStorageClassDetails(name: string): Observable<any> {
    return this.k8sService.clusterId$.pipe(
      switchMap((clusterId: string) => {
        if (!clusterId) {
          return of({ data: null });
        }
        return this.http.get(endpoints.STORAGE_CLASS_LIST + '/' + name, { cluster_id: clusterId });
      })
    );
  }
  /* 
        Create & update
      */
  applyManifest(data: any, manifestFor: any): Observable<any> {
    switch (manifestFor) {
      case 'storage-class':
        return this.http.post(endpoints.STORAGE_CLASS_LIST, data, { cluster_id: this.k8sService.clusterIdSnapshot });
      default:
        return null;
    }
  }

  /* 
      Delete
    */
  deleteStorageClass(name: any): Observable<any> {
    return this.http.delete(endpoints.STORAGE_CLASS_LIST + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot
    });
  }
}
