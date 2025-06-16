package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/krack8/lighthouse/pkg/common/log"
	db "github.com/krack8/lighthouse/pkg/controller/auth/config"
	"github.com/krack8/lighthouse/pkg/controller/auth/enum"
	"github.com/krack8/lighthouse/pkg/controller/auth/models"
	"github.com/krack8/lighthouse/pkg/controller/auth/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var storage Storage     // Implement the Storage interface
var crypto utils.Crypto // Implement the Crypto interface

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
	filter := bson.M{"_id": objectID, "status": enum.VALID}
	result := db.ClusterCollection.FindOne(context.Background(), filter)
	if err := result.Decode(&cluster); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("cluster not found")
		}
		return nil, fmt.Errorf("failed to fetch cluster: %w", err)
	}

	return &cluster, nil
}

func (s *ClusterService) GetAllClusters(reqUsername string) ([]models.Cluster, error) {
	var filter bson.M
	requester, _ := GetUserByUsername(reqUsername)
	if requester == nil {
		return nil, errors.New("failed to fetch requester")
	}

	// If user is ADMIN, return all worker clusters
	if requester.UserType == "ADMIN" {
		filter = bson.M{"cluster_type": bson.M{"$eq": enum.WORKER}, "status": bson.M{"$eq": enum.VALID}}
	} else {
		// For USER type, only return clusters in their ClusterIdList
		if len(requester.ClusterIdList) == 0 {
			return []models.Cluster{}, nil // Return empty list if user has no clusters
		}

		// Convert cluster IDs to ObjectIDs
		var objectIDs []primitive.ObjectID
		for _, idStr := range requester.ClusterIdList {
			objectID, err := primitive.ObjectIDFromHex(idStr)
			if err != nil {
				return nil, fmt.Errorf("invalid cluster ID format: %s, error: %w", idStr, err)
			}
			objectIDs = append(objectIDs, objectID)
		}

		// Filter by user's cluster IDs and also maintain the worker type and valid status filters
		filter = bson.M{
			"_id":          bson.M{"$in": objectIDs},
			"cluster_type": bson.M{"$eq": enum.WORKER},
			"status":       bson.M{"$eq": enum.VALID},
		}
	}

	cursor, err := db.ClusterCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch clusters: %w", err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			// Log the error but don't return it as it's a cleanup operation
			log.Logger.Warnw("error closing cursor: %v", err)
		}
	}(cursor, context.Background())

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

func (s *ClusterService) GetClusterList() ([]models.Cluster, error) {
	// Filter for AGENT clusters
	filter := bson.M{"cluster_type": bson.M{"$eq": enum.WORKER}, "status": bson.M{"$eq": enum.VALID}}

	cursor, err := db.ClusterCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch non-master clusters: %w", err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, context.Background())

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

// UpdateClusterStatusToActive updates the cluster's status to "active"
func UpdateClusterStatusToActive(clusterID primitive.ObjectID) error {
	// Update the cluster status to "active" for the given cluster ID
	_, err := db.ClusterCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": clusterID},
		bson.M{"$set": bson.M{"is_active": true, "cluster_status": enum.CONNECTED}},
	)
	return err
}

