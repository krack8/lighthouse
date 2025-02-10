package services

import (
	"context"
	"fmt"
	"github.com/krack8/lighthouse/pkg/auth/enum"
	"github.com/krack8/lighthouse/pkg/auth/errors"
	"github.com/krack8/lighthouse/pkg/auth/models"
	"github.com/krack8/lighthouse/pkg/auth/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type TokenManager struct {
	storage Storage
	crypto  utils.Crypto
}

type TokenValidationResult struct {
	IsValid   bool
	ClusterID string
	Error     error
	Token     *models.TokenValidation
}

type TokenValidator interface {
	ValidateToken(ctx context.Context, token string) (*TokenValidationResult, error)
	CreateToken(ctx context.Context, clusterID primitive.ObjectID, createdBy string) (*models.TokenValidation, error)
}

func NewTokenManager(storage Storage, crypto utils.Crypto) *TokenManager {
	return &TokenManager{
		storage: storage,
		crypto:  crypto,
	}
}

// ValidateToken validates the combined token
func (tm *TokenManager) ValidateToken(ctx context.Context, combinedToken string) (*TokenValidationResult, error) {
	result := &TokenValidationResult{
		IsValid: false,
	}

	// Parse the combined token (encrypted raw token + cluster ID + signature)
	clusterID, rawToken, err := tm.crypto.ParseCombinedToken(combinedToken)
	if err != nil {
		result.Error = errors.ErrInvalidSignature
		return result, nil
	}

	result.ClusterID = clusterID.Hex()

	// Retrieve the stored combined token from the database based on the cluster ID
	tokenValidation, err := tm.storage.GetToken(ctx, rawToken)
	if err != nil {
		result.Error = errors.ErrTokenNotFound
		return result, nil
	}

	if tokenValidation == nil {
		result.Error = errors.ErrTokenNotFound
		return result, nil
	}

	// The combined token is already stored, we now verify its validity
	if err := tm.validateTokenStatus(tokenValidation); err != nil {
		result.Error = err
		return result, nil
	}

	// Update the "last used" timestamp for the token
	if err := tm.storage.UpdateLastUsed(ctx, combinedToken); err != nil {
		fmt.Printf("Failed to update last used time: %v\n", err)
	}

	result.IsValid = true
	result.Token = tokenValidation
	return result, nil
}

// Store the combined token
func (tm *TokenManager) CreateToken(clusterID primitive.ObjectID) (*models.TokenValidation, error) {
	// Generate a raw token
	rawToken, err := tm.crypto.GenerateSecureToken(32)
	if err != nil {
		return nil, fmt.Errorf("failed to generate secure token: %w", err)
	}

	// Create the combined token
	combinedToken, err := tm.crypto.CreateCombinedToken(rawToken, clusterID)
	if err != nil {
		return nil, fmt.Errorf("failed to create combined token: %w", err)
	}

	// Store the token in the database
	tokenValidation := &models.TokenValidation{
		ClusterID: clusterID,
		TokenHash: combinedToken,
		IsValid:   true,
		CreatedBy: string(enum.SYSTEM),
		CreatedAt: time.Now(),
	}

	if err := tm.storage.StoreToken(context.Background(), tokenValidation); err != nil {
		return nil, fmt.Errorf("failed to store token: %w", err)
	}

	return tokenValidation, nil
}

// validateTokenStatus checks the validity, expiry, and status of the token
func (tm *TokenManager) validateTokenStatus(token *models.TokenValidation) error {
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
