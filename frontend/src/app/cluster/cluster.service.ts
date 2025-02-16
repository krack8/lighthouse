import { Injectable } from '@angular/core';
import { HttpService } from '@core-ui/services/http.service';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import * as endpoint from './cluster.endpoint';
import { Utils } from '@shared-ui/utils';
import { ICluster } from './cluster.model';

interface StatusObject {
  [key: string]: any;
}

@Injectable()
export class ClusterService {
  constructor(private httpService: HttpService) {}

  /*   APIs   */
  // Definition: Used To Get Cluster List
  // @queryParams: Pass QueryParams for Single Cluster Details Example: queryParams={ clusterName: this.clusterService.clusterNameSnapshot }
  // @dataFilter: Filter By Cluster State
  //
  getClusters(queryParams?: any): Observable<ICluster[]> {
    return this.httpService.get(endpoint.GET_CLUSTERS, queryParams && queryParams);
  }

  /**
   * @description Used to get single cluster details
   * @param {Object} param - {clusterId} || {ClusterName: }
   * @return {Observable<any>} Observable<ClusterData as any>
   */
  getCluster(clusterId: string): Observable<ICluster> {
    return this.httpService.get(Utils.formatString(endpoint.GET_CLUSTER, clusterId));
  }

  createCluster(payload: any): Observable<ICluster> {
    return this.httpService.post(endpoint.GET_CLUSTERS, payload);
  }

  getHelmChart(clusterId: string): Observable<any> {
    return this.httpService.get(Utils.formatString(endpoint.GET_HELM_CHART, clusterId));
  }

}
