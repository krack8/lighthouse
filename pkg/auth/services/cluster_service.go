package services

import (
	"context"
	"errors"
	"fmt"
	db "github.com/krack8/lighthouse/pkg/auth/config"
	"time"

	"github.com/krack8/lighthouse/pkg/auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ClusterService handles cluster-related business logic
type ClusterService struct {
	collection Collection
}

// NewClusterService creates a new ClusterService instance
func NewClusterService(collection Collection) *ClusterService {
	return &ClusterService{
		collection: collection,
	}
}

// GetCluster retrieves a cluster by ID
func (s *ClusterService) GetCluster(clusterID string) (*models.Cluster, error) {
	if clusterID == "" {
		return nil, errors.New("cluster ID cannot be empty")
	}

	objectID, err := primitive.ObjectIDFromHex(clusterID)
	if err != nil {
		return nil, fmt.Errorf("invalid cluster ID format: %w", err)
	}

	var cluster models.Cluster
	filter := bson.M{"_id": objectID}
	result := db.ClusterCollection.FindOne(context.Background(), filter)
	if err := result.Decode(&cluster); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("cluster not found")
		}
		return nil, fmt.Errorf("failed to fetch cluster: %w", err)
	}

	return &cluster, nil
}

// GetAllclusters retrieves all clusters
func (s *ClusterService) GetAllClusters() ([]models.Cluster, error) {
	cursor, err := db.ClusterCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch clusters: %w", err)
	}
	defer cursor.Close(context.Background())

	var clusters []models.Cluster
	for cursor.Next(context.Background()) {
		var cluster models.Cluster
		if err := cursor.Decode(&cluster); err != nil {
			return nil, fmt.Errorf("failed to decode cluster: %w", err)
		}
		clusters = append(clusters, cluster)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return clusters, nil
}

func ValidateClusterToken(token string) (*models.Cluster, error) {
	var tokenValidation models.TokenValidation
	err := db.TokenCollection.FindOne(context.Background(), bson.M{
		"token":      token,
		"is_valid":   true,
		"expires_at": bson.M{"$gt": time.Now()},
	}).Decode(&tokenValidation)

	if err != nil {
		return nil, fmt.Errorf("invalid or expired token")
	}

	var cluster models.Cluster
	err = db.ClusterCollection.FindOne(context.Background(), bson.M{
		"_id":       tokenValidation.ClusterID,
		"is_active": true,
	}).Decode(&cluster)

	if err != nil {
		return nil, fmt.Errorf("cluster not found or inactive")
	}

	return &cluster, nil
}
