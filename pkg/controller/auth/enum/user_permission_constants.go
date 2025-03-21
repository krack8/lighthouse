package enum

// PermissionCategory represents different types of permissions
type PermissionCategory string

// PermissionName represents unique permission identifiers
type PermissionName string

// PermissionDescription represents permission descriptions
type PermissionDescription string

// ResourceGroup represents different type of resources in a category
type ResourceGroup string

const (
	//Permission categories
	DEFAULT         PermissionCategory = "DEFAULT"
	KUBERNETES      PermissionCategory = "KUBERNETES"
	USER_MANAGEMENT PermissionCategory = "USER_MANAGEMENT"
	HELM            PermissionCategory = "HELM"

	//CLUSTER Permission names
	DEFAULT_PERMISSION                PermissionName = "DEFAULT_PERMISSION"
	ADD_CLUSTER                       PermissionName = "ADD_CLUSTER"
	CREATE_NAMESPACE                  PermissionName = "CREATE_K8S_NAMESPACE"
	VIEW_NAMESPACE                    PermissionName = "VIEW_K8S_NAMESPACE"
	UPDATE_NAMESPACE                  PermissionName = "UPDATE_K8S_NAMESPACE"
	DELETE_NAMESPACE                  PermissionName = "DELETE_K8S_NAMESPACE"
	VIEW_DEPLOYMENT                   PermissionName = "VIEW_NAMESPACE_DEPLOYMENT"
	VIEW_REPLICA_SET                  PermissionName = "VIEW_NAMESPACE_REPLICA_SET"
	MANAGE_POD                        PermissionName = "MANAGE_NAMESPACE_POD"
	VIEW_POD                          PermissionName = "VIEW_NAMESPACE_POD"
	MANAGE_DEPLOYMENT                 PermissionName = "MANAGE_NAMESPACE_DEPLOYMENT"
	MANAGE_REPLICA_SET                PermissionName = "MANAGE_NAMESPACE_REPLICA_SET"
	VIEW_STATEFUL_SET                 PermissionName = "VIEW_NAMESPACE_STATEFUL_SET"
	MANAGE_STATEFUL_SET               PermissionName = "MANAGE_NAMESPACE_STATEFUL_SET"
	VIEW_DAEMON_SET                   PermissionName = "VIEW_NAMESPACE_DAEMON_SET"
	MANAGE_DAEMON_SET                 PermissionName = "MANAGE_NAMESPACE_DAEMON_SET"
	VIEW_SECRET                       PermissionName = "VIEW_NAMESPACE_SECRET"
	MANAGE_SECRET                     PermissionName = "MANAGE_NAMESPACE_SECRET"
	VIEW_CONFIG_MAP                   PermissionName = "VIEW_NAMESPACE_CONFIG_MAP"
	MANAGE_CONFIG_MAP                 PermissionName = "MANAGE_NAMESPACE_CONFIG_MAP"
	VIEW_SERVICE_ACCOUNT              PermissionName = "VIEW_NAMESPACE_SERVICE_ACCOUNT"
	MANAGE_SERVICE_ACCOUNT            PermissionName = "MANAGE_NAMESPACE_SERVICE_ACCOUNT"
	VIEW_SERVICE                      PermissionName = "VIEW_NAMESPACE_SERVICE"
	MANAGE_SERVICE                    PermissionName = "MANAGE_NAMESPACE_SERVICE"
	VIEW_INGRESS                      PermissionName = "VIEW_NAMESPACE_INGRESS"
	MANAGE_INGRESS                    PermissionName = "MANAGE_NAMESPACE_INGRESS"
	VIEW_CERTIFICATE                  PermissionName = "VIEW_NAMESPACE_CERTIFICATE"
	MANAGE_CERTIFICATE                PermissionName = "MANAGE_NAMESPACE_CERTIFICATE"
	VIEW_NAMESPACE_ROLE               PermissionName = "VIEW_NAMESPACE_ROLE"
	MANAGE_NAMESPACE_ROLE             PermissionName = "MANAGE_NAMESPACE_ROLE"
	VIEW_NAMESPACE_ROLE_BINDING       PermissionName = "VIEW_NAMESPACE_ROLE_BINDING"
	MANAGE_NAMESPACE_ROLE_BINDING     PermissionName = "MANAGE_NAMESPACE_ROLE_BINDING"
	VIEW_JOB                          PermissionName = "VIEW_NAMESPACE_JOB"
	MANAGE_JOB                        PermissionName = "MANAGE_NAMESPACE_JOB"
	VIEW_CRON_JOB                     PermissionName = "VIEW_NAMESPACE_CRON_JOB"
	MANAGE_CRON_JOB                   PermissionName = "MANAGE_NAMESPACE_CRON_JOB"
	VIEW_NAMESPACE_NETWORK_POLICY     PermissionName = "VIEW_NAMESPACE_NETWORK_POLICY"
	MANAGE_NAMESPACE_NETWORK_POLICY   PermissionName = "MANAGE_NAMESPACE_NETWORK_POLICY"
	VIEW_NAMESPACE_RESOURCE_QUOTA     PermissionName = "VIEW_NAMESPACE_RESOURCE_QUOTA"
	MANAGE_RESOURCE_QUOTA             PermissionName = "MANAGE_NAMESPACE_RESOURCE_QUOTA"
	VIEW_PERSISTENT_VOLUME            PermissionName = "VIEW_K8S_PERSISTENT_VOLUME"
	MANAGE_PERSISTENT_VOLUME          PermissionName = "MANAGE_K8S_PERSISTENT_VOLUME"
	VIEW_PERSISTENT_VOLUME_CLAIM      PermissionName = "VIEW_NAMESPACE_PERSISTENT_VOLUME_CLAIM"
	MANAGE_PERSISTENT_VOLUME_CLAIM    PermissionName = "MANAGE_NAMESPACE_PERSISTENT_VOLUME_CLAIM"
	VIEW_GATEWAY                      PermissionName = "VIEW_NAMESPACE_GATEWAY"
	MANAGE_GATEWAY                    PermissionName = "MANAGE_NAMESPACE_GATEWAY"
	VIEW_VIRTUAL_SERVICE              PermissionName = "VIEW_NAMESPACE_VIRTUAL_SERVICE"
	MANAGE_VIRTUAL_SERVICE            PermissionName = "MANAGE_NAMESPACE_VIRTUAL_SERVICE"
	VIEW_NODES                        PermissionName = "VIEW_K8S_NODES"
	MANAGE_NODE_TAINT                 PermissionName = "MANAGE_K8S_NODE_TAINT"
	DRAIN_NODE                        PermissionName = "DRAIN_K8S_NODE"
	VIEW_CLUSTER_ROLE                 PermissionName = "VIEW_K8S_CLUSTER_ROLE"
	MANAGE_CLUSTER_ROLE               PermissionName = "MANAGE_K8S_CLUSTER_ROLE"
	VIEW_CLUSTER_ROLE_BINDING         PermissionName = "VIEW_K8S_CLUSTER_ROLE_BINDING"
	MANAGE_CLUSTER_ROLE_BINDING       PermissionName = "MANAGE_K8S_CLUSTER_ROLE_BINDING"
	VIEW_STORAGE_CLASS                PermissionName = "VIEW_K8S_STORAGE_CLASS"
	MANAGE_STORAGE_CLASS              PermissionName = "MANAGE_K8S_STORAGE_CLASS"
	VIEW_CUSTOM_RESOURCES             PermissionName = "VIEW_K8S_CUSTOM_RESOURCES"
	MANAGE_CUSTOM_RESOURCES           PermissionName = "MANAGE_K8S_CUSTOM_RESOURCES"
	VIEW_CUSTOM_RESOURCE_DEFINITION   PermissionName = "VIEW_K8S_CUSTOM_RESOURCE_DEFINITION"
	MANAGE_CUSTOM_RESOURCE_DEFINITION PermissionName = "MANAGE_K8S_CUSTOM_RESOURCE_DEFINITION"
	VIEW_LOGS                         PermissionName = "VIEW_LOGS"
	VIEW_ENDPOINTS                    PermissionName = "VIEW_NAMESPACE_ENDPOINTS"
	MANAGE_ENDPOINTS                  PermissionName = "MANAGE_NAMESPACE_ENDPOINTS"
	VIEW_ENDPOINT_SLICE               PermissionName = "VIEW_NAMESPACE_ENDPOINT_SLICE"
	MANAGE_ENDPOINT_SLICE             PermissionName = "MANAGE_NAMESPACE_ENDPOINT_SLICE"
	VIEW_PDB                          PermissionName = "VIEW_NAMESPACE_PDB"
	MANAGE_PDB                        PermissionName = "MANAGE_NAMESPACE_PDB"
	VIEW_CONTROLLER_REVISION          PermissionName = "VIEW_NAMESPACE_CONTROLLER_REVISION"
	MANAGE_CONTROLLER_REVISION        PermissionName = "MANAGE_NAMESPACE_CONTROLLER_REVISION"
	VIEW_REPLICATION_CONTROLLER       PermissionName = "VIEW_NAMESPACE_REPLICATION_CONTROLLER"
	MANAGE_REPLICATION_CONTROLLER     PermissionName = "MANAGE_NAMESPACE_REPLICATION_CONTROLLER"

	//MANAGEMENT Permission names
	VIEW_USER   PermissionName = "VIEW_USER"
	MANAGE_USER PermissionName = "MANAGE_USER"

	VIEW_ROLE   PermissionName = "VIEW_ROLE"
	MANAGE_ROLE PermissionName = "MANAGE_ROLE"

	//CLUSTER Permission descriptions
	DEFAULT_PERMISSION_DESCRIPTION                PermissionDescription = "Default permission for basic access"
	ADD_CLUSTER_DESCRIPTION                       PermissionDescription = "Permission to add cluster with the control plane"
	CREATE_NAMESPACE_DESCRIPTION                  PermissionDescription = "Permission to create Kubernetes namespaces"
	VIEW_NAMESPACE_DESCRIPTION                    PermissionDescription = "Permission to view Kubernetes namespaces"
	UPDATE_NAMESPACE_DESCRIPTION                  PermissionDescription = "Permission to update Kubernetes namespaces"
	DELETE_NAMESPACE_DESCRIPTION                  PermissionDescription = "Permission to delete Kubernetes namespaces"
	VIEW_DEPLOYMENT_DESCRIPTION                   PermissionDescription = "Permission to view deployments"
	VIEW_REPLICA_SET_DESCRIPTION                  PermissionDescription = "Permission to view replica sets"
	MANAGE_POD_DESCRIPTION                        PermissionDescription = "Permission to create,update and delete pods"
	VIEW_POD_DESCRIPTION                          PermissionDescription = "Permission to view pods"
	MANAGE_DEPLOYMENT_DESCRIPTION                 PermissionDescription = "Permission to create,update and delete deployments"
	MANAGE_REPLICA_SET_DESCRIPTION                PermissionDescription = "Permission to create,update and delete replica sets"
	VIEW_STATEFUL_SET_DESCRIPTION                 PermissionDescription = "Permission to view stateful sets"
	MANAGE_STATEFUL_SET_DESCRIPTION               PermissionDescription = "Permission to create,update and delete stateful sets"
	VIEW_DAEMON_SET_DESCRIPTION                   PermissionDescription = "Permission to view daemon sets"
	MANAGE_DAEMON_SET_DESCRIPTION                 PermissionDescription = "Permission to create,update and delete daemon sets"
	VIEW_SECRET_DESCRIPTION                       PermissionDescription = "Permission to view secrets"
	MANAGE_SECRET_DESCRIPTION                     PermissionDescription = "Permission to create,update and delete secrets"
	VIEW_CONFIG_MAP_DESCRIPTION                   PermissionDescription = "Permission to view config maps"
	MANAGE_CONFIG_MAP_DESCRIPTION                 PermissionDescription = "Permission to create,update and delete config maps"
	VIEW_SERVICE_ACCOUNT_DESCRIPTION              PermissionDescription = "Permission to view service accounts"
	MANAGE_SERVICE_ACCOUNT_DESCRIPTION            PermissionDescription = "Permission to create,update and delete service accounts"
	VIEW_SERVICE_DESCRIPTION                      PermissionDescription = "Permission to view services"
	MANAGE_SERVICE_DESCRIPTION                    PermissionDescription = "Permission to create,update and delete services"
	VIEW_INGRESS_DESCRIPTION                      PermissionDescription = "Permission to view ingress resources"
	MANAGE_INGRESS_DESCRIPTION                    PermissionDescription = "Permission to create,update and delete ingress resources"
	VIEW_CERTIFICATE_DESCRIPTION                  PermissionDescription = "Permission to view certificates"
	MANAGE_CERTIFICATE_DESCRIPTION                PermissionDescription = "Permission to create,update and delete certificates"
	VIEW_NAMESPACE_ROLE_DESCRIPTION               PermissionDescription = "Permission to view roles"
	MANAGE_NAMESPACE_ROLE_DESCRIPTION             PermissionDescription = "Permission to create,update and delete roles"
	VIEW_NAMESPACE_ROLE_BINDING_DESCRIPTION       PermissionDescription = "Permission to view role bindings"
	MANAGE_NAMESPACE_ROLE_BINDING_DESCRIPTION     PermissionDescription = "Permission to create,update and delete role bindings"
	VIEW_JOB_DESCRIPTION                          PermissionDescription = "Permission to view jobs"
	MANAGE_JOB_DESCRIPTION                        PermissionDescription = "Permission to create,update and delete jobs"
	VIEW_CRON_JOB_DESCRIPTION                     PermissionDescription = "Permission to view cron jobs"
	MANAGE_CRON_JOB_DESCRIPTION                   PermissionDescription = "Permission to create,update and delete cron jobs"
	VIEW_NAMESPACE_NETWORK_POLICY_DESCRIPTION     PermissionDescription = "Permission to view network policies"
	MANAGE_NAMESPACE_NETWORK_POLICY_DESCRIPTION   PermissionDescription = "Permission to create,update and delete network policies"
	VIEW_NAMESPACE_RESOURCE_QUOTA_DESCRIPTION     PermissionDescription = "Permission to view resource quotas"
	MANAGE_RESOURCE_QUOTA_DESCRIPTION             PermissionDescription = "Permission to create,update and delete resource quotas"
	VIEW_PERSISTENT_VOLUME_DESCRIPTION            PermissionDescription = "Permission to view persistent volumes"
	MANAGE_PERSISTENT_VOLUME_DESCRIPTION          PermissionDescription = "Permission to create,update and delete persistent volumes"
	VIEW_PERSISTENT_VOLUME_CLAIM_DESCRIPTION      PermissionDescription = "Permission to view persistent volumes claim"
	MANAGE_PERSISTENT_VOLUME_CLAIM_DESCRIPTION    PermissionDescription = "Permission to create,update and delete persistent volumes claim"
	VIEW_GATEWAY_DESCRIPTION                      PermissionDescription = "Permission to view gateways"
	MANAGE_GATEWAY_DESCRIPTION                    PermissionDescription = "Permission to create,update and delete gateways"
	VIEW_VIRTUAL_SERVICE_DESCRIPTION              PermissionDescription = "Permission to view virtual services"
	MANAGE_VIRTUAL_SERVICE_DESCRIPTION            PermissionDescription = "Permission to create,update and delete virtual services"
	VIEW_NODES_DESCRIPTION                        PermissionDescription = "Permission to view Kubernetes nodes"
	MANAGE_NODE_TAINT_DESCRIPTION                 PermissionDescription = "Permission to create,update and delete Kubernetes node taints"
	DRAIN_NODE_DESCRIPTION                        PermissionDescription = "Permission to drain Kubernetes nodes"
	VIEW_CLUSTER_ROLE_DESCRIPTION                 PermissionDescription = "Permission to view Kubernetes cluster roles"
	MANAGE_CLUSTER_ROLE_DESCRIPTION               PermissionDescription = "Permission to create,update and delete Kubernetes cluster roles"
	VIEW_CLUSTER_ROLE_BINDING_DESCRIPTION         PermissionDescription = "Permission to view Kubernetes cluster role bindings"
	MANAGE_CLUSTER_ROLE_BINDING_DESCRIPTION       PermissionDescription = "Permission to create,update and delete Kubernetes cluster role bindings"
	VIEW_STORAGE_CLASS_DESCRIPTION                PermissionDescription = "Permission to view Kubernetes storage classes"
	MANAGE_STORAGE_CLASS_DESCRIPTION              PermissionDescription = "Permission to create,update and delete Kubernetes storage classes"
	VIEW_CUSTOM_RESOURCES_DESCRIPTION             PermissionDescription = "Permission to view Kubernetes custom resources"
	MANAGE_CUSTOM_RESOURCES_DESCRIPTION           PermissionDescription = "Permission to create,update and delete Kubernetes custom resources"
	VIEW_CUSTOM_RESOURCE_DEFINATION_DESCRIPTION   PermissionDescription = "Permission to view Kubernetes custom resource definitions"
	MANAGE_CUSTOM_RESOURCE_DEFINATION_DESCRIPTION PermissionDescription = "Permission to create,update and delete Kubernetes custom resource definitions"
	VIEW_LOGS_DESCRIPTION                         PermissionDescription = "Permission to view logs"
	VIEW_ENDPOINTS_DESCRIPTION                    PermissionDescription = "Permission to view endpoints"
	MANAGE_ENDPOINTS_DESCRIPTION                  PermissionDescription = "Permission to create,update and delete endpoints"
	VIEW_ENDPOINT_SLICE_DESCRIPTION               PermissionDescription = "Permission to view endpoint slices"
	MANAGE_ENDPOINT_SLICE_DESCRIPTION             PermissionDescription = "Permission to create,update and delete endpoint slices"
	VIEW_PDB_DESCRIPTION                          PermissionDescription = "Permission to view pod disruption budgets"
	MANAGE_PDB_DESCRIPTION                        PermissionDescription = "Permission to create,update and delete pod disruption budgets"
	VIEW_CONTROLLER_REVISION_DESCRIPTION          PermissionDescription = "Permission to view controller revisions"
	MANAGE_CONTROLLER_REVISION_DESCRIPTION        PermissionDescription = "Permission to create,update and delete controller revisions"
	VIEW_REPLICATION_CONTROLLER_DESCRIPTION       PermissionDescription = "Permission to view replication controllers"
	MANAGE_REPLICATION_CONTROLLER_DESCRIPTION     PermissionDescription = "Permission to create,update and delete replication controllers"

	//MANAGEMENT Permission names
	VIEW_USER_DESCRIPTION   PermissionDescription = "Permission to view user"
	MANAGE_USER_DESCRIPTION PermissionDescription = "Permission to create,update and delete user"

	VIEW_ROLE_DESCRIPTION   PermissionDescription = "Permission to view roles"
	MANAGE_ROLE_DESCRIPTION PermissionDescription = "Permission to create,update and delete roles"

	//Resource Group Name
	NODE                    ResourceGroup = "NODE"
	NAMESPACE               ResourceGroup = "NAMESPACE"
	PERSISTENT_VOLUME       ResourceGroup = "PERSISTENT_VOLUME"
	CLUSTER_ROLE            ResourceGroup = "CLUSTER_ROLE"
	CLUSTER_ROLE_BINDING    ResourceGroup = "CLUSTER_ROLE_BINDING"
	STORAGE_CLASS           ResourceGroup = "STORAGE_CLASS"
	CUSTOM_RESOURCE         ResourceGroup = "CUSTOM_RESOURCE"
	CERTIFICATE             ResourceGroup = "CERTIFICATE"
	CONFIG_MAP              ResourceGroup = "CONFIG_MAP"
	CRON_JOB                ResourceGroup = "CRON_JOB"
	CONTROLLER_REVISION     ResourceGroup = "CONTROLLER_REVISION"
	DAEMONSET               ResourceGroup = "DAEMONSET"
	DEPLOYMENT              ResourceGroup = "DEPLOYMENT"
	STATEFULSET             ResourceGroup = "STATEFULSET"
	ENDPOINT_SLICE          ResourceGroup = "ENDPOINT_SLICE"
	ENDPOINT                ResourceGroup = "ENDPOINT"
	INGRESS                 ResourceGroup = "INGRESS"
	JOB                     ResourceGroup = "JOB"
	NETWORK_POLICY          ResourceGroup = "NETWORK_POLICY"
	SECRET                  ResourceGroup = "SECRET"
	POD                     ResourceGroup = "POD"
	PERSISTENT_VOLUME_CLAIM ResourceGroup = "PERSISTENT_VOLUME_CLAIM"
	POD_DISRUPTION_BUDGET   ResourceGroup = "POD_DISRUPTION_BUDGET"
	REPLICASET              ResourceGroup = "REPLICASET"
	REPLICATION_CONTROLLER  ResourceGroup = "REPLICATION_CONTROLLER"
	RESOURCE_QUOTA          ResourceGroup = "RESOURCE_QUOTA"
	ROLE                    ResourceGroup = "ROLE"
	ROLE_BINDING            ResourceGroup = "ROLE_BINDING"
	SERVICE                 ResourceGroup = "SERVICE"
	SERVICE_ACCOUNT         ResourceGroup = "SERVICE_ACCOUNT"
	VIRTUAL_SERVICE         ResourceGroup = "VIRTUAL_SERVICE"
	GATEWAY                 ResourceGroup = "GATEWAY"

	USER ResourceGroup = "USER"
	NONE ResourceGroup = "NONE"
)
