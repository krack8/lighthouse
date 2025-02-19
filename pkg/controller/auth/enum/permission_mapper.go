package enum

// PermissionDefinition represents a permission's metadata
type PermissionDefinition struct {
	Name           PermissionName
	Description    PermissionDescription
	Category       PermissionCategory
	Resource_Group ResourceGroup
}

var PermissionInitializer = []PermissionDefinition{
	{
		Name:        DEFAULT_PERMISSION,
		Description: DEFAULT_PERMISSION_DESCRIPTION,
		Category:    DEFAULT,
	},
	{
		Name:        VIEW_USER,
		Description: VIEW_USER_DESCRIPTION,
		Category:    MANAGEMENT,
	},
	{
		Name:        MANAGE_USER,
		Description: MANAGE_USER_DESCRIPTION,
		Category:    MANAGEMENT,
	},
	{
		Name:        VIEW_ROLE,
		Description: VIEW_ROLE_DESCRIPTION,
		Category:    MANAGEMENT,
	},
	{
		Name:        MANAGE_ROLE,
		Description: MANAGE_ROLE_DESCRIPTION,
		Category:    MANAGEMENT,
	},
	{
		Name:        ADD_CLUSTER,
		Description: ADD_CLUSTER_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        CREATE_NAMESPACE,
		Description: CREATE_NAMESPACE_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        VIEW_NAMESPACE,
		Description: VIEW_NAMESPACE_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        UPDATE_NAMESPACE,
		Description: UPDATE_NAMESPACE_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        DELETE_NAMESPACE,
		Description: DELETE_NAMESPACE_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        VIEW_DEPLOYMENT,
		Description: VIEW_DEPLOYMENT_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        VIEW_REPLICA_SET,
		Description: VIEW_REPLICA_SET_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        MANAGE_POD,
		Description: MANAGE_POD_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        VIEW_POD,
		Description: VIEW_POD_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        MANAGE_DEPLOYMENT,
		Description: MANAGE_DEPLOYMENT_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        MANAGE_REPLICA_SET,
		Description: MANAGE_REPLICA_SET_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        VIEW_STATEFUL_SET,
		Description: VIEW_STATEFUL_SET_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        MANAGE_STATEFUL_SET,
		Description: MANAGE_STATEFUL_SET_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        VIEW_DAEMON_SET,
		Description: VIEW_DAEMON_SET_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        MANAGE_DAEMON_SET,
		Description: MANAGE_DAEMON_SET_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        VIEW_SECRET,
		Description: VIEW_SECRET_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        MANAGE_SECRET,
		Description: MANAGE_SECRET_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        VIEW_CONFIG_MAP,
		Description: VIEW_CONFIG_MAP_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        MANAGE_CONFIG_MAP,
		Description: MANAGE_CONFIG_MAP_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        VIEW_SERVICE_ACCOUNT,
		Description: VIEW_SERVICE_ACCOUNT_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        MANAGE_SERVICE_ACCOUNT,
		Description: MANAGE_SERVICE_ACCOUNT_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        VIEW_SERVICE,
		Description: VIEW_SERVICE_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        MANAGE_SERVICE,
		Description: MANAGE_SERVICE_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        VIEW_INGRESS,
		Description: VIEW_INGRESS_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        MANAGE_INGRESS,
		Description: MANAGE_INGRESS_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        VIEW_CERTIFICATE,
		Description: VIEW_CERTIFICATE_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        MANAGE_CERTIFICATE,
		Description: MANAGE_CERTIFICATE_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        VIEW_NAMESPACE_ROLE,
		Description: VIEW_NAMESPACE_ROLE_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        MANAGE_NAMESPACE_ROLE,
		Description: MANAGE_NAMESPACE_ROLE_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        VIEW_NAMESPACE_ROLE_BINDING,
		Description: VIEW_NAMESPACE_ROLE_BINDING_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        MANAGE_NAMESPACE_ROLE_BINDING,
		Description: MANAGE_NAMESPACE_ROLE_BINDING_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        VIEW_JOB,
		Description: VIEW_JOB_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        MANAGE_JOB,
		Description: MANAGE_JOB_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        VIEW_CUSTOM_RESOURCES,
		Description: VIEW_CUSTOM_RESOURCES_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        MANAGE_CUSTOM_RESOURCES,
		Description: MANAGE_CUSTOM_RESOURCES_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        VIEW_CUSTOM_RESOURCE_DEFINITION,
		Description: VIEW_CUSTOM_RESOURCE_DEFINATION_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        MANAGE_CUSTOM_RESOURCE_DEFINITION,
		Description: MANAGE_CUSTOM_RESOURCE_DEFINATION_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        VIEW_LOGS,
		Description: VIEW_LOGS_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        MANAGE_ENDPOINTS,
		Description: MANAGE_ENDPOINTS_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        VIEW_ENDPOINT_SLICE,
		Description: VIEW_ENDPOINT_SLICE_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        MANAGE_ENDPOINT_SLICE,
		Description: MANAGE_ENDPOINT_SLICE_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        VIEW_PDB,
		Description: VIEW_PDB_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        MANAGE_PDB,
		Description: MANAGE_PDB_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        VIEW_CONTROLLER_REVISION,
		Description: VIEW_CONTROLLER_REVISION_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        MANAGE_CONTROLLER_REVISION,
		Description: MANAGE_CONTROLLER_REVISION_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        VIEW_REPLICATION_CONTROLLER,
		Description: VIEW_REPLICATION_CONTROLLER_DESCRIPTION,
		Category:    CLUSTER,
	},
	{
		Name:        MANAGE_REPLICATION_CONTROLLER,
		Description: MANAGE_REPLICATION_CONTROLLER_DESCRIPTION,
		Category:    CLUSTER,
	},
	// Add more permission definitions here
}
