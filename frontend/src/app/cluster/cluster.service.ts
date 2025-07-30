import { Injectable } from '@angular/core';
import { HttpService } from '@core-ui/services/http.service';
import { BehaviorSubject, Observable, of } from 'rxjs';
import * as endpoint from './cluster.endpoint';
import { Utils } from '@shared-ui/utils';
import { ICluster } from './cluster.model';
import { switchMap } from 'rxjs/operators';
import { ToastrService } from '@sdk-ui/ui';

interface StatusObject {
  [key: string]: any;
}

@Injectable()
export class ClusterService {

  private clusterListSubject: BehaviorSubject<any[]> = new BehaviorSubject<any[]>([]);
  public clusterList$: Observable<any[]> = this.clusterListSubject.asObservable();
  
  constructor(private httpService: HttpService, private toastrService: ToastrService) {}

  getAllClusterList(): void {
    this.getClusters().subscribe({
      next: (clusters: ICluster[]) => {
        this.setClusterList(clusters);
      },
      error: (error: any) => {
        this.toastrService.error('Failed to load clusters', 'Error');
        this.setClusterList([]);
      }
    });
  }

  get clusterListSnapshot(): ICluster[] {
    return this.clusterListSubject.getValue();
  }

  setClusterList(clusterList: ICluster[]): void {
    this.clusterListSubject.next(clusterList);
  }

  /*   APIs   */
  // Definition: Used To Get Cluster List
  // @queryParams: Pass QueryParams for Single Cluster Details Example: queryParams={ clusterName: this.clusterService.clusterNameSnapshot }
  // @dataFilter: Filter By Cluster State
  //
  getClusters(queryParams?: any): Observable<ICluster[]> {
    return this.httpService.get(endpoint.GET_CLUSTERS, queryParams && queryParams)
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

  deleteCluster(id: string): Observable<any> {
    return this.httpService.delete(endpoint.DELETE_CLUSTER + id);
  }

  getClustersList(){
    return this.httpService.get(endpoint.GET_CLUSTERS_LIST);
  }

  renameCluster(id: string, payload: any): Observable<any> {
    return this.httpService.put(endpoint.RENAME_CLUSTER + id, payload);
  }
}
