package usecases

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"shipping-management/internal/domain/entities"
	"shipping-management/internal/domain/interfaces"
	"shipping-management/internal/infrastructure/adapters/auth"
)

type LoginUseCaseImpl struct {
	userRepo        interfaces.UserRepository
	jwtService      *auth.JWTService
	passwordService *auth.PasswordService
}

func NewLoginUseCase(
	userRepo interfaces.UserRepository,
	jwtService *auth.JWTService,
	passwordService *auth.PasswordService,
) interfaces.LoginUseCase {
	return &LoginUseCaseImpl{
		userRepo:        userRepo,
		jwtService:      jwtService,
		passwordService: passwordService,
	}
}

func (l *LoginUseCaseImpl) Execute(ctx context.Context, request entities.LoginRequest) (entities.LoginResponse, error) {
	// Buscar usuário por email
	user, err := l.userRepo.GetUserByEmailWithPassword(ctx, request.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.LoginResponse{}, errors.New("credenciais inválidas")
		}
		return entities.LoginResponse{}, err
	}

	// Verificar senha
	if err := l.passwordService.CheckPassword(user.Password, request.Password); err != nil {
		return entities.LoginResponse{}, errors.New("credenciais inválidas")
	}

	// Gerar token JWT
	token, err := l.jwtService.GenerateToken(user)
	if err != nil {
		return entities.LoginResponse{}, errors.New("erro interno do servidor")
	}

	// Limpar senha do usuário antes de retornar
	user.Password = ""

	return entities.LoginResponse{
		Token: token,
		User:  user,
	}, nil
}
