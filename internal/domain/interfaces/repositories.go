package interfaces

import (
	"context"
	"shipping-management/internal/domain/entities"
)

type PackageRepository interface {
	CreatePackage(ctx context.Context, product entities.Package) (entities.Package, error)
	GetPackageByUuid(ctx context.Context, uuid string) (entities.Package, error)
	UpdatePackage(ctx context.Context, pkg entities.Package) (entities.Package, error)
}

type CarrierRepository interface {
	GetRegionByState(state string) string
	GetCarriersByRegion(ctx context.Context, uuid string) ([]entities.Carrier, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, user entities.User) (entities.User, error)
	GetUserByUUID(ctx context.Context, uuid string) (entities.User, error)
	GetUserByEmail(ctx context.Context, email string) (entities.User, error)
	GetAllUsers(ctx context.Context, limit, offset int) ([]entities.User, error)
	UpdateUser(ctx context.Context, user entities.User) (entities.User, error)
	DeleteUser(ctx context.Context, uuid string) error
}
