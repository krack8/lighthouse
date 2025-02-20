package models

import (
	"github.com/krack8/lighthouse/pkg/controller/auth/enum"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Cluster struct {
	ID                       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name                     string             `json:"name" bson:"name"`
	ClusterType              enum.ClusterType   `json:"cluster_type" bson:"cluster_type"` // "MASTER", "WORKER"
	Token                    TokenValidation    `json:"-" bson:"token"`
	AgentGroup               string             `json:"agent_group" bson:"agent_group"`
	ControllerGrpcServerHost string             `json:"controller_grpc_server_host" bson:"controller_grpc_server_host"`
	IsActive                 bool               `json:"is_active" bson:"is_active"`
	Details                  interface{}        `json:"details" bson:"details"`
	Status                   enum.Status        `json:"status" bson:"status"`
	ClusterStatus            enum.ClusterStatus `json:"cluster_status" bson:"cluster_status"`
	CreatedAt                time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt                time.Time          `json:"updated_at" bson:"updated_at"`
	CreatedBy                string             `json:"created_by" bson:"created_by"`
	UpdatedBy                string             `json:"updated_by" bson:"updated_by"`
}

type TokenValidation struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ClusterID     primitive.ObjectID `json:"cluster_id" bson:"cluster_id"`
	RawTokenHash  string             `json:"raw_token" bson:"raw_token"`   // Bcrypt encrypted raw token
	CombinedToken string             `json:"auth_token" bson:"auth_token"` // Bcrypt encrypted agent auth token
	IsValid       bool               `json:"is_valid" bson:"is_valid"`
	ExpiresAt     time.Time          `json:"expires_at" bson:"expires_at"`
	TokenStatus   enum.TokenStatus   `json:"token_status" bson:"token_status"`
	Status        enum.Status        `json:"status" bson:"status"`
	CreatedAt     time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at"`
	LastUsed      time.Time          `json:"last_used" bson:"last_used"`
	CreatedBy     string             `json:"created_by" bson:"created_by"`
	UpdatedBy     string             `json:"updated_by" bson:"updated_by"`
}
