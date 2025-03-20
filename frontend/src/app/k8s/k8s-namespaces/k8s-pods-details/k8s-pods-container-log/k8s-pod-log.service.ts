import { Inject, Injectable, NgZone, Optional } from '@angular/core';
import { APP_ENV } from '@core-ui/constants';
import { IAppEnv } from '@core-ui/interfaces';
import { RequesterService } from '@core-ui/services';
import { K8sService } from '@k8s/k8s.service';
import { Observable } from 'rxjs';
import { K8sNamespacesService } from '../../k8s-namespaces.service';

export const LOCAL_STORAGE_LIGHTHOUSE_LOG_KEY = 'ngx-webstorage|kc-lighthouse-logs-token';

@Injectable()
export class K8sPodLogService {
  baseUrl: string;

  constructor(
    @Optional() @Inject(APP_ENV) private _env: IAppEnv,
    private k8sService: K8sService,
    private namespaceService: K8sNamespacesService,
    private requester: RequesterService,
    private ngZone: NgZone
  ) {
    this.baseUrl = _env?.apiEndPoint;
  }

  getPodsWsUrl(name: string, qp?: any): string {
    const filterParams = qp;

    const logsWsUrl = this.k8sService.clusterInfoSnapshot.lighthouseWsUrl;

    let uri = `${this.baseUrl}/v1/pod/${name}/logs/stream?cluster_id=${this.k8sService.clusterIdSnapshot}&namespace=${this.namespaceService.selectedNamespaceSnapshot}`;
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

  getPodLogsStream(url: string): Observable<string> {
    return new Observable<string>((observer) => {  
      const xhr = new XMLHttpRequest();
      xhr.open('GET', url, true);
      xhr.setRequestHeader('Authorization', `Bearer ${this.requester.get().token}`);
      xhr.onreadystatechange = () => {
        if (xhr.readyState >= 3) {
          if (xhr.status === 200) {
            this.ngZone.run(() => {
              observer.next(xhr.responseText);
            });
          } else if (xhr.readyState === 4 && xhr.status !== 200) {
            this.ngZone.run(() => {
              observer.error(`Error: ${xhr.status} - ${xhr.statusText}`);
            });
          }
        }
      };
      xhr.onerror = (error) => {
        this.ngZone.run(() => {
          observer.error(error); 
        });
      };

      xhr.send();
        return () => {
        xhr.abort();
      };
    });
  }
}













