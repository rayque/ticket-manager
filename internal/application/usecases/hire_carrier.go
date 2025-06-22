package usecases

import (
	"context"
	"shipping-management/internal/domain/app_errors"
	"shipping-management/internal/domain/entities"
	"shipping-management/internal/domain/interfaces"
)

type HireCarrierUseCase struct {
	packageRepository interfaces.PackageRepository
}

func NewHireCarrierForPackageDelivery(
	packageRepository interfaces.PackageRepository,
) interfaces.HireCarrierForPackageDeliveryUseCase {
	return &HireCarrierUseCase{
		packageRepository: packageRepository,
	}
}

func (h HireCarrierUseCase) Execute(ctx context.Context, packageUuid string, carrierUuid string) (entities.Package, error) {
	pkg, err := h.packageRepository.GetPackageByUuid(ctx, packageUuid)
	if err != nil {
		return entities.Package{}, err
	}

	if pkg.Status != entities.StatusCreated {
		return entities.Package{}, app_errors.ErrInvalidPackageStatus
	}

	pkg.CarrierUUID = carrierUuid
	pkg.Status = entities.StatusWaitingForCollection

	updatedPkg, err := h.packageRepository.UpdatePackage(ctx, pkg)
	if err != nil {
		return entities.Package{}, err
	}

	return updatedPkg, nil
}
