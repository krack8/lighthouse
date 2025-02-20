package tasks

import (
	k8s2 "github.com/krack8/lighthouse/pkg/common/k8s"
)

func InitTaskRegistry() {
	//namespace
	RegisterTask(k8s2.NamespaceService().GetNamespaceList, k8s2.GetNamespaceListInputParams{})
	RegisterTask(k8s2.NamespaceService().GetNamespaceNameList, k8s2.GetNamespaceNamesInputParams{})
	RegisterTask(k8s2.NamespaceService().GetNamespaceDetails, k8s2.GetNamespaceInputParams{})
	RegisterTask(k8s2.NamespaceService().DeployNamespace, k8s2.DeployNamespaceInputParams{})
	RegisterTask(k8s2.NamespaceService().DeleteNamespace, k8s2.DeleteNamespaceInputParams{})

	//certficate
	RegisterTask(k8s2.CertificateService().GetCertificateList, k8s2.GetCertificateListInputParams{})
	RegisterTask(k8s2.CertificateService().GetCertificateDetails, k8s2.GetCertificateDetailsInputParams{})
	RegisterTask(k8s2.CertificateService().DeployCertificate, k8s2.DeployCertificateInputParams{})
	RegisterTask(k8s2.CertificateService().DeleteCertificate, k8s2.DeleteCertificateInputParams{})

	//clusterRole
	RegisterTask(k8s2.ClusterRoleService().GetClusterRoleList, k8s2.GetClusterRoleListInputParams{})
	RegisterTask(k8s2.ClusterRoleService().GetClusterRoleDetails, k8s2.GetClusterRoleDetailsInputParams{})
	RegisterTask(k8s2.ClusterRoleService().DeployClusterRole, k8s2.DeployClusterRoleInputParams{})
	RegisterTask(k8s2.ClusterRoleService().DeleteClusterRole, k8s2.DeleteClusterRoleInputParams{})

	//clusterRoleBinding
	RegisterTask(k8s2.ClusterRoleBindingService().GetClusterRoleBindingList, k8s2.GetClusterRoleBindingListInputParams{})
	RegisterTask(k8s2.ClusterRoleBindingService().GetClusterRoleBindingDetails, k8s2.GetClusterRoleBindingDetailsInputParams{})
	RegisterTask(k8s2.ClusterRoleBindingService().DeployClusterRoleBinding, k8s2.DeployClusterRoleBindingInputParams{})
	RegisterTask(k8s2.ClusterRoleBindingService().DeleteClusterRoleBinding, k8s2.DeleteClusterRoleBindingInputParams{})

	//configMap
	RegisterTask(k8s2.ConfigMapService().GetConfigMapList, k8s2.GetConfigMapListInputParams{})
	RegisterTask(k8s2.ConfigMapService().GetConfigMapDetails, k8s2.GetConfigMapDetailsInputParams{})
	RegisterTask(k8s2.ConfigMapService().DeployConfigMap, k8s2.DeployConfigMapInputParams{})
	RegisterTask(k8s2.ConfigMapService().DeleteConfigMap, k8s2.DeleteConfigMapInputParams{})

	//controllerRevision
	RegisterTask(k8s2.ControllerRevisionService().GetControllerRevisionList, k8s2.GetControllerRevisionListInputParams{})
	RegisterTask(k8s2.ControllerRevisionService().GetControllerRevisionDetails, k8s2.GetControllerRevisionDetailsInputParams{})
	RegisterTask(k8s2.ControllerRevisionService().DeployControllerRevision, k8s2.DeployControllerRevisionInputParams{})
	RegisterTask(k8s2.ControllerRevisionService().DeleteControllerRevision, k8s2.DeleteControllerRevisionInputParams{})

	//CRD
	RegisterTask(k8s2.CrdService().GetCrdList, k8s2.GetCrdListInputParams{})
	RegisterTask(k8s2.CrdService().GetCrdDetails, k8s2.GetCrdDetailsInputParams{})
	RegisterTask(k8s2.CrdService().DeployCrd, k8s2.DeployCrdInputParams{})
	RegisterTask(k8s2.CrdService().DeleteCrd, k8s2.DeleteCrdInputParams{})

	//customResource
	RegisterTask(k8s2.CustomResourceService().GetCustomResourceList, k8s2.GetCustomResourceListInputParams{})
	RegisterTask(k8s2.CustomResourceService().GetCustomResourceDetails, k8s2.GetCustomResourceDetailsInputParams{})
	RegisterTask(k8s2.CustomResourceService().DeployCustomResource, k8s2.DeployCustomResourceInputParams{})
	RegisterTask(k8s2.CustomResourceService().DeleteCustomResource, k8s2.DeleteCustomResourceInputParams{})

	//cronJob
	RegisterTask(k8s2.CronJobService().GetCronJobList, k8s2.GetCronJobListInputParams{})
	RegisterTask(k8s2.CronJobService().GetCronJobDetails, k8s2.GetCronJobInputParams{})
	RegisterTask(k8s2.CronJobService().DeployCronJob, k8s2.DeployCronJobInputParams{})
	RegisterTask(k8s2.CronJobService().DeleteCronJob, k8s2.DeleteCronJobInputParams{})

	//daemonSet
	RegisterTask(k8s2.DaemonSetService().GetDaemonSetList, k8s2.GetDaemonSetListInputParams{})
	RegisterTask(k8s2.DaemonSetService().GetDaemonSetDetails, k8s2.GetDaemonSetDetailsInputParams{})
	RegisterTask(k8s2.DaemonSetService().GetDaemonSetStats, k8s2.GetDaemonSetStatsInputParams{})
	RegisterTask(k8s2.DaemonSetService().DeployDaemonSet, k8s2.DeployDaemonSetInputParams{})
	RegisterTask(k8s2.DaemonSetService().DeleteDaemonSet, k8s2.DeleteDaemonSetInputParams{})

	//deployment
	RegisterTask(k8s2.DeploymentService().GetDeploymentList, k8s2.GetDeploymentListInputParams{})
	RegisterTask(k8s2.DeploymentService().GetDeploymentDetails, k8s2.GetDeploymentDetailsInputParams{})
	RegisterTask(k8s2.DeploymentService().GetDeploymentStats, k8s2.GetDeploymentStatsInputParams{})
	RegisterTask(k8s2.DeploymentService().GetDeploymentPodList, k8s2.GetDeploymentPodListInputParams{})
	RegisterTask(k8s2.DeploymentService().DeployDeployment, k8s2.DeployDeploymentInputParams{})
	RegisterTask(k8s2.DeploymentService().DeleteDeployment, k8s2.DeleteDeploymentInputParams{})

	//endpoints
	RegisterTask(k8s2.EndpointsService().GetEndpointsList, k8s2.GetEndpointsListInputParams{})
	RegisterTask(k8s2.EndpointsService().GetEndpointsDetails, k8s2.GetEndpointsDetailsInputParams{})
	RegisterTask(k8s2.EndpointsService().DeployEndpoints, k8s2.DeployEndpointsInputParams{})
	RegisterTask(k8s2.EndpointsService().DeleteEndpoints, k8s2.DeleteEndpointsInputParams{})

	//endpointSlice
	RegisterTask(k8s2.EndpointSliceService().GetEndpointSliceList, k8s2.GetEndpointSliceListInputParams{})
	RegisterTask(k8s2.EndpointSliceService().GetEndpointSliceDetails, k8s2.GetEndpointSliceDetailsInputParams{})
	RegisterTask(k8s2.EndpointSliceService().DeployEndpointSlice, k8s2.DeployEndpointSliceInputParams{})
	RegisterTask(k8s2.EndpointSliceService().DeleteEndpointSlice, k8s2.DeleteEndpointSliceInputParams{})

	//event
	RegisterTask(k8s2.EventService().GetEventList, k8s2.GetEventListInputParams{})
	RegisterTask(k8s2.EventService().GetEventDetails, k8s2.GetEventDetailsInputParams{})

	//hpa
	RegisterTask(k8s2.HpaService().GetHpaList, k8s2.GetHpaListInputParams{})
	RegisterTask(k8s2.HpaService().GetHpaDetails, k8s2.GetHpaDetailsInputParams{})

	//ingress
	RegisterTask(k8s2.IngressService().GetIngressList, k8s2.GetIngressListInputParams{})
	RegisterTask(k8s2.IngressService().GetIngressDetails, k8s2.GetIngressDetailsInputParams{})
	RegisterTask(k8s2.IngressService().DeployIngress, k8s2.DeployIngressInputParams{})
	RegisterTask(k8s2.IngressService().DeleteIngress, k8s2.DeleteIngressInputParams{})

	//istioGateway
	RegisterTask(k8s2.IstioGatewayService().GetIstioGatewayList, k8s2.GetIstioGatewayListInputParams{})
	RegisterTask(k8s2.IstioGatewayService().GetIstioGatewayDetails, k8s2.GetIstioGatewayDetailsInputParams{})
	RegisterTask(k8s2.IstioGatewayService().DeployIstioGateway, k8s2.DeployIstioGatewayInputParams{})
	RegisterTask(k8s2.IstioGatewayService().DeleteIstioGateway, k8s2.DeleteIstioGatewayInputParams{})

	//job
	RegisterTask(k8s2.JobService().GetJobList, k8s2.GetJobListInputParams{})
	RegisterTask(k8s2.JobService().GetJobDetails, k8s2.GetJobInputParams{})
	RegisterTask(k8s2.JobService().DeployJob, k8s2.DeployJobInputParams{})
	RegisterTask(k8s2.JobService().DeleteJob, k8s2.DeleteJobInputParams{})

	//loadBalancer
	RegisterTask(k8s2.LoadBalancerService().GetLoadBalancerList, k8s2.GetLoadBalancerListInputParams{})
	RegisterTask(k8s2.LoadBalancerService().GetLoadBalancerDetails, k8s2.GetLoadBalancerDetailsInputParams{})

	//Manifest
	RegisterTask(k8s2.ManifestService().DeployManifest, k8s2.DeployManifestInputParams{})

	//networkPolicy
	RegisterTask(k8s2.NetworkPolicyService().GetNetworkPolicyList, k8s2.GetNetworkPolicyListInputParams{})
	RegisterTask(k8s2.NetworkPolicyService().GetNetworkPolicyDetails, k8s2.GetNetworkPolicyDetailsInputParams{})

	//node
	RegisterTask(k8s2.NodeService().GetNodeList, k8s2.GetNodeListInputParams{})
	RegisterTask(k8s2.NodeService().GetNodeDetails, k8s2.GetNodeInputParams{})
	RegisterTask(k8s2.NodeService().NodeCordon, k8s2.NodeCordonInputParams{})
	RegisterTask(k8s2.NodeService().NodeTaint, k8s2.NodeTaintInputParams{})
	RegisterTask(k8s2.NodeService().NodeUnTaint, k8s2.NodeUnTaintInputParams{})

	//pod
	RegisterTask(k8s2.PodService().GetPodList, k8s2.GetPodListInputParams{})
	RegisterTask(k8s2.PodService().GetPodDetails, k8s2.GetPodDetailsInputParams{})
	RegisterTask(k8s2.PodService().GetPodLogs, k8s2.GetPodLogsInputParams{})
	RegisterTask(k8s2.PodService().GetPodStats, k8s2.GetPodStatsInputParams{})
	RegisterTask(k8s2.PodService().DeployPod, k8s2.DeployPodInputParams{})
	RegisterTask(k8s2.PodService().DeletePod, k8s2.DeletePodInputParams{})

	//podDisruptionBudget
	RegisterTask(k8s2.PodDisruptionBudgetsService().GetPodDisruptionBudgetsList, k8s2.GetPodDisruptionBudgetsListInputParams{})
	RegisterTask(k8s2.PodDisruptionBudgetsService().GetPodDisruptionBudgetsDetails, k8s2.GetPodDisruptionBudgetsDetailsInputParams{})
	RegisterTask(k8s2.PodDisruptionBudgetsService().DeployPodDisruptionBudgets, k8s2.DeployPodDisruptionBudgetsInputParams{})
	RegisterTask(k8s2.PodDisruptionBudgetsService().DeletePodDisruptionBudgets, k8s2.DeletePodDisruptionBudgetsInputParams{})

	//podMetrics
	RegisterTask(k8s2.PodMetricsService().GetPodMetricsList, k8s2.GetPodMetricsListInputParams{})
	RegisterTask(k8s2.PodMetricsService().GetPodMetricsDetails, k8s2.GetPodMetricsDetailsInputParams{})

	//pv
	RegisterTask(k8s2.PvService().GetPvList, k8s2.GetPvListInputParams{})
	RegisterTask(k8s2.PvService().GetPvDetails, k8s2.GetPvDetailsInputParams{})
	RegisterTask(k8s2.PvService().DeployPv, k8s2.DeployPvInputParams{})
	RegisterTask(k8s2.PvService().DeletePv, k8s2.DeletePvInputParams{})

	//pvc
	RegisterTask(k8s2.PvcService().GetPvcList, k8s2.GetPvcListInputParams{})
	RegisterTask(k8s2.PvcService().GetPvcDetails, k8s2.GetPvcDetailsInputParams{})
	RegisterTask(k8s2.PvcService().DeployPvc, k8s2.DeployPvcInputParams{})
	RegisterTask(k8s2.PvcService().DeletePvc, k8s2.DeletePvcInputParams{})

	//replicaSet
	RegisterTask(k8s2.ReplicaSetService().GetReplicaSetList, k8s2.GetReplicaSetListInputParams{})
	RegisterTask(k8s2.ReplicaSetService().GetReplicaSetDetails, k8s2.GetReplicaSetDetailsInputParams{})
	RegisterTask(k8s2.ReplicaSetService().DeployReplicaSet, k8s2.DeployReplicaSetInputParams{})
	RegisterTask(k8s2.ReplicaSetService().DeleteReplicaSet, k8s2.DeleteReplicaSetInputParams{})

	//replicationController
	RegisterTask(k8s2.ReplicationControllerService().GetReplicationControllerList, k8s2.GetReplicationControllerListInputParams{})
	RegisterTask(k8s2.ReplicationControllerService().GetReplicationControllerDetails, k8s2.GetReplicationControllerDetailsInputParams{})
	RegisterTask(k8s2.ReplicationControllerService().DeployReplicationController, k8s2.DeployReplicationControllerInputParams{})
	RegisterTask(k8s2.ReplicationControllerService().DeleteReplicationController, k8s2.DeleteReplicationControllerInputParams{})

	//resourceQuota
	RegisterTask(k8s2.ResourceQuotaService().GetResourceQuotaList, k8s2.GetResourceQuotaListInputParams{})
	RegisterTask(k8s2.ResourceQuotaService().GetResourceQuotaDetails, k8s2.GetResourceQuotaDetailsInputParams{})
	RegisterTask(k8s2.ResourceQuotaService().DeployResourceQuota, k8s2.DeployResourceQuotaInputParams{})
	RegisterTask(k8s2.ResourceQuotaService().DeleteResourceQuota, k8s2.DeleteResourceQuotaInputParams{})

	//role
	RegisterTask(k8s2.RoleService().GetRoleList, k8s2.GetRoleListInputParams{})
	RegisterTask(k8s2.RoleService().GetRoleDetails, k8s2.GetRoleDetailsInputParams{})
	RegisterTask(k8s2.RoleService().DeployRole, k8s2.DeployRoleInputParams{})
	RegisterTask(k8s2.RoleService().DeleteRole, k8s2.DeleteRoleInputParams{})

	//roleBinding
	RegisterTask(k8s2.RoleBindingService().GetRoleBindingList, k8s2.GetRoleBindingListInputParams{})
	RegisterTask(k8s2.RoleBindingService().GetRoleBindingDetails, k8s2.GetRoleBindingDetailsInputParams{})
	RegisterTask(k8s2.RoleBindingService().DeployRoleBinding, k8s2.DeployRoleBindingInputParams{})
	RegisterTask(k8s2.RoleBindingService().DeleteRoleBinding, k8s2.DeleteRoleBindingInputParams{})

	//serviceAccount
	RegisterTask(k8s2.ServiceAccountService().GetServiceAccountList, k8s2.GetServiceAccountListInputParams{})
	RegisterTask(k8s2.ServiceAccountService().GetServiceAccountDetails, k8s2.GetServiceAccountDetailsInputParams{})
	RegisterTask(k8s2.ServiceAccountService().DeployServiceAccount, k8s2.DeployServiceAccountInputParams{})
	RegisterTask(k8s2.ServiceAccountService().DeleteServiceAccount, k8s2.DeleteServiceAccountInputParams{})

	//secret
	RegisterTask(k8s2.SecretService().GetSecretList, k8s2.GetSecretListInputParams{})
	RegisterTask(k8s2.SecretService().GetSecretDetails, k8s2.GetSecretDetailsInputParams{})
	RegisterTask(k8s2.SecretService().DeploySecret, k8s2.DeploySecretInputParams{})
	RegisterTask(k8s2.SecretService().DeleteSecret, k8s2.DeleteSecretInputParams{})

	//statefulSet
	RegisterTask(k8s2.StatefulSetService().GetStatefulSetList, k8s2.GetStatefulSetListInputParams{})
	RegisterTask(k8s2.StatefulSetService().GetStatefulSetDetails, k8s2.GetStatefulSetDetailsInputParams{})
	RegisterTask(k8s2.StatefulSetService().GetStatefulSetStats, k8s2.GetStatefulSetStatsInputParams{})
	RegisterTask(k8s2.StatefulSetService().GetStatefulSetPodList, k8s2.GetStatefulSetPodListInputParams{})
	RegisterTask(k8s2.StatefulSetService().DeployStatefulSet, k8s2.DeployStatefulSetInputParams{})
	RegisterTask(k8s2.StatefulSetService().DeleteStatefulSet, k8s2.DeleteStatefulSetInputParams{})

	//storageClass
	RegisterTask(k8s2.StorageClassService().GetStorageClassList, k8s2.GetStorageClassListInputParams{})
	RegisterTask(k8s2.StorageClassService().GetStorageClassDetails, k8s2.GetStorageClassDetailsInputParams{})
	RegisterTask(k8s2.StorageClassService().DeployStorageClass, k8s2.DeployStorageClassInputParams{})
	RegisterTask(k8s2.StorageClassService().DeleteStorageClass, k8s2.DeleteStorageClassInputParams{})

	//svc
	RegisterTask(k8s2.SvcService().GetSvcList, k8s2.GetSvcListInputParams{})
	RegisterTask(k8s2.SvcService().GetSvcDetails, k8s2.GetSvcDetailsInputParams{})
	RegisterTask(k8s2.SvcService().DeploySvc, k8s2.DeploySvcInputParams{})
	RegisterTask(k8s2.SvcService().DeleteSvc, k8s2.DeleteSvcInputParams{})

	//virtualService
	RegisterTask(k8s2.VirtualServiceService().GetVirtualServiceList, k8s2.GetVirtualServiceListInputParams{})
	RegisterTask(k8s2.VirtualServiceService().GetVirtualServiceDetails, k8s2.GetVirtualServiceDetailsInputParams{})
	RegisterTask(k8s2.VirtualServiceService().DeployVirtualService, k8s2.DeployVirtualServiceInputParams{})
	RegisterTask(k8s2.VirtualServiceService().DeleteVirtualService, k8s2.DeleteVirtualServiceInputParams{})

	//volumeSnapshot
	RegisterTask(k8s2.VolumeSnapshotService().GetVolumeSnapshotList, k8s2.GetVolumeSnapshotListInputParams{})
	RegisterTask(k8s2.VolumeSnapshotService().GetVolumeSnapshotDetails, k8s2.GetVolumeSnapshotDetailsInputParams{})
	RegisterTask(k8s2.VolumeSnapshotService().DeployVolumeSnapshot, k8s2.DeployVolumeSnapshotInputParams{})
	RegisterTask(k8s2.VolumeSnapshotService().DeleteVolumeSnapshot, k8s2.DeleteVolumeSnapshotInputParams{})

	//volumeSnapshotClass
	RegisterTask(k8s2.VolumeSnapshotClassService().GetVolumeSnapshotClassList, k8s2.GetVolumeSnapshotClassListInputParams{})
	RegisterTask(k8s2.VolumeSnapshotClassService().GetVolumeSnapshotClassDetails, k8s2.GetVolumeSnapshotClassDetailsInputParams{})

	//volumeSnapshotContent
	RegisterTask(k8s2.VolumeSnapshotContentService().GetVolumeSnapshotContentList, k8s2.GetVolumeSnapshotContentListInputParams{})
	RegisterTask(k8s2.VolumeSnapshotContentService().GetVolumeSnapshotContentDetails, k8s2.GetVolumeSnapshotContentDetailsInputParams{})
}
