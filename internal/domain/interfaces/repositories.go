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
