package usecases

import (
	"context"
	"shipping-management/internal/domain/entities"
	"shipping-management/internal/domain/interfaces"
)

type CreateUser struct {
	userRepository interfaces.UserRepository
	uuidGenerator  interfaces.UUID
}

func NewCreateUser(userRepository interfaces.UserRepository, uuidGenerator interfaces.UUID) *CreateUser {
	return &CreateUser{
		userRepository: userRepository,
		uuidGenerator:  uuidGenerator,
	}
}

func (cu *CreateUser) Execute(ctx context.Context, user entities.User) (entities.User, error) {
	user.UUID = cu.uuidGenerator.Generate()

	// Verificar se email já existe
	existingUser, err := cu.userRepository.GetUserByEmail(ctx, user.Email)
	if err == nil && existingUser.ID != 0 {
		return entities.User{}, &interfaces.DuplicateEmailError{Email: user.Email}
	}

	return cu.userRepository.CreateUser(ctx, user)
}
