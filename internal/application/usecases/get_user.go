package usecases

import (
	"context"
	"shipping-management/internal/domain/entities"
	"shipping-management/internal/domain/interfaces"
)

type GetUser struct {
	userRepository interfaces.UserRepository
}

func NewGetUser(userRepository interfaces.UserRepository) *GetUser {
	return &GetUser{
		userRepository: userRepository,
	}
}

func (gu *GetUser) Execute(ctx context.Context, uuid string) (entities.User, error) {
	return gu.userRepository.GetUserByUUID(ctx, uuid)
}
