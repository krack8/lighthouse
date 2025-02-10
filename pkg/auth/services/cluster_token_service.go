package services

import (
	"context"
	"errors"
	db "github.com/krack8/lighthouse/pkg/auth/config"
	"github.com/krack8/lighthouse/pkg/auth/enum"
	"github.com/krack8/lighthouse/pkg/auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Storage interface {
	StoreToken(ctx context.Context, token *models.TokenValidation) error
	ValidateToken(ctx context.Context, token string, clusterID primitive.ObjectID) (*models.TokenValidation, error)
	TokenExists(ctx context.Context, token string) (bool, error)
	InvalidateToken(ctx context.Context, token string, clusterID primitive.ObjectID, updater string) error
	GetToken(ctx context.Context, token string) (*models.TokenValidation, error) // Retrieve a token by its value
	UpdateLastUsed(ctx context.Context, token string) error                      // Update the last-used timestamp for a token
}

type MongoStorage struct {
	collection *mongo.Collection
}

func NewMongoTokenStorage() Storage {
	return &MongoStorage{
		collection: db.TokenCollection,
	}
}

func (ms *MongoStorage) StoreToken(ctx context.Context, token *models.TokenValidation) error {
	_, err := ms.collection.InsertOne(ctx, token)
	return err
}

func (ms *MongoStorage) ValidateToken(ctx context.Context, token string, clusterID primitive.ObjectID) (*models.TokenValidation, error) {
	var tokenValidation models.TokenValidation
	err := ms.collection.FindOne(ctx, bson.M{
		"token":      token,
		"cluster_id": clusterID,
		"is_valid":   true,
		"status":     enum.VALID,
		"expires_at": bson.M{"$gt": time.Now()},
	}).Decode(&tokenValidation)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &tokenValidation, nil
}

func (ms *MongoStorage) TokenExists(ctx context.Context, token string) (bool, error) {
	count, err := ms.collection.CountDocuments(ctx, bson.M{"token": token})
	return count > 0, err
}

func (ms *MongoStorage) InvalidateToken(ctx context.Context, token string, clusterID primitive.ObjectID, updater string) error {
	_, err := ms.collection.UpdateOne(
		ctx,
		bson.M{
			"token":      token,
			"cluster_id": clusterID,
			"is_valid":   true,
		},
		bson.M{
			"$set": bson.M{
				"is_valid":   false,
				"status":     enum.DELETED,
				"updated_at": time.Now(),
				"updated_by": updater,
			},
		},
	)
	return err
}

func (m *MongoStorage) GetToken(ctx context.Context, token string) (*models.TokenValidation, error) {
	// Query MongoDB for the token
	var result models.TokenValidation
	err := m.collection.FindOne(ctx, bson.M{"token": token}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // Token not found
		}
		return nil, err
	}
	return &result, nil
}

func (m *MongoStorage) UpdateLastUsed(ctx context.Context, token string) error {
	// Update the last-used timestamp for the token
	_, err := m.collection.UpdateOne(
		ctx,
		bson.M{"token": token},
		bson.M{"$set": bson.M{"last_used": time.Now()}},
	)
	return err
}
