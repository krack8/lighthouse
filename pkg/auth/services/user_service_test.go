package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/krack8/lighthouse/pkg/auth/models"
	"github.com/krack8/lighthouse/pkg/auth/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MockCollection is a mock implementation of the Collection interface.
type MockCollection struct {
	mock.Mock
}

// MockSingleResult is a custom mock for mongo.SingleResult
type MockSingleResult struct {
	err  error
	user *models.User
}

func NewMockSingleResult(user *models.User, err error) *MockSingleResult {
	return &MockSingleResult{
		user: user,
		err:  err,
	}
}

func (m *MockSingleResult) Decode(v interface{}) error {
	if m.err != nil {
		return m.err
	}
	if m.user == nil {
		return mongo.ErrNoDocuments
	}

	userPtr, ok := v.(*models.User)
	if !ok {
		return errors.New("invalid type assertion")
	}
	*userPtr = *m.user
	return nil
}

func (m *MockCollection) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockCollection) FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult {
	args := m.Called(ctx, filter)
	if mockResult, ok := args.Get(0).(*mongo.SingleResult); ok {
		return mockResult
	}
	// Return a new mock result if the type doesn't match
	return &mongo.SingleResult{}
}

func (m *MockCollection) Find(ctx context.Context, filter interface{}) (*mongo.Cursor, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.Cursor), args.Error(1)
}

func (m *MockCollection) UpdateOne(ctx context.Context, filter, update interface{}) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func (m *MockCollection) DeleteOne(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.DeleteResult), args.Error(1)
}

func TestUserService_CreateUser(t *testing.T) {
	mockCollection := new(MockCollection)
	service := &services.UserService{
		UserCollection: mockCollection,
		Context:        context.Background(),
	}

	user := &models.User{
		FirstName: "John",
		LastName:  "Doe",
		Username:  "johndoe",
	}

	mockCollection.On("InsertOne", service.Context, user).
		Return(&mongo.InsertOneResult{}, nil)

	createdUser, err := service.CreateUser(user)

	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, "John", createdUser.FirstName)
	mockCollection.AssertExpectations(t)
}

func TestUserService_GetUser(t *testing.T) {
	mockCollection := new(MockCollection)
	service := &services.UserService{
		UserCollection: mockCollection,
		Context:        context.Background(),
	}

	/*t.Run("Success", func(t *testing.T) {
		mockUserID := primitive.NewObjectID()
		expectedUser := &models.User{
			FirstName: "Jane",
			LastName:  "Doe",
			Username:  "janedoe",
		}

		// Create a mock result with the expected user
		mockResult := NewMockSingleResult(expectedUser, nil)
		mockCollection.On("FindOne", service.Context, mock.Anything).
			Return(mockResult)

		user, err := service.GetUser(mockUserID.Hex())

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedUser.FirstName, user.FirstName)
		assert.Equal(t, expectedUser.LastName, user.LastName)
		assert.Equal(t, expectedUser.Username, user.Username)
		mockCollection.AssertExpectations(t)
	})*/

	t.Run("UserNotFound", func(t *testing.T) {
		mockUserID := primitive.NewObjectID()

		// Create a mock result with no user and ErrNoDocuments
		mockResult := NewMockSingleResult(nil, mongo.ErrNoDocuments)
		mockCollection.On("FindOne", service.Context, mock.Anything).
			Return(mockResult)

		user, err := service.GetUser(mockUserID.Hex())

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "user not found", err.Error())
		mockCollection.AssertExpectations(t)
	})

	t.Run("InvalidUserID", func(t *testing.T) {
		user, err := service.GetUser("invalid-id")

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "invalid user ID", err.Error())
	})

	t.Run("DatabaseError", func(t *testing.T) {
		mockUserID := primitive.NewObjectID()

		// Create a mock result with a database error
		mockResult := NewMockSingleResult(nil, errors.New("database error"))
		mockCollection.On("FindOne", service.Context, mock.Anything).
			Return(mockResult)

		user, err := service.GetUser(mockUserID.Hex())

		assert.Error(t, err)
		assert.Nil(t, user)
		mockCollection.AssertExpectations(t)
	})
}

func TestUserService_UpdateUser(t *testing.T) {
	mockCollection := new(MockCollection)
	service := &services.UserService{
		UserCollection: mockCollection,
		Context:        context.Background(),
	}

	mockUserID := primitive.NewObjectID()
	updatedUser := &models.User{
		FirstName: "Jane",
		LastName:  "Smith",
	}

	mockCollection.On("UpdateOne", service.Context, mock.Anything, mock.Anything).
		Return(&mongo.UpdateResult{MatchedCount: 1}, nil)

	err := service.UpdateUser(mockUserID.Hex(), updatedUser)

	assert.NoError(t, err)
	mockCollection.AssertExpectations(t)
}

func TestUserService_DeleteUser(t *testing.T) {
	mockCollection := new(MockCollection)
	service := &services.UserService{
		UserCollection: mockCollection,
		Context:        context.Background(),
	}

	mockUserID := primitive.NewObjectID()

	mockCollection.On("DeleteOne", service.Context, mock.Anything).
		Return(&mongo.DeleteResult{DeletedCount: 1}, nil)

	err := service.DeleteUser(mockUserID.Hex())

	assert.NoError(t, err)
	mockCollection.AssertExpectations(t)
}

// Test function
func TestUserService_GetUserByUsername(t *testing.T) {
	mockCollection := new(MockCollection)
	service := &services.UserService{
		UserCollection: mockCollection,
		Context:        context.Background(),
	}

	/*t.Run("Success", func(t *testing.T) {
		expectedUser := &models.User{
			FirstName: "Alice",
			LastName:  "Wonder",
			Username:  "alice",
		}

		mockResult := NewMockSingleResult(expectedUser, nil)
		mockCollection.On("FindOne", service.Context, mock.Anything).
			Return(mockResult)

		user, err := services.GetUserByUsername("alice")

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedUser.Username, user.Username)
		assert.Equal(t, expectedUser.FirstName, user.FirstName)
		assert.Equal(t, expectedUser.LastName, user.LastName)
		mockCollection.AssertExpectations(t)
	})*/

	t.Run("UserNotFound", func(t *testing.T) {
		mockResult := NewMockSingleResult(nil, mongo.ErrNoDocuments)
		mockCollection.On("FindOne", service.Context, mock.Anything).
			Return(mockResult)

		user, err := services.GetUserByUsername("nonexistent")

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "user not found", err.Error())
		mockCollection.AssertExpectations(t)
	})

	t.Run("EmptyUsername", func(t *testing.T) {
		user, err := services.GetUserByUsername("")

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "username cannot be empty", err.Error())
	})

	t.Run("DatabaseError", func(t *testing.T) {
		mockResult := NewMockSingleResult(nil, errors.New("database error"))
		mockCollection.On("FindOne", service.Context, mock.Anything).
			Return(mockResult)

		user, err := services.GetUserByUsername("alice")

		assert.Error(t, err)
		assert.Nil(t, user)
		mockCollection.AssertExpectations(t)
	})
}
