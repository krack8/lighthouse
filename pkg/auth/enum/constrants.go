package enum

// ClusterType represents different types of clusters
type ClusterType string

// Mongo data Status represents
type Status string

// Agent token Status represents
type TokenStatus string

const (
	MASTER ClusterType = "MASTER"
	AGENT  ClusterType = "AGENT"
)

const (
	VALID   Status = "V"
	DELETED Status = "D"
	HIDDEN  Status = "H"
	SYSTEM  Status = "SYSTEM"
)

const (
	TokenStatusValid   TokenStatus = "ACTIVE"
	TokenStatusExpired TokenStatus = "EXPIRED"
	TokenStatusRevoked TokenStatus = "REVOKED"
)
