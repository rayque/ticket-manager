package usecases

import (
	"context"
	"github.com/stretchr/testify/assert"
	"shipping-management/internal/domain/entities"
	"testing"
)

func TestGetUser_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockUserRepository)
	ctx := context.Background()

	expectedUser := entities.User{
		ID:      1,
		UUID:    "123e4567-e89b-12d3-a456-426614174000",
		Name:    "João Silva",
		Email:   "joao@email.com",
		Phone:   "11999999999",
		Address: "Rua A, 123",
	}

	mockRepo.On("GetUserByUUID", ctx, "123e4567-e89b-12d3-a456-426614174000").Return(expectedUser, nil)

	useCase := NewGetUser(mockRepo)

	// Act
	result, err := useCase.Execute(ctx, "123e4567-e89b-12d3-a456-426614174000")

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)
	mockRepo.AssertExpectations(t)
}

func TestGetUser_NotFound(t *testing.T) {
	// Arrange
	mockRepo := new(MockUserRepository)
	ctx := context.Background()

	mockRepo.On("GetUserByUUID", ctx, "nonexistent-uuid").Return(entities.User{}, assert.AnError)

	useCase := NewGetUser(mockRepo)

	// Act
	result, err := useCase.Execute(ctx, "nonexistent-uuid")

	// Assert
	assert.Error(t, err)
	assert.Equal(t, entities.User{}, result)
	mockRepo.AssertExpectations(t)
}

func TestGetAllUsers_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockUserRepository)
	ctx := context.Background()

	expectedUsers := []entities.User{
		{
			ID:    1,
			UUID:  "uuid1",
			Name:  "João Silva",
			Email: "joao@email.com",
		},
		{
			ID:    2,
			UUID:  "uuid2",
			Name:  "Maria Santos",
			Email: "maria@email.com",
		},
	}

	mockRepo.On("GetAllUsers", ctx, 10, 0).Return(expectedUsers, nil)

	useCase := NewGetAllUsers(mockRepo)

	// Act
	result, err := useCase.Execute(ctx, 10, 0)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, result)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)
}

func TestGetAllUsers_WithDefaultLimit(t *testing.T) {
	// Arrange
	mockRepo := new(MockUserRepository)
	ctx := context.Background()

	expectedUsers := []entities.User{}

	mockRepo.On("GetAllUsers", ctx, 10, 0).Return(expectedUsers, nil)

	useCase := NewGetAllUsers(mockRepo)

	// Act
	result, err := useCase.Execute(ctx, 0, -5) // Test default values

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, result)
	mockRepo.AssertExpectations(t)
}
