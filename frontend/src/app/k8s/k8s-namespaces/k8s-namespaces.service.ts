import { Injectable } from '@angular/core';
import { K8sService } from '@k8s/k8s.service';
import { BehaviorSubject, Observable, of } from 'rxjs';
import { map, switchMap } from 'rxjs/operators';
import * as endpoints from './k8s-namespaces.endpoints';
import { Utils } from '@shared-ui/utils';
import { HttpService } from '@core-ui/services';

@Injectable()
export class K8sNamespacesService {
  private selectedNamespaceSubject = new BehaviorSubject<string>('');
  selectedNamespace$: Observable<string> = this.selectedNamespaceSubject.asObservable();

  private namespacesSubject = new BehaviorSubject<any[]>([]);
  namespaces$: Observable<any[]> = this.namespacesSubject.asObservable();

  constructor(
    private k8sService: K8sService,
    private http: HttpService
  ) {}

  changeSelectedNamespace(_namespace: string): void {
    this.selectedNamespaceSubject.next(_namespace);
  }

  get selectedNamespaceSnapshot(): string {
    return this.selectedNamespaceSubject.value;
  }

  changeNamespace(items: any[]): void {
    this.namespacesSubject.next(items);
  }

  get namespacesSnapshot(): any[] {
    return this.namespacesSubject.value;
  }

  forceGetNamespaces(): Observable<any> {
    return this.namespaces$.pipe(
      switchMap(data => {
        if (data?.length) {
          return data;
        }
        return this.http
          .get(endpoints.NAMESPACE_NAME_LIST, { cluster_id: this.k8sService.clusterIdSnapshot })
          .pipe(map((res: any) => res?.data || []));
      })
    );
  }

  //new api

  getNamespaces(queryParam?: any): Observable<any> {
    return this.k8sService.clusterId$.pipe(
      switchMap((clusterId: string) => {
        if (!clusterId) {
          return of({ data: null });
        }
        return this.http.get(endpoints.NAMESPACE, { cluster_id: clusterId, ...queryParam });
      })
    );
  }

  getNamespacesDetailsV1(name: string): Observable<any> {
    return this.k8sService.clusterId$.pipe(
      switchMap((clusterId: string) => {
        if (!clusterId) {
          return of({ data: null });
        }
        return this.http.get(endpoints.NAMESPACE + '/' + name, { cluster_id: clusterId });
      })
    );
  }

