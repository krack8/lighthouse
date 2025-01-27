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

}
