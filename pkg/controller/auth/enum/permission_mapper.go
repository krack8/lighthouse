package enum

// PermissionDefinition represents a permission's metadata
type PermissionDefinition struct {
	Name          PermissionName
	Description   PermissionDescription
	Category      PermissionCategory
	ResourceGroup ResourceGroup
}

var PermissionInitializer = []PermissionDefinition{
	{
		Name:          DEFAULT_PERMISSION,
		Description:   DEFAULT_PERMISSION_DESCRIPTION,
		Category:      DEFAULT,
		ResourceGroup: NONE,
	},
	{
		Name:          VIEW_USER,
		Description:   VIEW_USER_DESCRIPTION,
		Category:      MANAGEMENT,
		ResourceGroup: USER,
	},
	{
		Name:          MANAGE_USER,
		Description:   MANAGE_USER_DESCRIPTION,
		Category:      MANAGEMENT,
		ResourceGroup: USER,
	},
	{
		Name:          VIEW_ROLE,
		Description:   VIEW_ROLE_DESCRIPTION,
		Category:      MANAGEMENT,
		ResourceGroup: USER,
	},
	{
		Name:          MANAGE_ROLE,
		Description:   MANAGE_ROLE_DESCRIPTION,
		Category:      MANAGEMENT,
		ResourceGroup: USER,
	},
	{
		Name:          ADD_CLUSTER,
		Description:   ADD_CLUSTER_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: USER,
	},
	{
		Name:          VIEW_NODES,
		Description:   VIEW_NODES_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: NODE,
	},
	{
		Name:          MANAGE_NODE_TAINT,
		Description:   MANAGE_NODE_TAINT_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: NODE,
	},
	{
		Name:          DRAIN_NODE,
		Description:   DRAIN_NODE_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: NODE,
	},
	{
		Name:          VIEW_NAMESPACE,
		Description:   VIEW_NAMESPACE_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: NAMESPACE,
	},
	{
		Name:          CREATE_NAMESPACE,
		Description:   CREATE_NAMESPACE_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: NAMESPACE,
	},
	{
		Name:          UPDATE_NAMESPACE,
		Description:   UPDATE_NAMESPACE_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: NAMESPACE,
	},
	{
		Name:          DELETE_NAMESPACE,
		Description:   DELETE_NAMESPACE_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: NAMESPACE,
	},
	{
		Name:          VIEW_DEPLOYMENT,
		Description:   VIEW_DEPLOYMENT_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: DEPLOYMENT,
	},
	{
		Name:          MANAGE_DEPLOYMENT,
		Description:   MANAGE_DEPLOYMENT_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: DEPLOYMENT,
	},
	{
		Name:          VIEW_POD,
		Description:   VIEW_POD_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: POD,
	},
	{
		Name:          MANAGE_POD,
		Description:   MANAGE_POD_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: POD,
	},
	{
		Name:          VIEW_LOGS,
		Description:   VIEW_LOGS_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: POD,
	},
	{
		Name:          VIEW_REPLICA_SET,
		Description:   VIEW_REPLICA_SET_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: REPLICASET,
	},
	{
		Name:          MANAGE_REPLICA_SET,
		Description:   MANAGE_REPLICA_SET_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: REPLICASET,
	},
	{
		Name:          VIEW_STATEFUL_SET,
		Description:   VIEW_STATEFUL_SET_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: STATEFULSET,
	},
	{
		Name:          MANAGE_STATEFUL_SET,
		Description:   MANAGE_STATEFUL_SET_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: STATEFULSET,
	},
	{
		Name:          VIEW_DAEMON_SET,
		Description:   VIEW_DAEMON_SET_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: DAEMONSET,
	},
	{
		Name:          MANAGE_DAEMON_SET,
		Description:   MANAGE_DAEMON_SET_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: DAEMONSET,
	},
	{
		Name:          VIEW_CONFIG_MAP,
		Description:   VIEW_CONFIG_MAP_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: CONFIG_MAP,
	},
	{
		Name:          MANAGE_CONFIG_MAP,
		Description:   MANAGE_CONFIG_MAP_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: CONFIG_MAP,
	},
	{
		Name:          VIEW_SECRET,
		Description:   VIEW_SECRET_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: SECRET,
	},
	{
		Name:          MANAGE_SECRET,
		Description:   MANAGE_SECRET_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: SECRET,
	},
	{
		Name:          VIEW_SERVICE,
		Description:   VIEW_SERVICE_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: SERVICE,
	},
	{
		Name:          MANAGE_SERVICE,
		Description:   MANAGE_SERVICE_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: SERVICE,
	},
	{
		Name:          VIEW_SERVICE_ACCOUNT,
		Description:   VIEW_SERVICE_ACCOUNT_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: SERVICE_ACCOUNT,
	},
	{
		Name:          MANAGE_SERVICE_ACCOUNT,
		Description:   MANAGE_SERVICE_ACCOUNT_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: SERVICE_ACCOUNT,
	},
	{
		Name:          VIEW_INGRESS,
		Description:   VIEW_INGRESS_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: INGRESS,
	},
	{
		Name:          MANAGE_INGRESS,
		Description:   MANAGE_INGRESS_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: INGRESS,
	},
	{
		Name:          VIEW_PERSISTENT_VOLUME_CLAIM,
		Description:   VIEW_PERSISTENT_VOLUME_CLAIM_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: PERSISTENT_VOLUME_CLAIM,
	},
	{
		Name:          MANAGE_PERSISTENT_VOLUME_CLAIM,
		Description:   MANAGE_PERSISTENT_VOLUME_CLAIM_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: PERSISTENT_VOLUME_CLAIM,
	},
	{
		Name:          VIEW_CERTIFICATE,
		Description:   VIEW_CERTIFICATE_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: CERTIFICATE,
	},
	{
		Name:          MANAGE_CERTIFICATE,
		Description:   MANAGE_CERTIFICATE_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: CERTIFICATE,
	},
	{
		Name:          VIEW_NAMESPACE_ROLE,
		Description:   VIEW_NAMESPACE_ROLE_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: ROLE,
	},
	{
		Name:          MANAGE_NAMESPACE_ROLE,
		Description:   MANAGE_NAMESPACE_ROLE_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: ROLE,
	},
	{
		Name:          VIEW_NAMESPACE_ROLE_BINDING,
		Description:   VIEW_NAMESPACE_ROLE_BINDING_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: ROLE_BINDING,
	},
	{
		Name:          MANAGE_NAMESPACE_ROLE_BINDING,
		Description:   MANAGE_NAMESPACE_ROLE_BINDING_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: ROLE_BINDING,
	},
	{
		Name:          VIEW_JOB,
		Description:   VIEW_JOB_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: JOB,
	},
	{
		Name:          MANAGE_JOB,
		Description:   MANAGE_JOB_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: JOB,
	},
	{
		Name:          VIEW_CRON_JOB,
		Description:   VIEW_CRON_JOB_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: CRON_JOB,
	},
	{
		Name:          MANAGE_CRON_JOB,
		Description:   MANAGE_CRON_JOB_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: CRON_JOB,
	},
	{
		Name:          VIEW_NAMESPACE_NETWORK_POLICY,
		Description:   VIEW_NAMESPACE_NETWORK_POLICY_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: NETWORK_POLICY,
	},
	{
		Name:          MANAGE_NAMESPACE_NETWORK_POLICY,
		Description:   MANAGE_NAMESPACE_NETWORK_POLICY_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: NETWORK_POLICY,
	},
	{
		Name:          VIEW_NAMESPACE_RESOURCE_QUOTA,
		Description:   VIEW_NAMESPACE_RESOURCE_QUOTA_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: RESOURCE_QUOTA,
	},
	{
		Name:          MANAGE_RESOURCE_QUOTA,
		Description:   MANAGE_RESOURCE_QUOTA_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: RESOURCE_QUOTA,
	},
	{
		Name:          VIEW_VIRTUAL_SERVICE,
		Description:   VIEW_VIRTUAL_SERVICE_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: VIRTUAL_SERVICE,
	},
	{
		Name:          MANAGE_VIRTUAL_SERVICE,
		Description:   MANAGE_VIRTUAL_SERVICE_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: VIRTUAL_SERVICE,
	},
	//{
	//	Name:          VIEW_ENDPOINTS,
	//	Description:   VIEW_ENDPOINTS_DESCRIPTION,
	//	Category:      CLUSTER,
	//	ResourceGroup: ENDPOINT,
	//},
	//{
	//	Name:          MANAGE_ENDPOINTS,
	//	Description:   MANAGE_ENDPOINTS_DESCRIPTION,
	//	Category:      CLUSTER,
	//	ResourceGroup: ENDPOINT,
	//},
	{
		Name:          VIEW_ENDPOINT_SLICE,
		Description:   VIEW_ENDPOINT_SLICE_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: ENDPOINT_SLICE,
	},
	{
		Name:          MANAGE_ENDPOINT_SLICE,
		Description:   MANAGE_ENDPOINT_SLICE_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: ENDPOINT_SLICE,
	},
	{
		Name:          VIEW_PDB,
		Description:   VIEW_PDB_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: POD_DISRUPTION_BUDGET,
	},
	{
		Name:          MANAGE_PDB,
		Description:   MANAGE_PDB_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: POD_DISRUPTION_BUDGET,
	},
	{
		Name:          VIEW_CONTROLLER_REVISION,
		Description:   VIEW_CONTROLLER_REVISION_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: CONTROLLER_REVISION,
	},
	{
		Name:          MANAGE_CONTROLLER_REVISION,
		Description:   MANAGE_CONTROLLER_REVISION_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: CONTROLLER_REVISION,
	},
	{
		Name:          VIEW_REPLICATION_CONTROLLER,
		Description:   VIEW_REPLICATION_CONTROLLER_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: REPLICATION_CONTROLLER,
	},
	{
		Name:          MANAGE_REPLICATION_CONTROLLER,
		Description:   MANAGE_REPLICATION_CONTROLLER_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: REPLICATION_CONTROLLER,
	},
	{
		Name:          VIEW_PERSISTENT_VOLUME,
		Description:   VIEW_PERSISTENT_VOLUME_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: PERSISTENT_VOLUME,
	},
	{
		Name:          MANAGE_PERSISTENT_VOLUME,
		Description:   MANAGE_PERSISTENT_VOLUME_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: PERSISTENT_VOLUME,
	},
	{
		Name:          VIEW_CLUSTER_ROLE,
		Description:   VIEW_CLUSTER_ROLE_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: CLUSTER_ROLE,
	},
	{
		Name:          MANAGE_CLUSTER_ROLE,
		Description:   MANAGE_CLUSTER_ROLE_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: CLUSTER_ROLE,
	},
	{
		Name:          VIEW_CLUSTER_ROLE_BINDING,
		Description:   VIEW_CLUSTER_ROLE_BINDING_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: CLUSTER_ROLE_BINDING,
	},
	{
		Name:          MANAGE_CLUSTER_ROLE_BINDING,
		Description:   MANAGE_CLUSTER_ROLE_BINDING_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: CLUSTER_ROLE_BINDING,
	},
	{
		Name:          VIEW_STORAGE_CLASS,
		Description:   VIEW_STORAGE_CLASS_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: STORAGE_CLASS,
	},
	{
		Name:          MANAGE_STORAGE_CLASS,
		Description:   MANAGE_STORAGE_CLASS_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: STORAGE_CLASS,
	},
	{
		Name:          VIEW_CUSTOM_RESOURCE_DEFINITION,
		Description:   VIEW_CUSTOM_RESOURCE_DEFINATION_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: CUSTOM_RESOURCE,
	},
	{
		Name:          MANAGE_CUSTOM_RESOURCE_DEFINITION,
		Description:   MANAGE_CUSTOM_RESOURCE_DEFINATION_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: CUSTOM_RESOURCE,
	},
	{
		Name:          VIEW_CUSTOM_RESOURCES,
		Description:   VIEW_CUSTOM_RESOURCES_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: CUSTOM_RESOURCE,
	},
	{
		Name:          MANAGE_CUSTOM_RESOURCES,
		Description:   MANAGE_CUSTOM_RESOURCES_DESCRIPTION,
		Category:      CLUSTER,
		ResourceGroup: CUSTOM_RESOURCE,
	},
	// Add more permission definitions here
}
