package enum

// TableName is an enum for all table names in the database.
type TableName string

const (
	UsersTable       TableName = "users"
	PermissionsTable TableName = "permissions"
	RolesTable       TableName = "roles"
	ClusterTable     TableName = "cluster"
	TokenTable       TableName = "token"
)
