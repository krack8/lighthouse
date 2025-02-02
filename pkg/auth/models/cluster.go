package models

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/krack8/lighthouse/pkg/auth/enum"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Cluster struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name            string             `json:"name" bson:"name"`
	ClusterType     enum.ClusterType   `json:"cluster_type" bson:"cluster_type"` // e.g., "MASTER", "AGENT"
	Token           string             `json:"-" bson:"token"`
	MasterClusterId string             `json:"masterClusterId" bson:"masterClusterId"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at"`
	IsActive        bool               `json:"is_active" bson:"is_active"`
}

type TokenValidation struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ClusterID primitive.ObjectID `json:"cluster_id" bson:"cluster_id"`
	Token     string             `json:"token" bson:"token"`
	IsValid   bool               `json:"is_valid" bson:"is_valid"`
	ExpiresAt time.Time          `json:"expires_at" bson:"expires_at"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

// Generate secure token
func generateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
