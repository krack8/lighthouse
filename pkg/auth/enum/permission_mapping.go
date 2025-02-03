package enum

// PermissionDefinition represents a permission's metadata
type PermissionDefinition struct {
	Description PermissionDescription
	Category    PermissionCategory
}

// PermissionDefinitions maps permission names to their definitions
var PermissionDefinitions = map[PermissionName]PermissionDefinition{
	DEFAULT_PERMISSION: {
		Description: DEFAULT_PERMISSION_DESCRIPTION,
		Category:    DEFAULT,
	},
	VIEW_NAMESPACE: {
		Description: VIEW_NAMESPACE_DESCRIPTION,
		Category:    CLUSTER,
	},
	MANAGE_NAMESPACE: {
		Description: MANAGE_NAMESPACE_DESCRIPTION,
		Category:    CLUSTER,
	},
	/*	VIEW_ENDPOINT_SLICE: {
			Description: VIEW_ENDPOINT_SLICE_DESCRIPTION,
			Category:    CLUSTER,
		},
		MANAGE_ENDPOINT_SLICE: {
			Description: MANAGE_ENDPOINT_SLICE_DESCRIPTION,
			Category:    CLUSTER,
		},
		VIEW_PDB: {
			Description: VIEW_PDB_DESCRIPTION,
			Category:    CLUSTER,
		},
		MANAGE_PDB: {
			Description: MANAGE_PDB_DESCRIPTION,
			Category:    CLUSTER,
		},*/
	// Add more permission definitions here
}
