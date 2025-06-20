package usecases

import (
	"context"
	"math"
	"shipping-management/internal/domain/entities"
	"shipping-management/internal/domain/interfaces"
)

type PackageQuotationUseCase struct {
	packageRepository interfaces.PackageRepository
	carrierRepository interfaces.CarrierRepository
}

func NewPackageQuotationUseCase(
	packageRepository interfaces.PackageRepository,
	carrierRepository interfaces.CarrierRepository,
) interfaces.PackageQuotationUseCase {
	return &PackageQuotationUseCase{
		packageRepository: packageRepository,
		carrierRepository: carrierRepository,
	}
}

func (p PackageQuotationUseCase) Execute(ctx context.Context, uuid string) ([]entities.Quotation, error) {
	pack, err := p.packageRepository.GetPackageByUuid(ctx, uuid)
	if err != nil {
		return []entities.Quotation{}, err
	}

	region := p.carrierRepository.GetRegionByState(ctx, pack.Destination)
	carriers, err := p.carrierRepository.GetCarriersByRegion(ctx, region)
	if err != nil {
		return []entities.Quotation{}, err
	}

	quotations := make([]entities.Quotation, 0, len(carriers))
	for _, carrier := range carriers {
		regionMap := make(map[string]entities.Region)
		for _, reg := range carrier.Region {
			regionMap[reg.Name] = reg
		}
		if reg, ok := regionMap[region]; ok {
			quotations = append(quotations, entities.Quotation{
				Carrier:           carrier.Name,
				CarrierUUID:       carrier.UUID,
				Price:             getPrice(reg, pack),
				DeliveryTimeByDay: reg.DeliveryTime,
			})
		}
	}

	return quotations, err
}

func getPrice(reg entities.Region, pack entities.Package) float64 {
	return math.Round((reg.PricePerKg*pack.Weight)*100) / 100
}
