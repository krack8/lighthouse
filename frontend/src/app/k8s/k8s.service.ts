import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';
import * as pluralize from 'pluralize';
import * as endpoints from './k8s.endpoints';
import { HttpService } from '@core-ui/services';

@Injectable()
export class K8sService {
  private clusterIdSubject = new BehaviorSubject<string>('');
  clusterId$: Observable<string> = this.clusterIdSubject.asObservable();

  private clusterInfoSubject = new BehaviorSubject<object>({});
  clusterInfo$: Observable<object> = this.clusterInfoSubject.asObservable();

  constructor(private http: HttpService) {}

  get clusterIdSnapshot(): string {
    return this.clusterIdSubject.value;
  }

  changeClusterId(clusterId: string): void {
    this.clusterIdSubject.next(clusterId);
  }

  get clusterInfoSnapshot(): any {
    return this.clusterInfoSubject.value;
  }

  changeClusterInfo(clusterInfo: object): void {
    this.clusterInfoSubject.next(clusterInfo);
  }

  applyManifest(data: any, manifestFor: any): Observable<any> {
    const kind = data.kind;
    const qp = {
      resource: pluralize(kind)?.toLocaleLowerCase(),
      kind: kind
    };
    return this.http.post(endpoints.APPLY_MANIFEST, data, {
      ...qp,
      cluster_id: this.clusterIdSnapshot
    });
  }
}