// CreateCluster creates a new cluster and inserts it into the database
func (s *ClusterService) CreateAgentCluster(name, controllerGrpcServerHost string) (*models.Cluster, error) {
	agentClusterID := primitive.NewObjectID()

	// Generate a raw token
	crypto, _ := utils.NewCryptoImpl()

	rawToken, err := crypto.GenerateSecureToken(32)
	if err != nil {
		log.Logger.Errorw("Failed to generate secure token:", "err", err.Error())
	}

	// Create the combined token
	combinedToken, err := crypto.CreateCombinedToken(rawToken, agentClusterID)
	if err != nil {
		log.Logger.Errorw("Failed to create combined token:", "err", err.Error())
	}

	// Create token validations
	agentToken := models.TokenValidation{
		ID:            primitive.NewObjectID(),
		ClusterID:     agentClusterID,
		RawTokenHash:  utils.HashPassword(rawToken),
		CombinedToken: combinedToken,
		IsValid:       true,
		ExpiresAt:     time.Now().AddDate(1, 0, 0), // Token valid for 1 year
		Status:        enum.VALID,
		TokenStatus:   enum.TokenStatusValid,
		CreatedBy:     string(enum.SYSTEM),
		UpdatedBy:     string(enum.SYSTEM),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	_, err = db.TokenCollection.InsertOne(context.Background(), agentToken)
	if err != nil {
		log.Logger.Errorw("Error creating token validations:", "err", err.Error())
	}

	// Create a new cluster
	cluster := &models.Cluster{
		ID:                       agentClusterID,
		Name:                     name,
		ClusterType:              enum.WORKER, // Set default cluster type to WORKER
		Token:                    agentToken,
		AgentGroup:               agentClusterID.Hex(),
		IsActive:                 false,
		ControllerGrpcServerHost: controllerGrpcServerHost,
		Status:                   enum.VALID,
		ClusterStatus:            enum.PENDING,
		CreatedAt:                time.Now(),
		UpdatedAt:                time.Now(),
		CreatedBy:                string(enum.SYSTEM),
		UpdatedBy:                string(enum.SYSTEM),
	}

	// Insert the new cluster into the MongoDB collection
	_, err = db.ClusterCollection.InsertOne(context.Background(), cluster)
	if err != nil {
		return nil, fmt.Errorf("failed to insert cluster into database: %w", err)
	}

	return cluster, nil
}

// DeleteClusterByID deletes a cluster by its ID
func (s *ClusterService) DeleteClusterByID(clusterId string) error {
	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(clusterId)
	if err != nil {
		return fmt.Errorf("invalid cluster ID format: %v", err)
	}

	_, err = db.ClusterCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": objectID},
		bson.M{"$set": bson.M{"status": enum.DELETED, "cluster_status": enum.DISCONNECTED}},
	)
	if err != nil {
		return fmt.Errorf("failed to delete cluster: %v", err)
	}
	_, err = db.TokenCollection.UpdateOne(
		context.Background(),
		bson.M{"cluster_id": objectID},
		bson.M{"$set": bson.M{"status": enum.DELETED, "updated_at": time.Now()}},
	)
	if err != nil {
		return fmt.Errorf("failed to delete token: %v", err)
	}
	return nil
}

// UpdateUser updates a cluster by ID
func (s *ClusterService) RenameCluster(userID string, updatedUser *models.Cluster) error {
	if updatedUser == nil {
		return errors.New("input cannot be nil")
	}

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("invalid cluster ID format: %w", err)
	}

	// First fetch the existing cluster
	var existingCluster models.Cluster
	filter := bson.M{"_id": objectID, "status": enum.VALID}
	err = db.ClusterCollection.FindOne(context.Background(), filter).Decode(&existingCluster)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("cluster not found")
		}
		return fmt.Errorf("failed to fetch existing cluster: %w", err)
	}

	// Create update map with only non-empty fields
	updateFields := bson.M{}

	if updatedUser.Name != "" {
		updateFields["name"] = updatedUser.Name
	}

	if updatedUser.DefaultCluster == true {
		filter := bson.M{"status": enum.VALID}

		cursor, err := db.ClusterCollection.Find(context.Background(), filter)
		if err != nil {
			return fmt.Errorf("failed to fetch valid clusters: %w", err)
		}
		defer cursor.Close(context.Background())

		var clusters []models.Cluster
		if err = cursor.All(context.Background(), &clusters); err != nil {
			return fmt.Errorf("failed to decode clusters: %w", err)
		}

		for _, cluster := range clusters {
			update := bson.M{
				"$set": bson.M{"default_cluster": false},
			}
			_, err := db.ClusterCollection.UpdateByID(context.Background(), cluster.ID, update)
			if err != nil {
				return fmt.Errorf("failed to update cluster %v: %w", cluster.ID, err)
			}
		}
		updateFields["default_cluster"] = updatedUser.DefaultCluster
	}

	// Always update the UpdatedAt timestamp
	updateFields["updated_at"] = time.Now()

	// Only perform update if there are fields to update
	if len(updateFields) > 0 {
		update := bson.M{"$set": updateFields}
		result, err := db.ClusterCollection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return fmt.Errorf("failed to update cluster: %w", err)
		}

		if result.MatchedCount == 0 {
			return errors.New("cluster not found")
		}
	}

	return nil
}
