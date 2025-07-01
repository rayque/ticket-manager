package usecases

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"shipping-management/internal/domain/entities"
	"shipping-management/internal/domain/interfaces"
	"testing"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user entities.User) (entities.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByUUID(ctx context.Context, uuid string) (entities.User, error) {
	args := m.Called(ctx, uuid)
	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByEmail(ctx context.Context, email string) (entities.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockUserRepository) GetAllUsers(ctx context.Context, limit, offset int) ([]entities.User, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).([]entities.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, user entities.User) (entities.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(entities.User), args.Error(1)
}

func (m *MockUserRepository) DeleteUser(ctx context.Context, uuid string) error {
	args := m.Called(ctx, uuid)
	return args.Error(0)
}

type MockUUID struct {
	mock.Mock
}

func (m *MockUUID) NewUUID() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockUUID) Generate() string {
	args := m.Called()
	return args.String(0)
}

func TestCreateUser_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockUserRepository)
	mockUUID := new(MockUUID)
	ctx := context.Background()

	inputUser := entities.User{
		Name:    "João Silva",
		Email:   "joao@email.com",
		Phone:   "11999999999",
		Address: "Rua A, 123",
	}

	expectedUUID := "123e4567-e89b-12d3-a456-426614174000"
	expectedUser := entities.User{
		ID:      1,
		UUID:    expectedUUID,
		Name:    "João Silva",
		Email:   "joao@email.com",
		Phone:   "11999999999",
		Address: "Rua A, 123",
	}

	mockUUID.On("Generate").Return(expectedUUID)
	mockRepo.On("GetUserByEmail", ctx, "joao@email.com").Return(entities.User{}, assert.AnError)
	mockRepo.On("CreateUser", ctx, mock.MatchedBy(func(user entities.User) bool {
		return user.UUID == expectedUUID && user.Email == "joao@email.com"
	})).Return(expectedUser, nil)

	useCase := NewCreateUser(mockRepo, mockUUID)

	// Act
	result, err := useCase.Execute(ctx, inputUser)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)
	mockRepo.AssertExpectations(t)
	mockUUID.AssertExpectations(t)
}

func TestCreateUser_DuplicateEmail(t *testing.T) {
	// Arrange
	mockRepo := new(MockUserRepository)
	mockUUID := new(MockUUID)
	ctx := context.Background()

	inputUser := entities.User{
		Name:  "João Silva",
		Email: "joao@email.com",
	}

	existingUser := entities.User{
		ID:    1,
		Email: "joao@email.com",
	}

	expectedUUID := "123e4567-e89b-12d3-a456-426614174000"

	mockUUID.On("Generate").Return(expectedUUID)
	mockRepo.On("GetUserByEmail", ctx, "joao@email.com").Return(existingUser, nil)

	useCase := NewCreateUser(mockRepo, mockUUID)

	// Act
	result, err := useCase.Execute(ctx, inputUser)

	// Assert
	assert.Error(t, err)
	assert.IsType(t, &interfaces.DuplicateEmailError{}, err)
	assert.Equal(t, entities.User{}, result)
	mockRepo.AssertExpectations(t)
	mockUUID.AssertExpectations(t)
}
