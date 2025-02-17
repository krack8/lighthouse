package services

import (
	"context"
	"fmt"
	db "github.com/krack8/lighthouse/pkg/auth/config"
	"github.com/krack8/lighthouse/pkg/auth/enum"
	"github.com/krack8/lighthouse/pkg/auth/models"
	"github.com/krack8/lighthouse/pkg/auth/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Storage interface {
	StoreToken(ctx context.Context, token *models.TokenValidation) error
	GetToken(ctx context.Context, token string) (*models.TokenValidation, error)
	TokenExists(ctx context.Context, token string) (bool, error)
	InvalidateToken(ctx context.Context, token string, clusterID primitive.ObjectID, updater string) error
	UpdateLastUsed(ctx context.Context, token string) error // Update the last-used timestamp for a token
}

type MongoStorage struct {
	collection *mongo.Collection
}

var mongoUpdate *MongoStorage

func (ms *MongoStorage) StoreToken(ctx context.Context, token *models.TokenValidation) error {
	_, err := ms.collection.InsertOne(ctx, token)
	return err
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

func (m *MongoStorage) UpdateLastUsed(ctx context.Context, token string) error {
	// Update the last-used timestamp for the token
	_, err := db.TokenCollection.UpdateOne(
		ctx,
		bson.M{"token": token},
		bson.M{"$set": bson.M{"last_used": time.Now()}},
	)
	return err
}

func (m *MongoStorage) GetToken(ctx context.Context, token string) (*models.TokenValidation, error) {
	// Query MongoDB for the token
	var result models.TokenValidation
	err := db.TokenCollection.FindOne(ctx, bson.M{"auth_token": token, "status": enum.VALID}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// validateTokenStatus checks the validity, expiry, and status of the token
func validateTokenStatus(token *models.TokenValidation) error {
	// 1. Check if the token is active (valid status)
	if token.TokenStatus != enum.TokenStatusValid {
		return fmt.Errorf("token is not active")
	}

	// 2. Check if the token is expired
	if time.Now().After(token.ExpiresAt) {
		//mark token as invalid.
		return fmt.Errorf("token has expired")
	}

	// 3. Check if the token is marked as valid in the database
	if !token.IsValid {
		return fmt.Errorf("token is invalid")
	}

	/*// 4. Optionally, check if the token is marked as used (if you want a "last used" validation)
	if !token.LastUsed.IsZero() && time.Since(token.LastUsed) > 3000*time.Hour { // Example check
		return fmt.Errorf("token has not been used for a long time")
	}*/

	// If all checks pass, the token is valid
	return nil
}

// ValidateToken validates the combined token and returns true/false along with clusterID
func ValidateToken(combinedToken string, validator *models.TokenValidation) (bool, primitive.ObjectID, error) {
	// Parse the combined token (encrypted raw token + cluster ID + signature)

	crypto, _ := utils.NewCryptoImpl()
	clusterID, rawToken, err := crypto.ParseCombinedToken(combinedToken)
	if err != nil {
		fmt.Printf("Error decoding combined token: %v\n", err)
		return false, primitive.NilObjectID, err
	}

	// Check if the decoded token matches the db token
	err = bcrypt.CompareHashAndPassword([]byte(validator.RawTokenHash), []byte(rawToken))
	if err != nil {
		fmt.Println("Token not matched. Unauthorized worker.")
		return false, primitive.NilObjectID, err
	}

	// Check if the decoded cluster ID matches the stored cluster ID
	if validator.ClusterID != clusterID {
		fmt.Println("Cluster ID mismatch.")
		return false, primitive.NilObjectID, err
	}

	// Validate the token's status, expiry, and validity
	if err := validateTokenStatus(validator); err != nil {
		fmt.Printf("Token validation failed: %v\n", err)
		return false, primitive.NilObjectID, err
	}

	// Update the "last used" timestamp for the token
	if err := mongoUpdate.UpdateLastUsed(context.Background(), combinedToken); err != nil {
		fmt.Printf("Failed to update last used time: %v\n", err)
		return false, primitive.NilObjectID, err
	}

	// If all checks pass, return true indicating the token is valid along with the clusterID
	return true, clusterID, nil
}
