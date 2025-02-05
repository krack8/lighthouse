import { Injectable } from '@angular/core';
import { HttpService } from '@core-ui/services/http.service';
import { RequesterService } from '@core-ui/services/requester.service';
import { BehaviorSubject, Observable } from 'rxjs';
import { map, tap } from 'rxjs/operators';
import * as endpoint from './cluster.endpoint';
import { Utils } from '@shared-ui/utils';
import { Params } from '@angular/router';

interface StatusObject {
  [key: string]: any;
}

@Injectable()
export class ClusterService {
  // Cluster Router Name
  private clusterNameSubject = new BehaviorSubject('');
  clusterName$ = this.clusterNameSubject.asObservable();

  // Cluster Step Status
  private creationStepStatusSubject = new BehaviorSubject<StatusObject>({});
  creationStepStatus$ = this.creationStepStatusSubject.asObservable();

  // Cluster State
  private currentStateSubject = new BehaviorSubject('ACTIVE');
  currentState$ = this.currentStateSubject.asObservable();

  constructor(private httpService: HttpService) {}

  changeCreationStepStatus(statuses: any) {
    this.creationStepStatusSubject.next(statuses);
  }

  changeClusterName(name: any) {
    this.clusterNameSubject.next(name);
  }

  changeCurrentState(state: string) {
    this.currentStateSubject.next(state);
  }

  get clusterNameSnapshot(): string {
    return this.clusterNameSubject.value;
  }

  getCreationStepStatus(clusterName: string, force?: boolean): Observable<any> {
    if (force || !Object.keys(this.creationStepStatusSubject.value).length) {
      return this.mcGetClustersCreationStepStatus(clusterName).pipe(map(res => res.data));
    }
    return this.creationStepStatus$;
  }

  getManualCreationStepStatus(clusterName: string, force?: boolean): Observable<any> {
    if (force || !Object.keys(this.creationStepStatusSubject.value).length) {
      return this.mcGetManualClustersCreationStepStatus(clusterName).pipe(map(res => res.data));
    }
    return this.creationStepStatus$;
  }

  /*   APIs   */
  // Definition: Used To Get Cluster List
  // @queryParams: Pass QueryParams for Single Cluster Details Example: queryParams={ clusterName: this.clusterService.clusterNameSnapshot }
  // @dataFilter: Filter By Cluster State
  //
  mcGetClusters(queryParams?: any, clusterState?: string): Observable<any> {
    return this.httpService.get(endpoint.MC_GET_CLUSTERS, queryParams && queryParams).pipe(
      map(res => {
        if (clusterState) {
          const data = res.data.filter(item => item.currentState === clusterState);
          return {
            ...res,
            data
          };
        }
        return res;
      })
    );
  }

  /**
   * @description Used to get single cluster details
   * @param {Object} param - {clusterId} || {ClusterName: }
   * @return {Observable<any>} Observable<ClusterData as any>
   */
  mcGetCluster(clusterId: string): Observable<any> {
    return this.httpService.get(Utils.formatString(endpoint.MC_GET_CLUSTER, clusterId)).pipe(
      map(res => {
        if (res?.data) {
          this.changeCurrentState(res.data.currentState); // Used In Cluster Details Section
          return res?.data;
        }
      })
    );
  }

  mcGetManualCluster(clusterName: string): Observable<any> {
    return this.httpService.get(endpoint.MC_GET_MANUAL_CLUSTER, { clusterName }).pipe(
      map(res => {
        if (res?.data) {
          this.changeCurrentState(res.data.currentState); // Used In Cluster Details Section
          return res?.data;
        }
      })
    );
  }

  /**
   * @description Fetch cluster statuses
   * @effect Response loading on creationStepStatusSubject
   */
  mcGetClustersCreationStepStatus(clusterName: string): Observable<any> {
    return this.httpService.get(endpoint.MC_GET_CLUSTER_CREATION_STEP_STATUS, { clusterName }).pipe(
      tap((res: any) => {
        if (res.data) {
          this.changeCreationStepStatus(res.data);
        }
      })
    );
  }

  mcGetManualClustersCreationStepStatus(clusterName: string): Observable<any> {
    return this.httpService.get(endpoint.MC_GET_MANUAL_CLUSTER, { clusterName }).pipe(
      tap((res: any) => {
        if (res.data) {
          this.changeCreationStepStatus(res.data);
        }
      })
    );
  }

  mcGetClustersFullLog(clusterName: string): Observable<any> {
    return this.httpService.get(endpoint.MC_GET_CLUSTER_FULL_LOG + '/' + clusterName);
  }

  mcDeleteCluster(clusterId: string): Observable<any> {
    return this.httpService.delete(endpoint.MC_GET_CLUSTER + '/' + clusterId);
  }

  mcForceDeleteCluster(clusterId: string): Observable<any> {
    return this.httpService.delete(endpoint.MC_FORCE_DELETE + clusterId);
  }

  mcGetK8sVersions(): Observable<any> {
    return this.httpService.get(endpoint.MC_GET_AVAILABLE_K8S_VERSION_LIST);
  }

  mcGetAwsRegions(): Observable<any> {
    return this.httpService.get(endpoint.MC_GET_AVAILABLE_AWS_REGION_LIST);
  }

  mcGetGKENodegroup(payload: any): Observable<any> {
    return this.httpService.post(endpoint.MC_GET_AVAILABLE_NODE_TYPE_LIST, payload);
  }

  mcGetAwsNodegroup(qp: any): Observable<any> {
    return this.httpService.get(endpoint.MC_GET_AWS_AVAILABLE_NODE_TYPE_LIST, qp);
  }

  mcGetK8sVersionsForGcp(payload: any): Observable<any> {
    return this.httpService.post(endpoint.MC_GET_AVAILABLE_K8S_VERSION_LIST_FOR_GCP, payload);
  }

  mcVerifyDnsSettings(clusterId: any): Observable<any> {
    return this.httpService.get(endpoint.MC_VERIFY_DNS_SETTINGS + clusterId);
  }
  mcOnboardExistingCluster(formData?: any): Observable<any> {
    delete formData.inExistingNetwork;
    const payload = { ...formData };
    return this.httpService.post(endpoint.MC_ONBOARD_EXISTING_CLUSTER, payload);
  }

  // Api from facade
  getClusterById(id: string) {
    return this.httpService.get(Utils.formatString(endpoint.GET_CLUSTER_BY_ID, id));
  }

  updateCluster(payload: Record<string, any>) {
    return this.httpService.put(endpoint.UPATE_CLUSTER, payload);
  }

  getClusterReleaseInfoById() {
    return this.httpService.get(endpoint.GET_CLUSTER_RELEASE_BY_ID);
  }

  activateClusterUpgrade(clusterId: string) {
    return this.httpService.put(endpoint.ACTIVATE_CLUSTER_UPGRADE + clusterId, {});
  }

  updateClusterUpgradeStrategy(clusterId: string, payload: any) {
    return this.httpService.put(endpoint.UPDATE_CLUSTER_UPGRADE_STRAETGY + clusterId, payload);
  }
}
