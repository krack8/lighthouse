package config

import (
	"context"
	"fmt"
	"github.com/krack8/lighthouse/pkg/common/config"
	"github.com/krack8/lighthouse/pkg/common/k8s"
	"github.com/krack8/lighthouse/pkg/common/log"
	"github.com/krack8/lighthouse/pkg/controller/auth/enum"
	"github.com/krack8/lighthouse/pkg/controller/auth/models"
	"github.com/krack8/lighthouse/pkg/controller/auth/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB collections
var (
	PermissionCollection *mongo.Collection
	RoleCollection       *mongo.Collection
	UserCollection       *mongo.Collection
	ClusterCollection    *mongo.Collection
	TokenCollection      *mongo.Collection
)

// ConnectDB initializes the MongoDB client and collections.
func ConnectDB() (*mongo.Client, context.Context, error) {
	ctx := context.Background()

	// Get the MongoDB URI and database name from environment variables
	uri := SetEnvWithDefault("MONGO_URI", "mongodb://localhost:27017")
	if uri == "" {
		log.Logger.Errorw("MONGO_URI environment variable is not set", "err", "MONGO_URI env missing")
		return nil, nil, fmt.Errorf("MONGO_URI not set")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Logger.Errorw("Failed to connect to MongoDB", "err", err.Error())
		return nil, nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Logger.Errorw("Failed to ping MongoDB", "err", err.Error())
		return nil, nil, err
	}

	dbName := SetEnvWithDefault("DB_NAME", "lighthouse")
	if dbName == "" {
		log.Logger.Errorw("DB_NAME environment variable is not set", "err", "DB_NAME env missing")
		return nil, nil, fmt.Errorf("DB_NAME not set")
	}

	// Initialize the collections
	UserCollection = client.Database(dbName).Collection(string(enum.UsersTable))
	PermissionCollection = client.Database(dbName).Collection(string(enum.PermissionsTable))
	RoleCollection = client.Database(dbName).Collection(string(enum.RolesTable))
	ClusterCollection = client.Database(dbName).Collection(string(enum.ClusterTable))
	TokenCollection = client.Database(dbName).Collection(string(enum.TokenTable))

	log.Logger.Info("Successfully connected to MongoDB")
	return client, ctx, nil
}

// InitializeDefaultUser creates a default user if no users exist.
func InitializeDefaultUser() {
	count, err := UserCollection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		log.Logger.Errorw("Error counting documents in users collection", "err", err.Error())
	}

	defaultUserName := SetEnvWithDefault("USER_EMAIL", "admin@default.com")
	defaultPassword := SetEnvWithDefault("PASSWORD", "lighthouse")
	if count == 0 {
		defaultUser := models.User{
			Username:     defaultUserName,
			FirstName:    "Admin",
			LastName:     "User",
			Password:     utils.HashPassword(defaultPassword), // Use a hashed password here
			UserType:     "ADMIN",
			Roles:        []models.Role{},
			UserIsActive: true,
			IsVerified:   true,
			Phone:        "1234567890",
			Status:       enum.VALID,
			CreatedBy:    string(enum.SYSTEM),
			UpdatedBy:    string(enum.SYSTEM),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		_, err := UserCollection.InsertOne(context.Background(), defaultUser)
		if err != nil {
			log.Logger.Errorw("Error creating default user", "err", err.Error())
		}

		log.Logger.Info("Default user created successfully")
	} else {
		log.Logger.Info("Users already exist. No default user created.")
	}
}

func InitRBAC() {
	// Initialize permissions
	initializer := NewPermissionInitializer(PermissionCollection)
	if err := initializer.InitializePermissions(context.Background()); err != nil {
		log.Logger.Errorw("Error Initializing permissions", "err", err.Error())
	}

	var defaultPermission models.Permission
	// Find permissions by name
	err := PermissionCollection.FindOne(context.Background(), bson.M{"name": string(enum.DEFAULT_PERMISSION)}).Decode(&defaultPermission)
	if err != nil {
		log.Logger.Errorw("Default permission not found", "err", err.Error())
	}

	roleCount, err := RoleCollection.CountDocuments(context.Background(), bson.M{"status": enum.VALID})
	if err != nil {
		log.Logger.Errorw("Error counting documents in users collection", "err", err.Error())
	}

	var defaultPermissions []models.Permission

	defaultPermissions = append(defaultPermissions, defaultPermission)

	if roleCount == 0 {
		// Example role with permissions
		defaultRole := models.Role{
			ID:          primitive.NewObjectID(),
			Name:        "DEFAULT_ROLE",
			Description: "Basic API Permissions",
			Permissions: defaultPermissions,
			Status:      enum.VALID,
			CreatedBy:   string(enum.SYSTEM),
			UpdatedBy:   string(enum.SYSTEM),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		// Insert the Default Role into the collection
		_, err = RoleCollection.InsertOne(context.Background(), defaultRole)
		if err != nil {
			log.Logger.Errorw("Error inserting role", "err", err.Error())
		}
	}
}

// InitializeClusters creates default clusters if none exist
func InitializeClusters() {
	clusterCount, err := ClusterCollection.CountDocuments(context.Background(), bson.M{"status": enum.VALID})
	if err != nil {
		log.Logger.Errorw("Error counting clusters", "err", err.Error())
	}
	if clusterCount == 0 || clusterCount == 1 {
		agentClusterID := primitive.NewObjectID()
		// Generate a raw token
		crypto, _ := utils.NewCryptoImpl()

		rawToken, err := crypto.GenerateSecureToken(32)
		if err != nil {
			log.Logger.Errorw("Error ailed to generate secure token", "err", err.Error())
		}

		// Create the combined token
		combinedToken, err := crypto.CreateCombinedToken(rawToken, agentClusterID)
		if err != nil {
			log.Logger.Errorw("Error failed to create combined token", "err", err.Error())
		}

		var existingCluster models.Cluster
		if clusterCount == 1 {
			err = ClusterCollection.FindOne(context.Background(), bson.M{"status": enum.VALID}).Decode(&existingCluster)
			combinedToken = existingCluster.Token.CombinedToken
			agentClusterID = existingCluster.ID
			if err != nil {
				log.Logger.Errorw("Error: failed to fetch first cluster entry from DB", "err", err.Error())
			}
		}
		k8s.InitiateKubeClientSet()
		// create the secret
		_, err = utils.CreateOrUpdateSecret(config.AgentSecretName, config.ResourceNamespace, combinedToken, agentClusterID.Hex(), config.RunMode)
		if err != nil {
			log.Logger.Errorw("Error failed to get secret", "err", err.Error())
		}

		if clusterCount == 0 {
			// Generate a bcrypt hash of the raw token with a default cost
			hashRawToken, err := bcrypt.GenerateFromPassword([]byte(rawToken), bcrypt.DefaultCost)
			if err != nil {
				log.Logger.Errorw("Error generating token hash", "err", err.Error())
			}

			// Create token validations
			agentToken := models.TokenValidation{
				ID:            primitive.NewObjectID(),
				ClusterID:     agentClusterID,
				RawTokenHash:  string(hashRawToken),
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

			_, err = TokenCollection.InsertOne(context.Background(), agentToken)
			if err != nil {
				log.Logger.Errorw("Error creating token validations", "err", err.Error())
			}

			clusterName := SetEnvWithDefault("DEFAULT_CLUSTER_NAME", "default-cluster")
			controllerGrpcServerHost := SetEnvWithDefault("CONTROLLER_GRPC_SERVER_HOST", "localhost:50051")
			// Create worker cluster
			agentCluster := models.Cluster{
				ID:                       agentClusterID,
				Name:                     clusterName,
				ClusterType:              enum.WORKER,
				AgentGroup:               agentClusterID.Hex(),
				Token:                    agentToken,
				Status:                   enum.VALID,
				ClusterStatus:            enum.PENDING,
				ControllerGrpcServerHost: controllerGrpcServerHost,
				CreatedBy:                string(enum.SYSTEM),
				UpdatedBy:                string(enum.SYSTEM),
				CreatedAt:                time.Now(),
				UpdatedAt:                time.Now(),
				IsActive:                 true,
			}

			// Insert clusters
			clusters := []interface{}{agentCluster}
			_, err = ClusterCollection.InsertMany(context.Background(), clusters)
			if err != nil {
				log.Logger.Errorw("Error creating default clusters", "err", err.Error())
			}

			log.Logger.Info("Default clusters and token validations created successfully")
		} else {
			log.Logger.Warn("Clusters already exist. No default cluster created.")
		}
	}
}

func SetEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		// Set the environment variable with the default value
		os.Setenv(key, defaultValue)
		return defaultValue
	}
	return value
}
