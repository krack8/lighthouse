import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';

export const LOCAL_STORAGE_KEY = 'ngx-webstorage|kc-default-cluster';

@Injectable({
  providedIn: 'root'
})
export class SelectedClusterService {
  private clusterIdSubject: BehaviorSubject<string | null> = new BehaviorSubject<any>(null);
  public defaultClusterId$: Observable<string | null> = this.clusterIdSubject.asObservable();

  private selectedClusterIdSubject: BehaviorSubject<string | null> = new BehaviorSubject<any>(null);
  public selectedClusterId$: Observable<string | null> = this.selectedClusterIdSubject.asObservable();

  constructor() {
    const loc = localStorage.getItem(LOCAL_STORAGE_KEY);
    if (loc) {
      const cluster = JSON.parse(loc as string);
      if (cluster?.defaultClusterId) {
        this.clusterIdSubject.next(cluster.defaultClusterId);
      }
    }

    this.defaultClusterId$.subscribe(clusterId => {
      if (clusterId && this.selectedClusterId === null && window.location.href.split('/')[4] !== 'k8s') {
        this.setSelectedClusterId(clusterId);
      }
    });
  }

  saveDefaultCluster(cluster: any): void {
    localStorage.setItem(LOCAL_STORAGE_KEY, JSON.stringify(cluster));
    this.setDefaultClusterId(cluster.defaultClusterId);
  }

  clear() {
    localStorage.removeItem(LOCAL_STORAGE_KEY);
    this.setDefaultClusterId(null);
  }

  get defaultClusterId(): string | null {
    return this.clusterIdSubject.value;
  }

  setDefaultClusterId(clusterId: string): void {
    this.clusterIdSubject.next(clusterId);
  }

  // selected cluster id

  get selectedClusterId(): string | null {
    return this.selectedClusterIdSubject.value;
  }

  setSelectedClusterId(clusterId: string): void {
    this.selectedClusterIdSubject.next(clusterId);
  }
  
}