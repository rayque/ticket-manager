package usecases

import (
	"context"
	"shipping-management/internal/domain/entities"
	"shipping-management/internal/domain/interfaces"
)

type UpdateUser struct {
	userRepository interfaces.UserRepository
}

func NewUpdateUser(userRepository interfaces.UserRepository) *UpdateUser {
	return &UpdateUser{
		userRepository: userRepository,
	}
}

func (uu *UpdateUser) Execute(ctx context.Context, user entities.User) (entities.User, error) {
	// Verificar se o usuário existe
	existingUser, err := uu.userRepository.GetUserByUUID(ctx, user.UUID)
	if err != nil {
		return entities.User{}, err
	}

	// Se o email foi alterado, verificar se não existe outro usuário com o mesmo email
	if user.Email != existingUser.Email {
		userWithEmail, err := uu.userRepository.GetUserByEmail(ctx, user.Email)
		if err == nil && userWithEmail.ID != 0 && userWithEmail.UUID != user.UUID {
			return entities.User{}, &interfaces.DuplicateEmailError{Email: user.Email}
		}
	}

	// Manter dados importantes do usuário existente
	user.ID = existingUser.ID
	user.CreatedAt = existingUser.CreatedAt

	return uu.userRepository.UpdateUser(ctx, user)
}
