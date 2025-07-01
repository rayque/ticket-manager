package usecases

import (
	"context"
	"shipping-management/internal/domain/interfaces"
)

type DeleteUser struct {
	userRepository interfaces.UserRepository
}

func NewDeleteUser(userRepository interfaces.UserRepository) *DeleteUser {
	return &DeleteUser{
		userRepository: userRepository,
	}
}

func (du *DeleteUser) Execute(ctx context.Context, uuid string) error {
	// Verificar se o usuário existe antes de deletar
	_, err := du.userRepository.GetUserByUUID(ctx, uuid)
	if err != nil {
		return err
	}

	return du.userRepository.DeleteUser(ctx, uuid)
}
