package enum

// TableName is an enum for all table names in the database.
type TableName string

const (
	UsersTable       TableName = "users"       // Enum for the User table name
	PermissionsTable TableName = "permissions" // Enum for the User table name
	RolesTable       TableName = "roles"       // Enum for the User table name
)
