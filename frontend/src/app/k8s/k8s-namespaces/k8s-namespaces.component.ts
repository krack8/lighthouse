import { Component, OnDestroy, OnInit } from '@angular/core';
import { MatSelectChange } from '@angular/material/select';
import { ActivatedRoute, Router } from '@angular/router';
import icClose from '@iconify/icons-ic/close';
import icKeyboardBackspace from '@iconify/icons-ic/keyboard-backspace';
import icDown from '@iconify/icons-ic/twotone-keyboard-arrow-down';
import { K8sService } from '@k8s/k8s.service';
import { k8sRoutesMap, k8sRoutesPermissionMap } from '@shared-ui/utils';
import { filter, takeWhile } from 'rxjs/operators';
import { K8sNamespacesService } from './k8s-namespaces.service';

@Component({
  selector: 'kc-k8s-namespaces',
  templateUrl: './k8s-namespaces.component.html',
  styleUrls: ['./k8s-namespaces.component.scss']
})
export class K8sNamespacesComponent implements OnInit, OnDestroy {
  icClose = icClose;
  icDown = icDown;
  icKeyboardBackspace = icKeyboardBackspace;
  isAlive: boolean = true;
  selectedNamespace!: string;
  selectedResource!: string;
  clusterId!: string;
  namespaces: any[] = [];
  searchNamespaceTerm: string = '';
  resourcesSearchTerm: string = '';
  isLoading: boolean = false;
  resourceRouteMap = k8sRoutesMap;
  resourcePermissionMap = k8sRoutesPermissionMap;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private namespaceService: K8sNamespacesService,
    private k8sService: K8sService
  ) {}

  ngOnInit(): void {
    this.clusterId = this.k8sService.clusterIdSnapshot;
    this.route.queryParams
      .pipe(
        takeWhile(() => this.isAlive),
        filter(qp => !!qp?.namespace)
      )
      .subscribe(qp => {
        this.selectedNamespace = qp?.namespace;
        this.namespaceService.changeSelectedNamespace(qp?.namespace);
      });

    this.getNamespaces();
  }

  ngOnDestroy(): void {
    this.isAlive = false;
  }

  getNamespaces(): void {
    this.isLoading = true;
    this.namespaceService.forceGetNamespaces().subscribe(data => {
      const namespace_obj = [];
      data.forEach(namespace => {
        namespace_obj.push({ name: namespace });
      });
      this.namespaces = namespace_obj;
      this.isLoading = false;
    });
  }

  onChangeNamespace(e: MatSelectChange): void {
    const urlSegments = window.location.href.split('/');
    const path = urlSegments[6].split('?')[0]; // gets the name of the namespace entity ex: {.../namespace/pods?namespace=klovercloud} returns 'pods'
    this.router.navigate([path], {
      queryParams: { namespace: e.value },
      relativeTo: this.route
    });
  }

  onChangeNamespaceResource(e: MatSelectChange): void {
    this.router.navigate([e.value], {
      queryParams: { namespace: this.selectedNamespace },
      relativeTo: this.route
    });
    this.resourcesSearchTerm = '';
  }

  navigateToResources(resource: string) {
    this.resourcesSearchTerm = '';
    this.router.navigate([resource], {
      queryParams: { namespace: this.selectedNamespace },
      relativeTo: this.route
    });
  }
}
