package interfaces

import (
	"context"
	"shipping-management/internal/domain/entities"
)

type CreatePackageUseCase interface {
	Execute(ctx context.Context, pack entities.Package) (entities.Package, error)
}

type GetPackageUseCase interface {
	Execute(ctx context.Context, uuid string) (entities.Package, error)
}

type UpdatePackageStatusUseCase interface {
	Execute(ctx context.Context, uuid string, status entities.Status) (entities.Package, error)
}

type HireCarrierForPackageDeliveryUseCase interface {
	Execute(ctx context.Context, packageUuid string, carrierUuid string) (entities.Package, error)
}

type PackageQuotationUseCase interface {
	Execute(ctx context.Context, packageUuid string) ([]entities.Quotation, error)
}

// Interfaces de autenticação
type LoginUseCase interface {
	Execute(ctx context.Context, request entities.LoginRequest) (entities.LoginResponse, error)
}

type RegisterUseCase interface {
	Execute(ctx context.Context, request entities.RegisterRequest) (entities.User, error)
}

// Interfaces de usuário
type CreateUserUseCase interface {
	Execute(ctx context.Context, user entities.User) (entities.User, error)
}

type GetUserUseCase interface {
	Execute(ctx context.Context, uuid string) (entities.User, error)
}

type GetAllUsersUseCase interface {
	Execute(ctx context.Context, limit, offset int) ([]entities.User, error)
}

type UpdateUserUseCase interface {
	Execute(ctx context.Context, user entities.User) (entities.User, error)
}

type DeleteUserUseCase interface {
	Execute(ctx context.Context, uuid string) error
}
