import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { RoleGuardService } from '@core-ui/guards';
import { K8sCertificatesDetailsComponent } from './k8s-certificates-details/k8s-certificates-details.component';
import { K8sCertificatesComponent } from './k8s-certificates/k8s-certificates.component';
import { K8sConfigMapsDetailsComponent } from './k8s-config-maps-details/k8s-config-maps-details.component';
import { K8sConfigMapsComponent } from './k8s-config-maps/k8s-config-maps.component';
import { K8sCronJobDetailsComponent } from './k8s-cron-job-details/k8s-cron-job-details.component';
import { K8sCronJobComponent } from './k8s-cron-job/k8s-cron-job.component';
import { K8sDaemonSetsDetailsComponent } from './k8s-daemon-sets-details/k8s-daemon-sets-details.component';
import { K8sDaemonSetsComponent } from './k8s-daemon-sets/k8s-daemon-sets.component';
import { K8sDeploymentsDetailsComponent } from './k8s-deployments-details/k8s-deployments-details.component';
import { K8sDeploymentsComponent } from './k8s-deployments/k8s-deployments.component';
import { K8sGatewayDetailsComponent } from './k8s-gateway-details/k8s-gateway-details.component';
import { K8sGatewayComponent } from './k8s-gateway/k8s-gateway.component';
import { K8sIngressesDetailsComponent } from './k8s-ingresses-details/k8s-ingresses-details.component';
import { K8sIngressesComponent } from './k8s-ingresses/k8s-ingresses.component';
import { K8sJobDetailsComponent } from './k8s-job-details/k8s-job-details.component';
import { K8sJobComponent } from './k8s-job/k8s-job.component';
import { K8sNamespacesListComponent } from './k8s-namespaces-list/k8s-namespaces-list.component';
import { K8sNamespacesComponent } from './k8s-namespaces.component';
import { K8sNetworkPolicyDetailsComponent } from './k8s-network-policy-details/k8s-network-policy-details.component';
import { K8sNetworkPolicyComponent } from './k8s-network-policy/k8s-network-policy.component';
import { K8sPodsDetailsComponent } from './k8s-pods-details/k8s-pods-details.component';
import { K8sPodsComponent } from './k8s-pods/k8s-pods.component';
import { K8sPvcsDetailsComponent } from './k8s-pvcs-details/k8s-pvcs-details.component';
import { K8sPvcsComponent } from './k8s-pvcs/k8s-pvcs.component';
import { K8sReplicaSetsDetailsComponent } from './k8s-replica-sets-details/k8s-replica-sets-details.component';
import { K8sReplicaSetsComponent } from './k8s-replica-sets/k8s-replica-sets.component';
import { K8sResourceQuotaDetailsComponent } from './k8s-resource-quota-details/k8s-resource-quota-details.component';
import { K8sResourceQuotaComponent } from './k8s-resource-quota/k8s-resource-quota.component';
import { K8sRoleBindingDetailsComponent } from './k8s-role-binding-details/k8s-role-binding-details.component';
import { K8sRoleBindingComponent } from './k8s-role-binding/k8s-role-binding.component';
import { K8sRoleDetailsComponent } from './k8s-role-details/k8s-role-details.component';
import { K8sRoleComponent } from './k8s-role/k8s-role.component';
import { K8sSecretsDetailsComponent } from './k8s-secrets-details/k8s-secrets-details.component';
import { K8sSecretsComponent } from './k8s-secrets/k8s-secrets.component';
import { K8sServiceAccountsDetailsComponent } from './k8s-service-accounts-details/k8s-service-accounts-details.component';
import { K8sServiceAccountsComponent } from './k8s-service-accounts/k8s-service-accounts.component';
import { K8sServiceDetailsComponent } from './k8s-service-details/k8s-service-details.component';
import { K8sServiceComponent } from './k8s-service/k8s-service.component';
import { K8sStatefulSetsDetailsComponent } from './k8s-stateful-sets-details/k8s-stateful-sets-details.component';
import { K8sStatefulSetsComponent } from './k8s-stateful-sets/k8s-stateful-sets.component';
import { K8sVirtualServiceDetailsComponent } from './k8s-virtual-service-details/k8s-virtual-service-details.component';
import { K8sVirtualServiceComponent } from './k8s-virtual-service/k8s-virtual-service.component';
import { K8sEndpointsComponent } from './k8s-endpoints/k8s-endpoints.component';
import { K8sEndpointsDetailsComponent } from './k8s-endpoints-details/k8s-endpoints-details.component';
import { K8sEndpointSliceComponent } from './k8s-endpoint-slice/k8s-endpoint-slice.component';
import { K } from '@angular/cdk/keycodes';
import { K8sEndpointSliceDetailsComponent } from './k8s-endpoint-slice-details/k8s-endpoint-slice-details.component';
import { K8sPodDisruptionBudgetsComponent } from './k8s-pod-disruption-budgets/k8s-pod-disruption-budgets.component';
import { K8sPodDisruptionBudgetsDetailsComponent } from './k8s-pod-disruption-budgets-details/k8s-pod-disruption-budgets-details.component';
import { K8sControllerRevisionComponent } from './k8s-controller-revision/k8s-controller-revision.component';
import { K8sControllerRevisionDetailsComponent } from './k8s-controller-revision-details/k8s-controller-revision-details.component';
import { K8sReplicationControllerComponent } from './k8s-replication-controller/k8s-replication-controller.component';
import { K8sReplicationControllerDetailsComponent } from './k8s-replication-controller-details/k8s-replication-controller-details.component';

