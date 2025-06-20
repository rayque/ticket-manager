package usecases

import (
	"context"
	"shipping-management/internal/domain/entities"
	"shipping-management/internal/domain/interfaces"
)

type GetPackageUseCase struct {
	packageRepository interfaces.PackageRepository
}

func NewGetPackage(packageRepository interfaces.PackageRepository) interfaces.GetPackageUseCase {
	return &GetPackageUseCase{
		packageRepository: packageRepository,
	}
}

func (g GetPackageUseCase) Execute(ctx context.Context, uuid string) (entities.Package, error) {
	return g.packageRepository.GetPackageByUuid(ctx, uuid)
}
