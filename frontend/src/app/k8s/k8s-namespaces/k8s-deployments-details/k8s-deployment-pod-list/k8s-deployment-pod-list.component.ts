import { Component, Inject, Input, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import icSearch from '@iconify/icons-ic/search';
import { ToastrService } from '@sdk-ui/ui';
import icRefresh from '@iconify/icons-ic/twotone-refresh';
import icLabel from '@iconify/icons-ic/twotone-label';
import icMoreVert from '@iconify/icons-ic/twotone-more-vert';
import icCross from '@iconify/icons-ic/cancel';
import { distinctUntilChanged, map, share, takeWhile, throttleTime, tap } from 'rxjs/operators';
import { DOCUMENT } from '@angular/common';
import { fromEvent } from 'rxjs';
import { K8sNamespacesService } from '@k8s/k8s-namespaces/k8s-namespaces.service';

@Component({
  selector: 'kc-k8s-deployment-pod-list',
  templateUrl: './k8s-deployment-pod-list.component.html',
  styleUrls: ['./k8s-deployment-pod-list.component.scss']
})
export class K8sDeploymentPodListComponent implements OnInit {
  @Input() requestData: any;
  icSearch = icSearch;
  icCross = icCross;
  icLabel = icLabel;
  icMoreVert = icMoreVert;
  icRefresh = icRefresh;
  searchBy = 'name';
  isAlive: boolean = true;
  loadingSpanner: boolean = false;
  data: any;
  searchTerm: string = '';
  queryParams: any;
  stats: any;
  statsLoaded: boolean = false;
  resourceToken: string = '';
  loadMoreData: boolean = false;
  remaining: string;

  constructor(
    private namespaceService: K8sNamespacesService,
    private route: ActivatedRoute,
    private toastr: ToastrService,
    @Inject(DOCUMENT) public document: any,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.queryParams = this.route.snapshot.queryParams;
    this.getPodList();
  }

  ngOnDestroy(): void {
    this.isAlive = false;
  }

  reloadList() {
    this.getPodList();
  }

  getPodList(queryParams?: any): void {
    this.loadingSpanner = true;
    this.namespaceService
      .getNamespaceDeploymentPodList(this.requestData?.deploymentName, this.requestData?.replicaset, queryParams)
      .subscribe({
        next: res => {
          if (res?.status === 'success') {
            this.resourceToken = res.data?.Resource;
            this.remaining = res.data?.Remaining;
            this.data = res?.data?.PodList;
          }
          if (res?.status === 'error') {
            this.toastr.error(res?.message);
          }
          this.loadingSpanner = false;
        },
        error: err => {
          this.toastr.error(err?.message);
        }
      });
  }

  onSearch() {
    if (this.searchBy === 'label') {
      const keyValuePairs = this.searchTerm.split(',');
      const jsonObject = {};
      keyValuePairs.forEach(pair => {
        if (pair.includes(':')) {
          const [key, value] = pair.split(':');
          jsonObject[key] = value;
        } else {
          this.toastr.error('Incorrect format for label search. Please use key:value format');
          return;
        }
      });
      if (JSON.stringify(jsonObject) === '{}') {
        return;
      }
      const jsonString = JSON.stringify(jsonObject);
      const qp = { labels: jsonString };
      this.getPodList(qp);
    }
    if (this.searchBy === 'name') {
      const qp = { q: this.searchTerm };
      this.getPodList(qp);
    }
  }

  clearSearch() {
    this.getPodList();
    this.searchTerm = '';
  }

  handleInputChange() {
    if (this.searchTerm.length === 0) {
      this.getPodList();
    }
  }

  ngAfterContentInit(): void {
    const content = this.document.querySelector('.sidenav-content');
    const scroll$ = fromEvent(content, 'scroll').pipe(
      takeWhile(() => this.isAlive),
      throttleTime(10), // only emit every 10 ms
      map((): boolean => {
        return content.offsetHeight + content.scrollTop + 80 >= content.scrollHeight;
      }),
      distinctUntilChanged(), // only emit when scrolling direction changed
      share() // share a single subscription to the underlying sequence in case of multiple subscribers
    );

    scroll$.subscribe((isBottom: boolean) => {
      if (isBottom && this.resourceToken.length !== 0 && !this.loadMoreData) {
        this.loadMore();
      }
    });
  }

  loadMore() {
    this.loadMoreData = true;
    const queryParam = { continue: this.resourceToken };

    const namespace$ = this.namespaceService.selectedNamespace$.pipe(tap(() => {}));
    this.namespaceService
      .getPods(queryParam)
      .pipe(takeWhile(() => this.isAlive))
      .subscribe({
        next: res => {
          if (res?.status === 'error') {
            this.toastr.error(res?.message);
          }
          this.loadMoreData = false;
          this.remaining = res.data.Remaining;
          this.resourceToken = res.data.Resource || '';
          this.data = this.data.concat(res.data.PodList) || [];
        },
        error: err => {
          this.loadingSpanner = false;
          this.toastr.error(err?.message);
        }
      });
  }

  calcMemory(containers, _type = 'requests'): any {
    if (containers.length === 1) {
      try {
        if (containers[0].resources[_type].memory !== undefined) {
          return containers[0].resources[_type].memory;
        } else {
          return '-';
        }
      } catch (err) {
        return '-';
      }
    } else if (containers.length > 1) {
      let count = 0;
      containers.map(container => {
        if (container.resources[_type] && container.resources[_type].memory && container.resources[_type].memory.includes('Gi')) {
          const asMi = parseFloat(container.resources[_type].memory) * 1024;
          count = count + asMi;
        } else if (container.resources[_type] && container.resources[_type].memory && container.resources[_type].memory.includes('Mi')) {
          const asNumber = parseFloat(container.resources[_type].memory);
          count = count + asNumber;
        }
      });
      if (count < 1024) {
        return `${count}Mi`;
      } else if (count < 1024 ** 2) {
        return `${(count / 1024).toFixed(2)}Gi`;
      } else if (count < 1024 ** 3) {
        return `${(count / 1024 ** 2).toFixed(2)}Tb`;
      }
      return count || '-';
    }
    return '-';
  }
  calcCpu(containers, _type = 'requests') {
    if (containers.length === 1) {
      try {
        return containers[0].resources[_type].cpu;
      } catch (err) {
        return '-';
      }
    } else if (containers.length > 1) {
      let count = 0;
      containers.map(container => {
        if (container.resources[_type] && container.resources[_type].cpu && container.resources[_type].cpu.includes('m')) {
          const asNumber = parseFloat(container.resources[_type].cpu);
          count = count + asNumber;
        } else if (container.resources[_type] && container.resources[_type].cpu) {
          const asMillie = parseFloat(container.resources[_type].cpu) * 1000;
          count = count + asMillie;
        }
      });
      if (count < 1000) {
        return `${count}m`;
      } else if (count < 1000 ** 2) {
        return `${(count / 1000).toFixed(2)}`;
      }
      return count || '-';
    }
    return '-';
  }

  containerRunning(containers: any[]) {
    let count = 0;
    if (containers) {
      containers.forEach(container => {
        if (container.ready === true) {
          count++;
        }
      });
      return count;
    } else {
      return '0';
    }
  }

  navigate(podName: string, namespace: string) {
    const url = this.router
      .createUrlTree([`../../pods/${podName}`], {
        relativeTo: this.route,
        queryParams: { namespace: namespace }
      })
      .toString();

    const fullUrl = window.location.origin + '/' + url;
    window.open(fullUrl, '_blank');
  }
}