const k8sNamespaceResourceRoutes: Routes = [
  {
    path: 'deployments',
    component: K8sDeploymentsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Deployments',
      toolbarTitle: 'Deployment',
      breadcrumb: 'Deployments',
      permissions: ['VIEW_NAMESPACE_DEPLOYMENT']
    }
  },
  {
    path: 'deployments/:name',
    component: K8sDeploymentsDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Deployments',
      toolbarTitle: 'Deployment',
      breadcrumb: 'Deployments',
      permissions: ['VIEW_NAMESPACE_DEPLOYMENT']
    }
  },
  {
    path: 'pods',
    component: K8sPodsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Pods',
      toolbarTitle: 'Pods',
      breadcrumb: 'Pods',
      permissions: ['VIEW_NAMESPACE_POD']
    }
  },
  {
    path: 'pods/:name',
    component: K8sPodsDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Pods',
      toolbarTitle: 'Pod',
      breadcrumb: 'Pods',
      permissions: ['VIEW_NAMESPACE_POD']
    }
  },
  {
    path: 'replica-sets',
    component: K8sReplicaSetsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | ReplicaSets',
      toolbarTitle: 'ReplicaSet',
      breadcrumb: 'ReplicaSets',
      refreshComponent: true,
      permissions: ['VIEW_NAMESPACE_REPLICA_SET']
    }
  },
  {
    path: 'replica-sets/:name',
    component: K8sReplicaSetsDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | ReplicaSets',
      toolbarTitle: 'ReplicaSet',
      breadcrumb: 'ReplicaSets',
      permissions: ['VIEW_NAMESPACE_REPLICA_SET']
    }
  },
  {
    path: 'stateful-sets',
    component: K8sStatefulSetsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | StatefulSets',
      toolbarTitle: 'StatefulSet',
      breadcrumb: 'StatefulSets',
      permissions: ['VIEW_NAMESPACE_STATEFUL_SET']
    }
  },
  {
    path: 'stateful-sets/:name',
    component: K8sStatefulSetsDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | StatefulSets',
      toolbarTitle: 'StatefulSet',
      breadcrumb: 'StatefulSets',
      permissions: ['VIEW_NAMESPACE_STATEFUL_SET']
    }
  },
  {
    path: 'daemon-sets',
    component: K8sDaemonSetsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | DaemonSets',
      toolbarTitle: 'DaemonSet',
      breadcrumb: 'DaemonSets',
      permissions: ['VIEW_NAMESPACE_DAEMON_SET']
    }
  },
  {
    path: 'daemon-sets/:name',
    component: K8sDaemonSetsDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | DaemonSets',
      toolbarTitle: 'DaemonSet',
      breadcrumb: 'DaemonSets',
      permissions: ['VIEW_NAMESPACE_DAEMON_SET']
    }
  },
  {
    path: 'gateway',
    component: K8sGatewayComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Gateway',
      toolbarTitle: 'Gateway',
      breadcrumb: 'Gateway',
      permissions: ['VIEW_NAMESPACE_GATEWAY']
    }
  },
  {
    path: 'gateway/:name',
    component: K8sGatewayDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Gateway',
      toolbarTitle: 'Gateway',
      breadcrumb: 'Gateway',
      permissions: ['VIEW_NAMESPACE_GATEWAY']
    }
  },
  {
    path: 'virtual-service',
    component: K8sVirtualServiceComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | VirtualService',
      toolbarTitle: 'VirtualService',
      breadcrumb: 'VirtualService',
      permissions: ['VIEW_NAMESPACE_VIRTUAL_SERVICE']
    }
  },
  {
    path: 'virtual-service/:name',
    component: K8sVirtualServiceDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | VirtualService',
      toolbarTitle: 'VirtualService',
      breadcrumb: 'VirtualService',
      permissions: ['VIEW_NAMESPACE_VIRTUAL_SERVICE']
    }
  },
  {
    path: 'config-maps',
    component: K8sConfigMapsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | ConfigMaps',
      toolbarTitle: 'ConfigMap',
      breadcrumb: 'ConfigMaps',
      permissions: ['VIEW_NAMESPACE_CONFIG_MAP']
    }
  },
  {
    path: 'config-maps/:name',
    component: K8sConfigMapsDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | ConfigMaps',
      toolbarTitle: 'ConfigMap',
      breadcrumb: 'ConfigMaps',
      permissions: ['VIEW_NAMESPACE_CONFIG_MAP']
    }
  },
  {
    path: 'pvcs',
    component: K8sPvcsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Persistent Volume Claims',
      toolbarTitle: 'Persistent Volume Claim',
      breadcrumb: 'Persistent Volume Claims',
      permissions: ['VIEW_NAMESPACE_PERSISTENT_VOLUME']
    }
  },
  {
    path: 'pvcs/:name',
    component: K8sPvcsDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Persistent Volume Claims',
      toolbarTitle: 'Persistent Volume Claim',
      breadcrumb: 'Persistent Volume Claims',
      permissions: ['VIEW_NAMESPACE_PERSISTENT_VOLUME']
    }
  },
  {
    path: 'ingresses',
    component: K8sIngressesComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Ingresses',
      toolbarTitle: 'Ingress',
      breadcrumb: 'Ingresses',
      permissions: ['VIEW_NAMESPACE_INGRESS']
    }
  },
  {
    path: 'ingresses/:name',
    component: K8sIngressesDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Ingresses',
      toolbarTitle: 'Ingress',
      breadcrumb: 'Ingresses',
      permissions: ['VIEW_NAMESPACE_INGRESS']
    }
  },
  {
    path: 'secrets',
    component: K8sSecretsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Secrets',
      toolbarTitle: 'Secret',
      breadcrumb: 'Secrets',
      permissions: ['VIEW_NAMESPACE_SECRET']
    }
  },
  {
    path: 'secrets/:name',
    component: K8sSecretsDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Secrets',
      toolbarTitle: 'Secret',
      breadcrumb: 'Secrets',
      permissions: ['VIEW_NAMESPACE_SECRET']
    }
  },
  {
    path: 'service',
    component: K8sServiceComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Service ',
      toolbarTitle: 'Service',
      breadcrumb: 'Service ',
      permissions: ['VIEW_NAMESPACE_SERVICE']
    }
  },
  {
    path: 'service/:name',
    component: K8sServiceDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Service',
      toolbarTitle: 'Service',
      breadcrumb: 'Service ',
      permissions: ['VIEW_NAMESPACE_SERVICE']
    }
  },
  {
    path: 'service-accounts',
    component: K8sServiceAccountsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Service Accounts',
      toolbarTitle: 'Service Account',
      breadcrumb: 'Service Accounts',
      permissions: ['VIEW_NAMESPACE_SERVICE_ACCOUNT']
    }
  },
  {
    path: 'service-accounts/:name',
    canActivate: [RoleGuardService],
    component: K8sServiceAccountsDetailsComponent,
    data: {
      title: 'Cluster | Namespace | Service Accounts',
      toolbarTitle: 'Service Account',
      breadcrumb: 'Service Accounts',
      permissions: ['VIEW_NAMESPACE_SERVICE_ACCOUNT']
    }
  },
  {
    path: 'certificates',
    component: K8sCertificatesComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Certificates',
      toolbarTitle: 'Certificate',
      breadcrumb: 'Certificates',
      permissions: ['VIEW_NAMESPACE_CERTIFICATE']
    }
  },
  {
    path: 'certificates/:name',
    component: K8sCertificatesDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Certificates',
      toolbarTitle: 'Certificate',
      breadcrumb: 'Certificates',
      permissions: ['VIEW_NAMESPACE_CERTIFICATE']
    }
  },
  {
    path: 'role',
    component: K8sRoleComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Role',
      toolbarTitle: 'Role',
      breadcrumb: 'Role',
      permissions: ['VIEW_NAMESPACE_ROLE']
    }
  },
  {
    path: 'role/:name',
    component: K8sRoleDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Role',
      toolbarTitle: 'Role',
      breadcrumb: 'Role',
      permissions: ['VIEW_NAMESPACE_ROLE']
    }
  },
  {
    path: 'role-binding',
    component: K8sRoleBindingComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Role Binding',
      toolbarTitle: 'Role Binding',
      breadcrumb: 'Role Binding',
      permissions: ['VIEW_NAMESPACE_ROLE_BINDING']
    }
  },
  {
    path: 'role-binding/:name',
    component: K8sRoleBindingDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Role Binding',
      toolbarTitle: 'Role Binding',
      breadcrumb: 'Role Binding',
      permissions: ['VIEW_NAMESPACE_ROLE_BINDING']
    }
  },
  {
    path: 'job',
    component: K8sJobComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Job',
      toolbarTitle: 'Job',
      breadcrumb: 'Job',
      permissions: ['VIEW_NAMESPACE_JOB']
    }
  },
  {
    path: 'job/:name',
    component: K8sJobDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Job',
      toolbarTitle: 'Job',
      breadcrumb: 'Job',
      permissions: ['VIEW_NAMESPACE_JOB']
    }
  },
  {
    path: 'cron-job',
    component: K8sCronJobComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Cron Job',
      toolbarTitle: 'Cron Job',
      breadcrumb: 'Cron Job',
      permissions: ['VIEW_NAMESPACE_CRON_JOB']
    }
  },
  {
    path: 'cron-job/:name',
    component: K8sCronJobDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Cron Job',
      toolbarTitle: 'Cron Job',
      breadcrumb: 'Cron Job',
      permissions: ['VIEW_NAMESPACE_CRON_JOB']
    }
  },
  {
    path: 'network-policy',
    component: K8sNetworkPolicyComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Network Policy',
      toolbarTitle: 'Network Policy',
      breadcrumb: 'Network Policy',
      permissions: ['VIEW_NAMESPACE_NETWORK_POLICY']
    }
  },
  {
    path: 'network-policy/:name',
    component: K8sNetworkPolicyDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Network Policy',
      toolbarTitle: 'Network Policy',
      breadcrumb: 'Network Policy',
      permissions: ['VIEW_NAMESPACE_NETWORK_POLICY']
    }
  },
  {
    path: 'resource-quota',
    component: K8sResourceQuotaComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Resource Quota',
      toolbarTitle: 'Resource Quota',
      breadcrumb: 'Resource Quota',
      permissions: ['VIEW_NAMESPACE_RESOURCE_QUOTA']
    }
  },
  {
    path: 'resource-quota/:name',
    component: K8sResourceQuotaDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace | Resource Quota',
      toolbarTitle: 'Resource Quota',
      breadcrumb: 'Resource Quota',
      permissions: ['VIEW_NAMESPACE_RESOURCE_QUOTA']
    }
  },
  {
    path: 'endpoints',
    component: K8sEndpointsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'KloverCloud | Cluster | Namespace | Endpoints',
      toolbarTitle: 'Endpoints',
      breadcrumb: 'Endpoints',
      permissions: ['VIEW_NAMESPACE_ENDPOINTS']
    }
  },
  {
    path: 'endpoints/:name',
    component: K8sEndpointsDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'KloverCloud | Cluster | Namespace | Endpoints',
      toolbarTitle: 'Endpoints',
      breadcrumb: 'Endpoints',
      permissions: ['VIEW_NAMESPACE_ENDPOINTS']
    }
  },
  {
    path: 'endpoints-slice',
    component: K8sEndpointSliceComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'KloverCloud | Cluster | Namespace | Endpoints Slice',
      toolbarTitle: 'Endpoints Slice',
      breadcrumb: 'Endpoints Slice',
      permissions: ['VIEW_NAMESPACE_ENDPOINT_SLICE']
    }
  },
  {
    path: 'endpoints-slice/:name',
    component: K8sEndpointSliceDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'KloverCloud | Cluster | Namespace | Endpoints Slice',
      toolbarTitle: 'Endpoints Slice',
      breadcrumb: 'Endpoints Slice',
      permissions: ['VIEW_NAMESPACE_ENDPOINT_SLICE']
    }
  },
  {
    path: 'PDB',
    component: K8sPodDisruptionBudgetsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'KloverCloud | Cluster | Namespace | Pod Disruption Budget',
      toolbarTitle: 'Pod Disruption Budget',
      breadcrumb: 'Pod Disruption Budget',
      permissions: ['VIEW_NAMESPACE_PDB']
    }
  },
  {
    path: 'PDB/:name',
    component: K8sPodDisruptionBudgetsDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'KloverCloud | Cluster | Namespace | Pod Disruption Budget',
      toolbarTitle: 'Pod Disruption Budget',
      breadcrumb: 'Pod Disruption Budget',
      permissions: ['VIEW_NAMESPACE_PDB']
    }
  },
  {
    path: 'controller-revision',
    component: K8sControllerRevisionComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'KloverCloud | Cluster | Namespace | Controller Revision',
      toolbarTitle: 'Controller Revision',
      breadcrumb: 'Controller Revision',
      permissions: ['VIEW_NAMESPACE_CONTROLLER_REVISION']
    }
  },
  {
    path: 'controller-revision/:name',
    component: K8sControllerRevisionDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'KloverCloud | Cluster | Namespace | Controller Revision',
      toolbarTitle: 'Controller Revision',
      breadcrumb: 'Controller Revision',
      permissions: ['VIEW_NAMESPACE_CONTROLLER_REVISION']
    }
  },
  {
    path: 'replication-controller',
    component: K8sReplicationControllerComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'KloverCloud | Cluster | Namespace | Replication Controller',
      toolbarTitle: 'Replication Controller',
      breadcrumb: 'Replication Controller',
      permissions: ['VIEW_NAMESPACE_REPLICATION_CONTROLLER']
    }
  },
  {
    path: 'replication-controller/:name',
    component: K8sReplicationControllerDetailsComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'KloverCloud | Cluster | Namespace | Replication Controller',
      toolbarTitle: 'Replication Controller',
      breadcrumb: 'Replication Controller',
      permissions: ['VIEW_NAMESPACE_REPLICATION_CONTROLLER']
    }
  }
];

const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    component: K8sNamespacesListComponent,
    canActivate: [RoleGuardService],
    data: {
      title: 'Cluster | Namespace List',
      toolbarTitle: 'Namespaces',
      breadcrumb: 'Namespaces',
      permissions: ['VIEW_K8S_NAMESPACE']
    }
  },
  {
    path: '',
    component: K8sNamespacesComponent,
    canActivate: [RoleGuardService],
    children: k8sNamespaceResourceRoutes
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class K8sNamespacesRoutingModule {}
