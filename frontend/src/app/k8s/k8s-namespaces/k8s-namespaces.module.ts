import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatDialogModule } from '@angular/material/dialog';
import { MatExpansionModule } from '@angular/material/expansion';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatMenuModule } from '@angular/material/menu';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatSelectModule } from '@angular/material/select';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatSortModule } from '@angular/material/sort';
import { MatTooltipModule } from '@angular/material/tooltip';
import { K8sUpdateModule } from '@k8s/k8s-update/k8s-update.module';
import { SharedModule } from '@shared-ui/shared.module';
import { IconModule } from '@visurel/iconify-angular';
import { Ng2SearchPipeModule } from 'ng2-search-filter';
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
import { K8sNamespacesDetailsComponent } from './k8s-namespaces-details/k8s-namespaces-details.component';
import { K8sNamespacesListComponent } from './k8s-namespaces-list/k8s-namespaces-list.component';
import { K8sNamespacesRoutingModule } from './k8s-namespaces-routing.module';
import { K8sNamespacesComponent } from './k8s-namespaces.component';
import { K8sNamespacesService } from './k8s-namespaces.service';
import { K8sNetworkPolicyDetailsComponent } from './k8s-network-policy-details/k8s-network-policy-details.component';
import { K8sNetworkPolicyComponent } from './k8s-network-policy/k8s-network-policy.component';
import { K8sPodsContainerLogComponent } from './k8s-pods-details/k8s-pods-container-log/k8s-pods-container-log.component';
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

import { CdkClipboardModule } from '@cdk-ui/clipboard';
import { CdkIconModule } from '@cdk-ui/icon';
import { ExpansionDataViewerTemplateComponent } from '@k8s/common/components/expansion-data-viewer-template/expansion-data-viewer-template.component';
import { JsonDataViewerTemplateComponent } from '@k8s/common/components/json-data-viewer-template/json-data-viewer-template.component';
import { K8sEventsTemplateComponent } from '@k8s/common/components/k8s-events-template/k8s-events-template.component';
import { MetadataTemplateComponent } from '@k8s/common/components/metadata-template/metadata-template.component';
import { AceEditorModule } from '@klovercloud/ace-editor/ace-editor.module';
import { K8sControllerRevisionDetailsComponent } from './k8s-controller-revision-details/k8s-controller-revision-details.component';
import { K8sControllerRevisionComponent } from './k8s-controller-revision/k8s-controller-revision.component';
import { K8sDeploymentPodListComponent } from './k8s-deployments-details/k8s-deployment-pod-list/k8s-deployment-pod-list.component';
import { K8sEndpointSliceDetailsComponent } from './k8s-endpoint-slice-details/k8s-endpoint-slice-details.component';
import { K8sEndpointSliceComponent } from './k8s-endpoint-slice/k8s-endpoint-slice.component';
import { K8sEndpointsDetailsComponent } from './k8s-endpoints-details/k8s-endpoints-details.component';
import { K8sEndpointsComponent } from './k8s-endpoints/k8s-endpoints.component';
import { K8sPodDisruptionBudgetsDetailsComponent } from './k8s-pod-disruption-budgets-details/k8s-pod-disruption-budgets-details.component';
import { K8sPodDisruptionBudgetsComponent } from './k8s-pod-disruption-budgets/k8s-pod-disruption-budgets.component';
import { GrafanaDashboardComponent } from './k8s-pods-details/grafana-dashboard/grafana-dashboard.component';
import { K8sReplicationControllerDetailsComponent } from './k8s-replication-controller-details/k8s-replication-controller-details.component';
import { K8sReplicationControllerComponent } from './k8s-replication-controller/k8s-replication-controller.component';
import { K8sStatefulsetPodListComponent } from './k8s-stateful-sets-details/k8s-statefulset-pod-list/k8s-statefulset-pod-list.component';
import { NgApexchartsModule } from 'ng-apexcharts';

@NgModule({
  declarations: [
    K8sNamespacesListComponent,
    K8sDeploymentsComponent,
    K8sNamespacesComponent,
    K8sPodsComponent,
    K8sReplicaSetsComponent,
    K8sConfigMapsComponent,
    K8sPvcsComponent,
    K8sIngressesComponent,
    K8sCertificatesComponent,
    K8sDeploymentsDetailsComponent,
    K8sReplicaSetsDetailsComponent,
    K8sStatefulSetsDetailsComponent,
    K8sStatefulSetsComponent,
    K8sPodsDetailsComponent,
    K8sServiceAccountsComponent,
    K8sServiceAccountsDetailsComponent,
    K8sSecretsComponent,
    K8sSecretsDetailsComponent,
    K8sConfigMapsDetailsComponent,
    K8sPvcsDetailsComponent,
    K8sIngressesDetailsComponent,
    K8sCertificatesDetailsComponent,
    K8sDaemonSetsComponent,
    K8sDaemonSetsDetailsComponent,
    K8sNamespacesDetailsComponent,
    K8sRoleBindingComponent,
    K8sRoleBindingDetailsComponent,
    K8sNetworkPolicyComponent,
    K8sNetworkPolicyDetailsComponent,
    K8sResourceQuotaComponent,
    K8sResourceQuotaDetailsComponent,
    K8sRoleComponent,
    K8sRoleDetailsComponent,
    K8sServiceComponent,
    K8sServiceDetailsComponent,
    K8sJobComponent,
    K8sJobDetailsComponent,
    K8sCronJobComponent,
    K8sCronJobDetailsComponent,
    K8sGatewayComponent,
    K8sVirtualServiceComponent,
    K8sVirtualServiceDetailsComponent,
    K8sGatewayDetailsComponent,
    K8sPodsContainerLogComponent,
    K8sDeploymentPodListComponent,
    K8sStatefulsetPodListComponent,
    K8sEndpointsComponent,
    K8sEndpointsDetailsComponent,
    K8sEndpointSliceComponent,
    K8sEndpointSliceDetailsComponent,
    K8sPodDisruptionBudgetsComponent,
    K8sPodDisruptionBudgetsDetailsComponent,
    K8sControllerRevisionComponent,
    K8sControllerRevisionDetailsComponent,
    K8sReplicationControllerComponent,
    K8sReplicationControllerDetailsComponent,
    GrafanaDashboardComponent
  ],
  imports: [
    CommonModule,
    JsonDataViewerTemplateComponent,
    MetadataTemplateComponent,
    K8sNamespacesRoutingModule,
    FormsModule,
    MatProgressSpinnerModule,
    MatIconModule,
    MatButtonModule,
    MatFormFieldModule,
    MatSelectModule,
    MatTooltipModule,
    MatSlideToggleModule,
    MatSortModule,
    MatMenuModule,
    MatDialogModule,
    MatExpansionModule,
    Ng2SearchPipeModule,
    SharedModule,
    IconModule,
    K8sUpdateModule,
    MatInputModule,
    ReactiveFormsModule,
    AceEditorModule,
    CdkClipboardModule,
    K8sEventsTemplateComponent,
    CdkIconModule,
    K8sEventsTemplateComponent,
    ExpansionDataViewerTemplateComponent,
    NgApexchartsModule
  ],
  providers: [K8sNamespacesService]
})
export class K8sNamespacesModule {}
