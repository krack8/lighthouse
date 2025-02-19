package services

import (
	"context"
	db "github.com/krack8/lighthouse/pkg/controller/auth/config"
	"github.com/krack8/lighthouse/pkg/controller/auth/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MockSingleResult is a mock implementation of mongo.SingleResult
type MockSingleResult struct {
	err    error
	result interface{}
}

func (m *MockSingleResult) Err() error {
	return m.err
}

func (m *MockSingleResult) Decode(v interface{}) error {
	if m.err != nil {
		return m.err
	}

	// Type assert and copy the mock result to v
	if user, ok := m.result.(*models.User); ok {
		*(v.(*models.User)) = *user
	}
	return nil
}

// MockCollection is a mock implementation of the Collection interface
type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockCollection) FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult {
	args := m.Called(ctx, filter)
	if args.Get(0) == nil {
		return &mongo.SingleResult{}
	}
	return args.Get(0).(*mongo.SingleResult)
}

func (m *MockCollection) Find(ctx context.Context, filter interface{}) (*mongo.Cursor, error) {
	args := m.Called(ctx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mongo.Cursor), args.Error(1)
}

func (m *MockCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func (m *MockCollection) DeleteOne(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	args := m.Called(ctx, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mongo.DeleteResult), args.Error(1)
}

// MockCursor is a mock implementation of the mongo.Cursor
type MockCursor struct {
	mock.Mock
}

func (m *MockCursor) Next(ctx context.Context) bool {
	args := m.Called(ctx)
	return args.Bool(0)
}

func (m *MockCursor) Decode(v interface{}) error {
	args := m.Called(v)
	if user, ok := v.(*models.User); ok {
		user.FirstName = "John"
		user.LastName = "Doe"
	}
	return args.Error(0)
}

func (m *MockCursor) Close(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockCursor) Err() error {
	args := m.Called()
	return args.Error(0)
}

func TestNewUserService(t *testing.T) {
	mockColl := new(MockCollection)
	service := NewUserService(mockColl)
	assert.NotNil(t, service)
	assert.Equal(t, mockColl, service.collection)
}

func TestCreateUser(t *testing.T) {
	mockColl := new(MockCollection)
	service := NewUserService(mockColl)

	t.Run("nil user", func(t *testing.T) {
		result, err := service.CreateUser(nil)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "user cannot be nil", err.Error())
	})
}

func TestGetUser(t *testing.T) {
	mockColl := new(MockCollection)
	service := NewUserService(mockColl)

	t.Run("invalid user ID", func(t *testing.T) {
		result, err := service.GetUser("invalid-id")
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "invalid user ID format")
	})

	t.Run("empty user ID", func(t *testing.T) {
		result, err := service.GetUser("")
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "user ID cannot be empty", err.Error())
	})

}

func TestUpdateUser(t *testing.T) {
	mockColl := new(MockCollection)
	service := NewUserService(mockColl)

	t.Run("nil user", func(t *testing.T) {
		err := service.UpdateUser(primitive.NewObjectID().Hex(), nil)
		assert.Error(t, err)
		assert.Equal(t, "updated user cannot be nil", err.Error())
	})

	t.Run("invalid user ID", func(t *testing.T) {
		err := service.UpdateUser("invalid-id", &models.User{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid user ID format")
	})

}

func TestDeleteUser(t *testing.T) {
	mockColl := new(MockCollection)
	service := NewUserService(mockColl)

	t.Run("invalid user ID", func(t *testing.T) {
		err := service.DeleteUser("invalid-id")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid user ID format")
	})

}

func TestGetUserByUsername(t *testing.T) {
	// Setup
	originalCollection := db.UserCollection

	// Teardown
	defer func() {
		db.UserCollection = originalCollection
	}()

	t.Run("empty username", func(t *testing.T) {
		result, err := GetUserByUsername("")
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "username cannot be empty", err.Error())
	})
}
