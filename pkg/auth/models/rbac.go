package models

import (
	"github.com/krack8/lighthouse/pkg/auth/enum"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Permission represents a specific action that can be performed on a resource
type Permission struct {
	ID           primitive.ObjectID      `json:"id" bson:"_id,omitempty"`
	Name         string                  `json:"name" bson:"name"`
	Description  string                  `json:"description" bson:"description"`
	EndpointList []Endpoint              `json:"endpoint_list" bson:"endpoint_list"`
	Category     enum.PermissionCategory `json:"category" bson:"category"`
	Status       enum.Status             `json:"status" bson:"status"`
	CreatedAt    time.Time               `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time               `json:"updated_at" bson:"updated_at"`
	CreatedBy    string                  `json:"created_by" bson:"created_by"`
	UpdatedBy    string                  `json:"updated_by" bson:"updated_by"`
}

type Endpoint struct {
	Route  string `json:"route" bson:"route"`   // URL path
	Method string `json:"method" bson:"method"` // HTTP method (GET, POST, etc.)
}

// Role represents a user role with associated permissions
type Role struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Permissions []Permission       `json:"permissions" bson:"permissions"`
	Status      enum.Status        `json:"status" bson:"status"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	CreatedBy   string             `json:"created_by" bson:"created_by"`
	UpdatedBy   string             `json:"updated_by" bson:"updated_by"`
}
