package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Permission represents a specific action that can be performed on a resource
type Permission struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Route       string             `json:"route" bson:"route"`   // URL path
	Method      string             `json:"method" bson:"method"` // HTTP method (GET, POST, etc.)
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

// Role represents a user role with associated permissions
type Role struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Permissions []Permission       `json:"permissions" bson:"permissions"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}
