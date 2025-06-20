package usecases

import (
	"context"
	"github.com/stretchr/testify/suite"
	"shipping-management/internal/domain/app_errors"
	"shipping-management/internal/domain/entities"
	"shipping-management/internal/domain/interfaces"
	mocks "shipping-management/mocks/internal_/domain/interfaces"
	"testing"
)

type PackageQuotationUseCaseTestSuite struct {
	suite.Suite
	PackageQuotationUseCase interfaces.PackageQuotationUseCase
	packageRepository       *mocks.PackageRepository
	carrierRepository       *mocks.CarrierRepository
}

func (p *PackageQuotationUseCaseTestSuite) SetupTest() {
	p.packageRepository = mocks.NewPackageRepository(p.T())
	p.carrierRepository = mocks.NewCarrierRepository(p.T())

	p.PackageQuotationUseCase = NewPackageQuotationUseCase(
		p.packageRepository,
		p.carrierRepository,
	)
}
func TestPackageQuotationUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(PackageQuotationUseCaseTestSuite))
}

func (p *PackageQuotationUseCaseTestSuite) TestPackageQuotationUseCase_Execute() {
	ctx := context.Background()
	uuid := "12345678-1234-1234-1234-123456789012"

	expectedPackage := entities.Package{
		UUID:        uuid,
		Product:     "Test Product",
		Weight:      10.0,
		Destination: "SP",
	}

	p.packageRepository.On("GetPackageByUuid", ctx, uuid).Return(expectedPackage, nil)

	expectedRegion := "SP"
	p.carrierRepository.On("GetRegionByState", ctx, expectedPackage.Destination).Return(expectedRegion)

	expectedCarriers := []entities.Carrier{
		{
			Name: "Carrier A",
			UUID: "carrier-a-uuid",
			Region: []entities.Region{
				{Name: expectedRegion, PricePerKg: 10.0, DeliveryTime: 2},
			},
		},
	}

	p.carrierRepository.On("GetCarriersByRegion", ctx, expectedRegion).Return(expectedCarriers, nil)

	expectedQuotation := []entities.Quotation{
		{
			Carrier:           "Carrier A",
			CarrierUUID:       "carrier-a-uuid",
			Price:             100.0,
			DeliveryTimeByDay: 2,
		},
	}

	result, err := p.PackageQuotationUseCase.Execute(ctx, uuid)

	p.NoError(err)
	p.Equal(expectedQuotation, result)
}

func (p *PackageQuotationUseCaseTestSuite) TestPackageQuotationUseCase_Execute_PackageNotFound() {
	ctx := context.Background()
	uuid := "12345678-1234-1234-1234-123456789012"

	p.packageRepository.On("GetPackageByUuid", ctx, uuid).Return(entities.Package{}, app_errors.ErrPackageNotFound)

	result, err := p.PackageQuotationUseCase.Execute(ctx, uuid)

	p.Error(err)
	p.Equal(app_errors.ErrPackageNotFound, err)
	p.Empty(result)
}

func (p *PackageQuotationUseCaseTestSuite) TestPackageQuotationUseCase_Execute_CarrierRepositoryError() {
	ctx := context.Background()
	uuid := "12345678-1234-1234-1234-123456789012"

	expectedPackage := entities.Package{
		UUID:        uuid,
		Product:     "Test Product",
		Weight:      10.0,
		Destination: "SP",
	}

	p.packageRepository.On("GetPackageByUuid", ctx, uuid).Return(expectedPackage, nil)

	expectedRegion := "SP"
	p.carrierRepository.On("GetRegionByState", ctx, expectedPackage.Destination).Return(expectedRegion)

	mockError := app_errors.ErrNoCarrierFound
	p.carrierRepository.On("GetCarriersByRegion", ctx, expectedRegion).Return(nil, mockError)

	result, err := p.PackageQuotationUseCase.Execute(ctx, uuid)

	p.Equal(mockError, err)
	p.Empty(result)
}
