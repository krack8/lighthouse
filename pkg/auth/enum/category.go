package enum

type PermissionCategory string

const (
	DEFAULT PermissionCategory = "DEFAULT"
	CLUSTER PermissionCategory = "CLUSTER"
	HELM    PermissionCategory = "HELM"
)
