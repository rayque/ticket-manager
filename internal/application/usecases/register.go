package usecases

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"shipping-management/internal/domain/entities"
	"shipping-management/internal/domain/interfaces"
	"shipping-management/internal/infrastructure/adapters/auth"
)

type RegisterUseCaseImpl struct {
	userRepo        interfaces.UserRepository
	passwordService *auth.PasswordService
	uuidGenerator   interfaces.UUID
}

func NewRegisterUseCase(
	userRepo interfaces.UserRepository,
	passwordService *auth.PasswordService,
	uuidGenerator interfaces.UUID,
) interfaces.RegisterUseCase {
	return &RegisterUseCaseImpl{
		userRepo:        userRepo,
		passwordService: passwordService,
		uuidGenerator:   uuidGenerator,
	}
}

func (r *RegisterUseCaseImpl) Execute(ctx context.Context, request entities.RegisterRequest) (entities.User, error) {
	// Verificar se o email já existe
	_, err := r.userRepo.GetUserByEmail(ctx, request.Email)
	if err == nil {
		return entities.User{}, errors.New("email já está em uso")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return entities.User{}, err
	}

	// Hash da senha
	hashedPassword, err := r.passwordService.HashPassword(request.Password)
	if err != nil {
		return entities.User{}, errors.New("erro interno do servidor")
	}

	// Criar usuário
	user := entities.User{
		UUID:     r.uuidGenerator.Generate(),
		Name:     request.Name,
		Email:    request.Email,
		Password: hashedPassword,
		Phone:    request.Phone,
		Address:  request.Address,
	}

	createdUser, err := r.userRepo.CreateUser(ctx, user)
	if err != nil {
		return entities.User{}, err
	}

	// Limpar senha antes de retornar
	createdUser.Password = ""

	return createdUser, nil
}
