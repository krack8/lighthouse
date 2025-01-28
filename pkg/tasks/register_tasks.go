package tasks

import "github.com/krack8/lighthouse/pkg/k8s"

func InitTaskRegistry() {
	//namespace
	RegisterTask(k8s.NamespaceService().GetNamespaceList, k8s.GetNamespaceListInputParams{})
	RegisterTask(k8s.NamespaceService().GetNamespaceNameList, k8s.GetNamespaceNamesInputParams{})
	RegisterTask(k8s.NamespaceService().GetNamespaceDetails, k8s.GetNamespaceInputParams{})
	RegisterTask(k8s.NamespaceService().DeployNamespace, k8s.DeployNamespaceInputParams{})
	RegisterTask(k8s.NamespaceService().DeleteNamespace, k8s.DeleteNamespaceInputParams{})

	//certficate
	RegisterTask(k8s.CertificateService().GetCertificateList, k8s.GetCertificateListInputParams{})
	RegisterTask(k8s.CertificateService().GetCertificateDetails, k8s.GetCertificateDetailsInputParams{})
	RegisterTask(k8s.CertificateService().DeployCertificate, k8s.DeployCertificateInputParams{})
	RegisterTask(k8s.CertificateService().DeleteCertificate, k8s.DeleteCertificateInputParams{})

	//clusterRole
	RegisterTask(k8s.ClusterRoleService().GetClusterRoleList, k8s.GetClusterRoleListInputParams{})
	RegisterTask(k8s.ClusterRoleService().GetClusterRoleDetails, k8s.GetClusterRoleDetailsInputParams{})
	RegisterTask(k8s.ClusterRoleService().DeployClusterRole, k8s.DeployClusterRoleInputParams{})
	RegisterTask(k8s.ClusterRoleService().DeleteClusterRole, k8s.DeleteClusterRoleInputParams{})

	//clusterRoleBinding
	RegisterTask(k8s.ClusterRoleBindingService().GetClusterRoleBindingList, k8s.GetClusterRoleBindingListInputParams{})
	RegisterTask(k8s.ClusterRoleBindingService().GetClusterRoleBindingDetails, k8s.GetClusterRoleBindingDetailsInputParams{})
	RegisterTask(k8s.ClusterRoleBindingService().DeployClusterRoleBinding, k8s.DeployClusterRoleBindingInputParams{})
	RegisterTask(k8s.ClusterRoleBindingService().DeleteClusterRoleBinding, k8s.DeleteClusterRoleBindingInputParams{})

	//configMap
	RegisterTask(k8s.ConfigMapService().GetConfigMapList, k8s.GetConfigMapListInputParams{})
	RegisterTask(k8s.ConfigMapService().GetConfigMapDetails, k8s.GetConfigMapDetailsInputParams{})
	RegisterTask(k8s.ConfigMapService().DeployConfigMap, k8s.DeployConfigMapInputParams{})
	RegisterTask(k8s.ConfigMapService().DeleteConfigMap, k8s.DeleteConfigMapInputParams{})

	//controllerRevision
	RegisterTask(k8s.ControllerRevisionService().GetControllerRevisionList, k8s.GetControllerRevisionListInputParams{})
	RegisterTask(k8s.ControllerRevisionService().GetControllerRevisionDetails, k8s.GetControllerRevisionDetailsInputParams{})
	RegisterTask(k8s.ControllerRevisionService().DeployControllerRevision, k8s.DeployControllerRevisionInputParams{})
	RegisterTask(k8s.ControllerRevisionService().DeleteControllerRevision, k8s.DeleteControllerRevisionInputParams{})

	//CRD
	RegisterTask(k8s.CrdService().GetCrdList, k8s.GetCrdListInputParams{})
	RegisterTask(k8s.CrdService().GetCrdDetails, k8s.GetCrdDetailsInputParams{})
	RegisterTask(k8s.CrdService().DeployCrd, k8s.DeployCrdInputParams{})
	RegisterTask(k8s.CrdService().DeleteCrd, k8s.DeleteCrdInputParams{})

	//customResource
	RegisterTask(k8s.CustomResourceService().GetCustomResourceList, k8s.GetCustomResourceListInputParams{})
	RegisterTask(k8s.CustomResourceService().GetCustomResourceDetails, k8s.GetCustomResourceDetailsInputParams{})
	RegisterTask(k8s.CustomResourceService().DeployCustomResource, k8s.DeployCustomResourceInputParams{})
	RegisterTask(k8s.CustomResourceService().DeleteCustomResource, k8s.DeleteCustomResourceInputParams{})

	//cronJob
	RegisterTask(k8s.CronJobService().GetCronJobList, k8s.GetCronJobListInputParams{})
	RegisterTask(k8s.CronJobService().GetCronJobDetails, k8s.GetCronJobInputParams{})
	RegisterTask(k8s.CronJobService().DeployCronJob, k8s.DeployCronJobInputParams{})
	RegisterTask(k8s.CronJobService().DeleteCronJob, k8s.DeleteCronJobInputParams{})

	//daemonSet
	RegisterTask(k8s.DaemonSetService().GetDaemonSetList, k8s.GetDaemonSetListInputParams{})
	RegisterTask(k8s.DaemonSetService().GetDaemonSetDetails, k8s.GetDaemonSetDetailsInputParams{})
	RegisterTask(k8s.DaemonSetService().GetDaemonSetStats, k8s.GetDaemonSetStatsInputParams{})
	RegisterTask(k8s.DaemonSetService().DeployDaemonSet, k8s.DeployDaemonSetInputParams{})
	RegisterTask(k8s.DaemonSetService().DeleteDaemonSet, k8s.DeleteDaemonSetInputParams{})

	//deployment
	RegisterTask(k8s.DeploymentService().GetDeploymentList, k8s.GetDeploymentListInputParams{})
	RegisterTask(k8s.DeploymentService().GetDeploymentDetails, k8s.GetDeploymentDetailsInputParams{})
	RegisterTask(k8s.DeploymentService().GetDeploymentStats, k8s.GetDeploymentStatsInputParams{})
	RegisterTask(k8s.DeploymentService().GetDeploymentPodList, k8s.GetDeploymentPodListInputParams{})
	RegisterTask(k8s.DeploymentService().DeployDeployment, k8s.DeployDeploymentInputParams{})
	RegisterTask(k8s.DeploymentService().DeleteDeployment, k8s.DeleteDeploymentInputParams{})

	//endpoints
	RegisterTask(k8s.EndpointsService().GetEndpointsList, k8s.GetEndpointsListInputParams{})
	RegisterTask(k8s.EndpointsService().GetEndpointsDetails, k8s.GetEndpointsDetailsInputParams{})
	RegisterTask(k8s.EndpointsService().DeployEndpoints, k8s.DeployEndpointsInputParams{})
	RegisterTask(k8s.EndpointsService().DeleteEndpoints, k8s.DeleteEndpointsInputParams{})

	//endpointSlice
	RegisterTask(k8s.EndpointSliceService().GetEndpointSliceList, k8s.GetEndpointSliceListInputParams{})
	RegisterTask(k8s.EndpointSliceService().GetEndpointSliceDetails, k8s.GetEndpointSliceDetailsInputParams{})
	RegisterTask(k8s.EndpointSliceService().DeployEndpointSlice, k8s.DeployEndpointSliceInputParams{})
	RegisterTask(k8s.EndpointSliceService().DeleteEndpointSlice, k8s.DeleteEndpointSliceInputParams{})

	//event
	RegisterTask(k8s.EventService().GetEventList, k8s.GetEventListInputParams{})
	RegisterTask(k8s.EventService().GetEventDetails, k8s.GetEventDetailsInputParams{})

	//hpa
	RegisterTask(k8s.HpaService().GetHpaList, k8s.GetHpaListInputParams{})
	RegisterTask(k8s.HpaService().GetHpaDetails, k8s.GetHpaDetailsInputParams{})

	//ingress
	RegisterTask(k8s.IngressService().GetIngressList, k8s.GetIngressListInputParams{})
	RegisterTask(k8s.IngressService().GetIngressDetails, k8s.GetIngressDetailsInputParams{})
	RegisterTask(k8s.IngressService().DeployIngress, k8s.DeployIngressInputParams{})
	RegisterTask(k8s.IngressService().DeleteIngress, k8s.DeleteIngressInputParams{})

	//istioGateway
	RegisterTask(k8s.IstioGatewayService().GetIstioGatewayList, k8s.GetIstioGatewayListInputParams{})
	RegisterTask(k8s.IstioGatewayService().GetIstioGatewayDetails, k8s.GetIstioGatewayDetailsInputParams{})
	RegisterTask(k8s.IstioGatewayService().DeployIstioGateway, k8s.DeployIstioGatewayInputParams{})
	RegisterTask(k8s.IstioGatewayService().DeleteIstioGateway, k8s.DeleteIstioGatewayInputParams{})

	//job
	RegisterTask(k8s.JobService().GetJobList, k8s.GetJobListInputParams{})
	RegisterTask(k8s.JobService().GetJobDetails, k8s.GetJobInputParams{})
	RegisterTask(k8s.JobService().DeployJob, k8s.DeployJobInputParams{})
	RegisterTask(k8s.JobService().DeleteJob, k8s.DeleteJobInputParams{})

	//loadBalancer
	RegisterTask(k8s.LoadBalancerService().GetLoadBalancerList, k8s.GetLoadBalancerListInputParams{})
	RegisterTask(k8s.LoadBalancerService().GetLoadBalancerDetails, k8s.GetLoadBalancerDetailsInputParams{})

	//Manifest
	RegisterTask(k8s.ManifestService().DeployManifest, k8s.DeployManifestInputParams{})

	//networkPolicy
	RegisterTask(k8s.NetworkPolicyService().GetNetworkPolicyList, k8s.GetNetworkPolicyListInputParams{})
	RegisterTask(k8s.NetworkPolicyService().GetNetworkPolicyDetails, k8s.GetNetworkPolicyDetailsInputParams{})

	//node
	RegisterTask(k8s.NodeService().GetNodeList, k8s.GetNodeListInputParams{})
	RegisterTask(k8s.NodeService().GetNodeDetails, k8s.GetNodeInputParams{})
	RegisterTask(k8s.NodeService().NodeCordon, k8s.NodeCordonInputParams{})
	RegisterTask(k8s.NodeService().NodeTaint, k8s.NodeTaintInputParams{})
	RegisterTask(k8s.NodeService().NodeUnTaint, k8s.NodeUnTaintInputParams{})

	//pod
	RegisterTask(k8s.PodService().GetPodList, k8s.GetPodListInputParams{})
	RegisterTask(k8s.PodService().GetPodDetails, k8s.GetPodDetailsInputParams{})
	RegisterTask(k8s.PodService().GetPodLogs, k8s.GetPodLogsInputParams{})
	RegisterTask(k8s.PodService().GetPodStats, k8s.GetPodStatsInputParams{})
	RegisterTask(k8s.PodService().DeployPod, k8s.DeployPodInputParams{})
	RegisterTask(k8s.PodService().DeletePod, k8s.DeletePodInputParams{})

	//podDisruptionBudget
	RegisterTask(k8s.PodDisruptionBudgetsService().GetPodDisruptionBudgetsList, k8s.GetPodDisruptionBudgetsListInputParams{})
	RegisterTask(k8s.PodDisruptionBudgetsService().GetPodDisruptionBudgetsDetails, k8s.GetPodDisruptionBudgetsDetailsInputParams{})
	RegisterTask(k8s.PodDisruptionBudgetsService().DeployPodDisruptionBudgets, k8s.DeployPodDisruptionBudgetsInputParams{})
	RegisterTask(k8s.PodDisruptionBudgetsService().DeletePodDisruptionBudgets, k8s.DeletePodDisruptionBudgetsInputParams{})

	//podMetrics
	RegisterTask(k8s.PodMetricsService().GetPodMetricsList, k8s.GetPodMetricsListInputParams{})
	RegisterTask(k8s.PodMetricsService().GetPodMetricsDetails, k8s.GetPodMetricsDetailsInputParams{})

	//pv
	RegisterTask(k8s.PvService().GetPvList, k8s.GetPvListInputParams{})
	RegisterTask(k8s.PvService().GetPvDetails, k8s.GetPvDetailsInputParams{})
	RegisterTask(k8s.PvService().DeployPv, k8s.DeployPvInputParams{})
	RegisterTask(k8s.PvService().DeletePv, k8s.DeletePvInputParams{})

	//pvc
	RegisterTask(k8s.PvcService().GetPvcList, k8s.GetPvcListInputParams{})
	RegisterTask(k8s.PvcService().GetPvcDetails, k8s.GetPvcDetailsInputParams{})
	RegisterTask(k8s.PvcService().DeployPvc, k8s.DeployPvcInputParams{})
	RegisterTask(k8s.PvcService().DeletePvc, k8s.DeletePvcInputParams{})

	//replicaSet
	RegisterTask(k8s.ReplicaSetService().GetReplicaSetList, k8s.GetReplicaSetListInputParams{})
	RegisterTask(k8s.ReplicaSetService().GetReplicaSetDetails, k8s.GetReplicaSetDetailsInputParams{})
	RegisterTask(k8s.ReplicaSetService().GetReplicaSetStats, k8s.GetReplicaSetStatsInputParams{})
	RegisterTask(k8s.ReplicaSetService().DeployReplicaSet, k8s.DeployReplicaSetInputParams{})
	RegisterTask(k8s.ReplicaSetService().DeleteReplicaSet, k8s.DeleteReplicaSetInputParams{})

	//replicationController
	RegisterTask(k8s.ReplicationControllerService().GetReplicationControllerList, k8s.GetReplicationControllerListInputParams{})
	RegisterTask(k8s.ReplicationControllerService().GetReplicationControllerList, k8s.GetReplicationControllerDetailsInputParams{})
	RegisterTask(k8s.ReplicationControllerService().DeployReplicationController, k8s.DeployReplicationControllerInputParams{})
	RegisterTask(k8s.ReplicationControllerService().DeleteReplicationController, k8s.DeleteReplicationControllerInputParams{})

	//resourceQuota
	RegisterTask(k8s.ResourceQuotaService().GetResourceQuotaList, k8s.GetResourceQuotaListInputParams{})
	RegisterTask(k8s.ResourceQuotaService().GetResourceQuotaDetails, k8s.GetResourceQuotaDetailsInputParams{})
	RegisterTask(k8s.ResourceQuotaService().DeployResourceQuota, k8s.DeployResourceQuotaInputParams{})
	RegisterTask(k8s.ResourceQuotaService().DeleteResourceQuota, k8s.DeleteResourceQuotaInputParams{})

	//role
	RegisterTask(k8s.RoleService().GetRoleList, k8s.GetRoleListInputParams{})
	RegisterTask(k8s.RoleService().GetRoleDetails, k8s.GetRoleDetailsInputParams{})
	RegisterTask(k8s.RoleService().DeployRole, k8s.DeployRoleInputParams{})
	RegisterTask(k8s.RoleService().DeleteRole, k8s.DeleteRoleInputParams{})

	//roleBinding
	RegisterTask(k8s.RoleBindingService().GetRoleBindingList, k8s.GetRoleBindingListInputParams{})
	RegisterTask(k8s.RoleBindingService().GetRoleBindingDetails, k8s.GetRoleBindingDetailsInputParams{})
	RegisterTask(k8s.RoleBindingService().DeployRoleBinding, k8s.DeployRoleBindingInputParams{})
	RegisterTask(k8s.RoleBindingService().DeleteRoleBinding, k8s.DeleteRoleBindingInputParams{})

	//serviceAccount
	RegisterTask(k8s.ServiceAccountService().GetServiceAccountList, k8s.GetServiceAccountListInputParams{})
	RegisterTask(k8s.ServiceAccountService().GetServiceAccountDetails, k8s.GetServiceAccountDetailsInputParams{})
	RegisterTask(k8s.ServiceAccountService().DeployServiceAccount, k8s.DeployServiceAccountInputParams{})
	RegisterTask(k8s.ServiceAccountService().DeleteServiceAccount, k8s.DeleteServiceAccountInputParams{})

	//secret
	RegisterTask(k8s.SecretService().GetSecretList, k8s.GetSecretListInputParams{})
	RegisterTask(k8s.SecretService().GetSecretDetails, k8s.GetSecretDetailsInputParams{})
	RegisterTask(k8s.SecretService().DeploySecret, k8s.DeploySecretInputParams{})
	RegisterTask(k8s.SecretService().DeleteSecret, k8s.DeleteSecretInputParams{})

	//statefulSet
	RegisterTask(k8s.StatefulSetService().GetStatefulSetList, k8s.GetStatefulSetListInputParams{})
	RegisterTask(k8s.StatefulSetService().GetStatefulSetDetails, k8s.GetStatefulSetDetailsInputParams{})
	RegisterTask(k8s.StatefulSetService().GetStatefulSetStats, k8s.GetStatefulSetStatsInputParams{})
	RegisterTask(k8s.StatefulSetService().GetStatefulSetPodList, k8s.GetStatefulSetPodListInputParams{})
	RegisterTask(k8s.StatefulSetService().DeployStatefulSet, k8s.DeployStatefulSetInputParams{})
	RegisterTask(k8s.StatefulSetService().DeleteStatefulSet, k8s.DeleteStatefulSetInputParams{})

	//storageClass
	RegisterTask(k8s.StorageClassService().GetStorageClassList, k8s.GetStorageClassListInputParams{})
	RegisterTask(k8s.StorageClassService().GetStorageClassDetails, k8s.GetStorageClassDetailsInputParams{})
	RegisterTask(k8s.StorageClassService().DeployStorageClass, k8s.DeployStorageClassInputParams{})
	RegisterTask(k8s.StorageClassService().DeleteStorageClass, k8s.DeleteStorageClassInputParams{})

	//svc
	RegisterTask(k8s.SvcService().GetSvcList, k8s.GetSvcListInputParams{})
	RegisterTask(k8s.SvcService().GetSvcDetails, k8s.GetSvcDetailsInputParams{})
	RegisterTask(k8s.SvcService().DeploySvc, k8s.DeploySvcInputParams{})
	RegisterTask(k8s.SvcService().DeleteSvc, k8s.DeleteSvcInputParams{})

	//virtualService
	RegisterTask(k8s.VirtualServiceService().GetVirtualServiceList, k8s.GetVirtualServiceListInputParams{})
	RegisterTask(k8s.VirtualServiceService().GetVirtualServiceDetails, k8s.GetVirtualServiceDetailsInputParams{})
	RegisterTask(k8s.VirtualServiceService().DeployVirtualService, k8s.DeployVirtualServiceInputParams{})
	RegisterTask(k8s.VirtualServiceService().DeleteVirtualService, k8s.DeleteVirtualServiceInputParams{})

	//volumeSnapshot
	RegisterTask(k8s.VolumeSnapshotService().GetVolumeSnapshotList, k8s.GetVolumeSnapshotListInputParams{})
	RegisterTask(k8s.VolumeSnapshotService().GetVolumeSnapshotDetails, k8s.GetVolumeSnapshotDetailsInputParams{})
	RegisterTask(k8s.VolumeSnapshotService().DeployVolumeSnapshot, k8s.DeployVolumeSnapshotInputParams{})
	RegisterTask(k8s.VolumeSnapshotService().DeleteVolumeSnapshot, k8s.DeleteVolumeSnapshotInputParams{})

	//volumeSnapshotClass
	RegisterTask(k8s.VolumeSnapshotClassService().GetVolumeSnapshotClassList, k8s.GetVolumeSnapshotClassListInputParams{})
	RegisterTask(k8s.VolumeSnapshotClassService().GetVolumeSnapshotClassDetails, k8s.GetVolumeSnapshotClassDetailsInputParams{})

	//volumeSnapshotContent
	RegisterTask(k8s.VolumeSnapshotContentService().GetVolumeSnapshotContentList, k8s.GetVolumeSnapshotContentListInputParams{})
	RegisterTask(k8s.VolumeSnapshotContentService().GetVolumeSnapshotContentDetails, k8s.GetVolumeSnapshotContentDetailsInputParams{})
}
