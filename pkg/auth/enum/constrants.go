package enum

// ClusterType represents different types of clusters
type ClusterType string

// Status represents
type Status string

const (
	MASTER ClusterType = "MASTER"
	AGENT  ClusterType = "AGENT"
)

const (
	VALID   Status = "V"
	DELETED Status = "D"
)
