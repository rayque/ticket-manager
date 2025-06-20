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
