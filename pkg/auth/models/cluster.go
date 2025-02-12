package models

import (
	"github.com/krack8/lighthouse/pkg/auth/enum"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Cluster struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name            string             `json:"name" bson:"name"`
	ClusterType     enum.ClusterType   `json:"cluster_type" bson:"cluster_type"` // "MASTER", "AGENT"
	Token           TokenValidation    `json:"-" bson:"token"`
	MasterClusterId string             `json:"master_cluster_id" bson:"master_cluster_id"`
	IsActive        bool               `json:"is_active" bson:"is_active"`
	Status          enum.Status        `json:"status" bson:"status"`
	ClusterStatus   enum.ClusterStatus `json:"cluster_status" bson:"cluster_status"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at"`
	CreatedBy       string             `json:"created_by" bson:"created_by"`
	UpdatedBy       string             `json:"updated_by" bson:"updated_by"`
}

type TokenValidation struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ClusterID   primitive.ObjectID `json:"cluster_id" bson:"cluster_id"`
	TokenHash   string             `json:"token" bson:"token"` // Bcrypt encrypted token
	IsValid     bool               `json:"is_valid" bson:"is_valid"`
	ExpiresAt   time.Time          `json:"expires_at" bson:"expires_at"`
	TokenStatus enum.TokenStatus   `json:"token_status" bson:"token_status"`
	Status      enum.Status        `json:"status" bson:"status"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	LastUsed    time.Time          `json:"last_used" bson:"last_used"`
	CreatedBy   string             `json:"created_by" bson:"created_by"`
	UpdatedBy   string             `json:"updated_by" bson:"updated_by"`
}
