package usecases

import (
	"context"
	"shipping-management/internal/domain/entities"
	"shipping-management/internal/domain/interfaces"
)

type UpdatePackageStatusUseCase struct {
	packageRepository interfaces.PackageRepository
}

func NewUpdatePackageStatus(packageRepository interfaces.PackageRepository) interfaces.UpdatePackageStatusUseCase {
	return &UpdatePackageStatusUseCase{
		packageRepository: packageRepository,
	}
}

func (u UpdatePackageStatusUseCase) Execute(ctx context.Context, uuid string, status entities.Status) (entities.Package, error) {
	pkg := entities.Package{
		UUID:   uuid,
		Status: status,
	}
	return u.packageRepository.UpdatePackage(ctx, pkg)
}
