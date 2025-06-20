package usecases

import (
	"context"
	"shipping-management/internal/domain/entities"
	"shipping-management/internal/domain/interfaces"
)

type CreatePackageUseCase struct {
	packageRepository interfaces.PackageRepository
	uuidCreator       interfaces.UUID
}

func NewCreatePackage(
	packageRepository interfaces.PackageRepository,
	uuidCreator interfaces.UUID,
) interfaces.CreatePackageUseCase {
	return &CreatePackageUseCase{
		packageRepository: packageRepository,
		uuidCreator:       uuidCreator,
	}
}

func (c *CreatePackageUseCase) Execute(ctx context.Context, pack entities.Package) (entities.Package, error) {
	pack.UUID = c.uuidCreator.NewUUID()
	return c.packageRepository.CreatePackage(ctx, pack)
}
