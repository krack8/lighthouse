package enum

// PermissionCategory represents different types of permissions
type PermissionCategory string

// PermissionName represents unique permission identifiers
type PermissionName string

// PermissionDescription represents permission descriptions
type PermissionDescription string

const (
	// Permission categories
	DEFAULT PermissionCategory = "DEFAULT"
	CLUSTER PermissionCategory = "CLUSTER"
	HELM    PermissionCategory = "HELM"

	// Permission names
	VIEW_NAMESPACE        PermissionName = "VIEW_NAMESPACE_ENDPOINTS"
	MANAGE_NAMESPACE      PermissionName = "MANAGE_NAMESPACE_ENDPOINTS"
	VIEW_ENDPOINT_SLICE   PermissionName = "VIEW_ENDPOINT_SLICE"
	MANAGE_ENDPOINT_SLICE PermissionName = "MANAGE_ENDPOINT_SLICE"
	VIEW_PDB              PermissionName = "VIEW_PDB"
	MANAGE_PDB            PermissionName = "MANAGE_PDB"

	// Permission descriptions
	VIEW_NAMESPACE_DESCRIPTION        PermissionDescription = "Permission to view namespace endpoints"
	MANAGE_NAMESPACE_DESCRIPTION      PermissionDescription = "Permission to manage namespace endpoints"
	VIEW_ENDPOINT_SLICE_DESCRIPTION   PermissionDescription = "Permission to view endpoint slices"
	MANAGE_ENDPOINT_SLICE_DESCRIPTION PermissionDescription = "Permission to manage endpoint slices"
	VIEW_PDB_DESCRIPTION              PermissionDescription = "Permission to view PDBs"
	MANAGE_PDB_DESCRIPTION            PermissionDescription = "Permission to manage PDBs"
)
