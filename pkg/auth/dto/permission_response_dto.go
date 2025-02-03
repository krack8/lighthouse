package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

// PermissionResponse represents the formatted API response
type PermissionResponse struct {
	DEFAULT  []PermissionDTO `json:"DEFAULT"`
	CLUSTER  []PermissionDTO `json:"CLUSTER"`
	HelmApps []PermissionDTO `json:"HELM_APPS"`
}

// PermissionDTO represents the simplified permission response
type PermissionDTO struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
}
