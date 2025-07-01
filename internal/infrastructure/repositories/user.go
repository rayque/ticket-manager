package repository

import (
	"context"
	"gorm.io/gorm"
	"shipping-management/internal/domain/entities"
	"shipping-management/internal/infrastructure/repositories/dtos"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user entities.User) (entities.User, error) {
	userDTO := dtos.UserDTO{
		UUID:     user.UUID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Phone:    user.Phone,
		Address:  user.Address,
	}

	if err := ur.db.WithContext(ctx).Create(&userDTO).Error; err != nil {
		return entities.User{}, err
	}

	return entities.User{
		ID:        userDTO.ID,
		UUID:      userDTO.UUID,
		Name:      userDTO.Name,
		Email:     userDTO.Email,
		Password:  userDTO.Password,
		Phone:     userDTO.Phone,
		Address:   userDTO.Address,
		CreatedAt: userDTO.CreatedAt,
		UpdatedAt: userDTO.UpdatedAt,
	}, nil
}

func (ur *UserRepository) GetUserByUUID(ctx context.Context, uuid string) (entities.User, error) {
	var userDTO dtos.UserDTO
	if err := ur.db.WithContext(ctx).Where("uuid = ?", uuid).First(&userDTO).Error; err != nil {
		return entities.User{}, err
	}

	return entities.User{
		ID:        userDTO.ID,
		UUID:      userDTO.UUID,
		Name:      userDTO.Name,
		Email:     userDTO.Email,
		Password:  userDTO.Password,
		Phone:     userDTO.Phone,
		Address:   userDTO.Address,
		CreatedAt: userDTO.CreatedAt,
		UpdatedAt: userDTO.UpdatedAt,
	}, nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (entities.User, error) {
	var userDTO dtos.UserDTO
	if err := ur.db.WithContext(ctx).Where("email = ?", email).First(&userDTO).Error; err != nil {
		return entities.User{}, err
	}

	return entities.User{
		ID:        userDTO.ID,
		UUID:      userDTO.UUID,
		Name:      userDTO.Name,
		Email:     userDTO.Email,
		Password:  userDTO.Password,
		Phone:     userDTO.Phone,
		Address:   userDTO.Address,
		CreatedAt: userDTO.CreatedAt,
		UpdatedAt: userDTO.UpdatedAt,
	}, nil
}

// GetUserByEmailWithPassword é específico para autenticação, incluindo a senha
func (ur *UserRepository) GetUserByEmailWithPassword(ctx context.Context, email string) (entities.User, error) {
	var userDTO dtos.UserDTO
	if err := ur.db.WithContext(ctx).Where("email = ?", email).First(&userDTO).Error; err != nil {
		return entities.User{}, err
	}

	return entities.User{
		ID:        userDTO.ID,
		UUID:      userDTO.UUID,
		Name:      userDTO.Name,
		Email:     userDTO.Email,
		Password:  userDTO.Password,
		Phone:     userDTO.Phone,
		Address:   userDTO.Address,
		CreatedAt: userDTO.CreatedAt,
		UpdatedAt: userDTO.UpdatedAt,
	}, nil
}

func (ur *UserRepository) GetAllUsers(ctx context.Context, limit, offset int) ([]entities.User, error) {
	var userDTOs []dtos.UserDTO
	if err := ur.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&userDTOs).Error; err != nil {
		return nil, err
	}

	users := make([]entities.User, len(userDTOs))
	for i, userDTO := range userDTOs {
		users[i] = entities.User{
			ID:        userDTO.ID,
			UUID:      userDTO.UUID,
			Name:      userDTO.Name,
			Email:     userDTO.Email,
			Phone:     userDTO.Phone,
			Address:   userDTO.Address,
			CreatedAt: userDTO.CreatedAt,
			UpdatedAt: userDTO.UpdatedAt,
		}
	}

	return users, nil
}

func (ur *UserRepository) UpdateUser(ctx context.Context, user entities.User) (entities.User, error) {
	userDTO := dtos.UserDTO{
		ID:      user.ID,
		UUID:    user.UUID,
		Name:    user.Name,
		Email:   user.Email,
		Phone:   user.Phone,
		Address: user.Address,
	}

	if err := ur.db.WithContext(ctx).Save(&userDTO).Error; err != nil {
		return entities.User{}, err
	}

	return entities.User{
		ID:        userDTO.ID,
		UUID:      userDTO.UUID,
		Name:      userDTO.Name,
		Email:     userDTO.Email,
		Phone:     userDTO.Phone,
		Address:   userDTO.Address,
		CreatedAt: user.CreatedAt,
		UpdatedAt: userDTO.UpdatedAt,
	}, nil
}

func (ur *UserRepository) DeleteUser(ctx context.Context, uuid string) error {
	return ur.db.WithContext(ctx).Where("uuid = ?", uuid).Delete(&dtos.UserDTO{}).Error
}
