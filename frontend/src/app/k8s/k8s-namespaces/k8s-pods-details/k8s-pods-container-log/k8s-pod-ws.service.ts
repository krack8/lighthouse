import { Injectable } from '@angular/core';
import { K8sService } from '@k8s/k8s.service';
import { K8sNamespacesService } from '../../k8s-namespaces.service';
import { HttpClient, HttpParams } from '@angular/common/http';
import { RequesterService } from '@core-ui/services';

export const LOCAL_STORAGE_LIGHTHOUSE_LOG_KEY = 'ngx-webstorage|kc-lighthouse-logs-token';

@Injectable()
export class K8sPodWebSocketService {
  baseUrl: string;
  lighthouseLogUrl: string;

  constructor(
    private k8sService: K8sService,
    private namespaceService: K8sNamespacesService,
    private http: HttpClient,
    private requester: RequesterService
  ) {}

  getPodsWsUrl(name: string, qp?: any): string {
    const filterParams = qp;

    //const logsWsUrl = this.k8sService.clusterInfoSnapshot.lighthouseWsUrl;

    const logsWsUrl = 'http://localhost:8080';

    let uri = `${logsWsUrl}/ws/pod/logs/stream/${name}?cluster_id=${this.k8sService.clusterIdSnapshot}&namespace=${this.namespaceService.selectedNamespaceSnapshot}`;
    if (filterParams.lines) {
      uri += `&lines=${filterParams.lines}`;
    }
    if (filterParams.container) {
      uri += `&container=${filterParams.container}`;
    }
    if (filterParams.since) {
      uri += `&since=${filterParams.since}`;
    }
    if (filterParams.timestamps) {
      uri += `&timestamps=${filterParams.timestamps}`;
    }
    if (filterParams.access_token) {
      uri += `&access_token=${filterParams.access_token}`;
    }
    return uri;
  }

  getToken() {
    if (this.k8sService.clusterInfoSnapshot.lighthouseLogUrl) {
      const baseUrl = this.k8sService.clusterInfoSnapshot.lighthouseLogUrl;

      // let trimmedUrl = baseUrl.split('https://');
      // const tempBaseUrl = 'http://' + trimmedUrl[1];

      const url = `${baseUrl}/api/v1/authenticate`;

      const params = new HttpParams().set('cluster_id', this.k8sService.clusterIdSnapshot);
      return this.http.get(url, { params });
    }
  }
}