  getDeploymentsV1(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_DEPLOYMENT, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getDeploymentStats(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_DEPLOYMENT_STATS, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot,
      scope: ''
    });
  }

  getDeploymentsDetailsV1(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_DEPLOYMENT + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getPods(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_POD, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getPodStats(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_POD_STATS, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getPodsDetails(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_POD + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getReplicaSets(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_REPLICA_SET, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getReplicaSetsDetails(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_REPLICA_SET + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getStatefulSets(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_STATEFUL_SET, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getStatefulSetStats(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_STATEFUL_SET_STATS, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getStatefulSetsDetails(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_STATEFUL_SET + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getDaemonSets(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_DAEMON_SET, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getDaemonSetsDetails(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_DAEMON_SET + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getDaemonSetsStats(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_DAEMON_SET_STATS, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getSecrets(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_SECRET, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getSecretsDetails(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_SECRET + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getGateway(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_GATEWAY, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getGatewayDetails(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_GATEWAY + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getVirtualService(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_VIRTUAL_SERVICE, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getVirtualServiceDetails(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_VIRTUAL_SERVICE + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getService(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_SERVICE, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getServiceDetails(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_SERVICE + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getServiceAccounts(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_SERVICE_ACCOUNT, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getServiceAccountsDetails(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_SERVICE_ACCOUNT + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getConfigMaps(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_CONFIG_MAP, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getConfigMapsDetails(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_CONFIG_MAP + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getPvcList(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_PERSISTENT_VOLUME, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }
  getPvcDetails(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_PERSISTENT_VOLUME + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getIngresses(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_INGRESS, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }
  getIngressDetails(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_INGRESS + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getCertificates(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_CERTIFICATE, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }
  getCertificatesDetails(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_CERTIFICATE + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getRole(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_ROLE, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }
  getRoleDetails(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_ROLE + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getRoleBindings(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_ROLE_BINDING, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }
  getRoleBindingsDetails(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_ROLE_BINDING + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getJob(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_JOB, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }
  getJobDetails(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_JOB + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getCronJob(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_CRON_JOB, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }
  getCronJobDetails(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_CRON_JOB + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getNetwokPolicy(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_NETWORK_POLICY, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }
  getNetwokPolicyDetails(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_NETWORK_POLICY + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getResourceQuota(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_RESOURCE_QUOTA, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }
  getResourceQuotaDetails(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_RESOURCE_QUOTA + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }
  getEndpoints(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_ENDPOINTS, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }
  getEndpointsDetails(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_ENDPOINTS + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }
  getEndpointSlice(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_ENDPOINT_SLICE, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }
  getEndpointSliceDetails(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_ENDPOINT_SLICE + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }
  getPdb(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_PDB, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }
  getPdbDetails(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_PDB + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }
  getControllerRevision(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_CONTROLLER_REVISION, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }
  getControllerRevisionDetails(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_CONTROLLER_REVISION + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }
  getReplicationController(queryParam?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_REPLICATION_CONTROLLER, {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }
  getReplicationControllerDetails(name: string): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_REPLICATION_CONTROLLER + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  // Create & update

  applyManifest(data: any, manifestFor: any): Observable<any> {
    switch (manifestFor) {
      case 'namespace':
        return this.http.post(endpoints.NAMESPACE, data, {
          cluster_id: this.k8sService.clusterIdSnapshot
        });
      case 'config-map':
        return this.http.post(endpoints.NAMESPACE_CONFIG_MAP, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      case 'secrets':
        return this.http.post(endpoints.NAMESPACE_SECRET, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      case 'service':
        return this.http.post(endpoints.NAMESPACE_SERVICE, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      case 'pods':
        return this.http.post(endpoints.NAMESPACE_POD, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      case 'replica-sets':
        return this.http.post(endpoints.NAMESPACE_REPLICA_SET, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      case 'daemonset':
        return this.http.post(endpoints.NAMESPACE_DAEMON_SET, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      case 'deployments':
        return this.http.post(endpoints.NAMESPACE_DEPLOYMENT, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      case 'service-account':
        return this.http.post(endpoints.NAMESPACE_SERVICE_ACCOUNT, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      case 'certificates':
        return this.http.post(endpoints.NAMESPACE_CERTIFICATE, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      case 'role':
        return this.http.post(endpoints.NAMESPACE_ROLE, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      case 'job':
        return this.http.post(endpoints.NAMESPACE_JOB, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      case 'cron-job':
        return this.http.post(endpoints.NAMESPACE_CRON_JOB, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      case 'role-binding':
        return this.http.post(endpoints.NAMESPACE_ROLE_BINDING, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      case 'network-policy':
        return this.http.post(endpoints.NAMESPACE_NETWORK_POLICY, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      case 'resource-quota':
        return this.http.post(endpoints.NAMESPACE_RESOURCE_QUOTA, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      case 'gateway':
        return this.http.post(endpoints.NAMESPACE_GATEWAY, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      case 'virtual-service':
        return this.http.post(endpoints.NAMESPACE_VIRTUAL_SERVICE, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      case 'pvc':
        return this.http.post(endpoints.NAMESPACE_PERSISTENT_VOLUME, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      case 'ingresses':
        return this.http.post(endpoints.NAMESPACE_INGRESS, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      case 'stateful-set':
        return this.http.post(endpoints.NAMESPACE_STATEFUL_SET, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      case 'endpoints':
        return this.http.post(endpoints.NAMESPACE_ENDPOINTS, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      case 'endpoint-slice':
        return this.http.post(endpoints.NAMESPACE_ENDPOINT_SLICE, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      case 'pdb':
        return this.http.post(endpoints.NAMESPACE_PDB, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      case 'controller-revision':
        return this.http.post(endpoints.NAMESPACE_CONTROLLER_REVISION, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      case 'replication-controller':
        return this.http.post(endpoints.NAMESPACE_REPLICATION_CONTROLLER, data, {
          cluster_id: this.k8sService.clusterIdSnapshot,
          namespace: this.selectedNamespaceSnapshot
        });
      default:
        return null;
    }
  }

  // Delete

  deleteNamespace(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot
    });
  }

  deleteNamespacePods(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_POD + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  deleteNamespaceReplicasets(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_REPLICA_SET + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  deleteNamespaceService(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_SERVICE + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  deleteNamespaceConfigMap(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_CONFIG_MAP + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  deleteNamespaceSecrets(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_SECRET + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  deleteNamespaceDeployment(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_DEPLOYMENT + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  deleteNamespaceDaemonset(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_DAEMON_SET + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  deleteNamespaceServiceAccount(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_SERVICE_ACCOUNT + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  deleteNamespaceCertificates(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_CERTIFICATE + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  deleteNamespaceRole(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_ROLE + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  deleteNamespaceRoleBinding(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_ROLE_BINDING + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  deleteNamespaceNetworkPolicy(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_NETWORK_POLICY + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  deleteNamespaceResourceQuota(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_RESOURCE_QUOTA + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  deleteNamespacePvc(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_PERSISTENT_VOLUME + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  deleteNamespaceIngresses(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_INGRESS + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  deleteNamespaceStatefulset(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_STATEFUL_SET + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  deleteJob(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_JOB + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  deleteCronJob(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_CRON_JOB + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  deleteNamespaceGateway(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_GATEWAY + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  deleteNamespaceVirtualService(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_VIRTUAL_SERVICE + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  deleteNamespaceEndpoint(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_ENDPOINTS + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  deleteNamespaceEndpointSlice(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_ENDPOINT_SLICE + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  deleteNamespacePdb(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_PDB + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  deleteNamespaceControllerRevision(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_CONTROLLER_REVISION + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  deleteNamespaceReplicationController(name: any): Observable<any> {
    return this.http.delete(endpoints.NAMESPACE_REPLICATION_CONTROLLER + '/' + name, {
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getLogsV1(params: any, pod: string): Observable<any> {
    return this.http.get(Utils.formatString(endpoints.GET_LOGS_V1, pod), {
      ...params,
      cluster_id: this.k8sService.clusterIdSnapshot
    });
  }

  getNamespaceDeploymentPodList(deploymentName: string, replicaset: string, queryParam?: any): Observable<any> {
    return this.http.get(Utils.formatString(endpoints.NAMESPACE_DEPLOYMENT_POD_LIST, deploymentName), {
      ...queryParam,
      rs: replicaset,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getNamespaceStatefulsetsPodList(statefulsetName: string, queryParam?: any): Observable<any> {
    return this.http.get(Utils.formatString(endpoints.NAMESPACE_STATEFUL_SETS_POD_LIST, statefulsetName), {
      ...queryParam,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }

  getEvents(name?: any): Observable<any> {
    return this.http.get(endpoints.NAMESPACE_EVENTS, {
      involved_object_name: name,
      cluster_id: this.k8sService.clusterIdSnapshot,
      namespace: this.selectedNamespaceSnapshot
    });
  }
}
