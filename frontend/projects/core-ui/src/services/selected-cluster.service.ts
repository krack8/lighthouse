import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';

export const LOCAL_STORAGE_KEY = 'ngx-webstorage|kc-default-cluster';

@Injectable({
  providedIn: 'root'
})
export class SelectedClusterService {
  private clusterIdSubject: BehaviorSubject<string | null> = new BehaviorSubject<any>(null);
  public clusterId$: Observable<string | null> = this.clusterIdSubject.asObservable();

  constructor() {
    const loc = localStorage.getItem(LOCAL_STORAGE_KEY);
    if (loc) {
      const cluster = JSON.parse(loc as string);
      if (cluster?.defaultClusterId) {
        this.clusterIdSubject.next(cluster.defaultClusterId);
      }
    }
  }

  saveDefaultCluster(cluster: any): void {
    localStorage.setItem(LOCAL_STORAGE_KEY, JSON.stringify(cluster));
    this.setClusterId(cluster.defaultClusterId);
  }

  clear() {
    localStorage.removeItem(LOCAL_STORAGE_KEY);
    this.setClusterId(null);
  }

  get defaultClusterId(): string | null {
    return this.clusterIdSubject.value;
  }

  setClusterId(clusterId: string): void {
    this.clusterIdSubject.next(clusterId);
  }
}