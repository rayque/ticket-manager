package usecases

import (
	"context"
	"shipping-management/internal/domain/entities"
	"shipping-management/internal/domain/interfaces"
)

type GetAllUsers struct {
	userRepository interfaces.UserRepository
}

func NewGetAllUsers(userRepository interfaces.UserRepository) *GetAllUsers {
	return &GetAllUsers{
		userRepository: userRepository,
	}
}

func (gau *GetAllUsers) Execute(ctx context.Context, limit, offset int) ([]entities.User, error) {
	if limit <= 0 {
		limit = 10 // valor padrão
	}
	if offset < 0 {
		offset = 0
	}

	return gau.userRepository.GetAllUsers(ctx, limit, offset)
}
