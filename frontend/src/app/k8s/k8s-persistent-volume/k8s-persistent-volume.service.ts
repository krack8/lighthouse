import { Injectable } from '@angular/core';
import { K8sService } from '@k8s/k8s.service';
import { Observable, of } from 'rxjs';
import { switchMap } from 'rxjs/operators';
import * as endpoints from './k8s-persistent-volume.endpoints';
import { HttpService } from '@core-ui/services';

@Injectable()
export class K8sPersistentVolumeService {
  constructor(
    private k8sService: K8sService,
    private http: HttpService
  ) {}

  getPersitentVolume(queryParam?: any): Observable<any> {
    return this.k8sService.clusterId$.pipe(
      switchMap((clusterId: string) => {
        if (!clusterId) {
          return of({ data: null });
        }
        return this.http.get(endpoints.PV_LIST, { ...queryParam, cluster_id: clusterId });
      })
    );
  }

  getPvDetails(name: string): Observable<any> {
    return this.k8sService.clusterId$.pipe(
      switchMap((clusterId: string) => {
        if (!clusterId) {
          return of({ data: null });
        }
        return this.http.get(endpoints.PV_LIST + '/' + name, { cluster_id: clusterId });
      })
    );
  }
  /* 
        Create & update
      */
  applyManifest(data: any, manifestFor: any): Observable<any> {
    switch (manifestFor) {
      case 'persistent-volume':
        return this.http.post(endpoints.PV_LIST, data, { cluster_id: this.k8sService.clusterIdSnapshot });
      default:
        return null;
    }
  }

  /* 
      Delete
    */
  deletePersistentVolume(name: any): Observable<any> {
    return this.http.delete(endpoints.PV_LIST + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot
    });
  }
}
