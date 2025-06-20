package usecases

import (
	"context"
	"errors"
	"github.com/stretchr/testify/suite"
	"shipping-management/internal/domain/app_errors"
	"shipping-management/internal/domain/entities"
	"shipping-management/internal/domain/interfaces"
	mocks "shipping-management/mocks/internal_/domain/interfaces"
	"testing"
)

type HireCarrierUseCaseTestSuite struct {
	suite.Suite
	HireCarrierUseCase interfaces.HireCarrierForPackageDeliveryUseCase
	packageRepository  *mocks.PackageRepository
}

func (h *HireCarrierUseCaseTestSuite) SetupTest() {
	h.packageRepository = mocks.NewPackageRepository(h.T())

	h.HireCarrierUseCase = NewHireCarrierForPackageDelivery(
		h.packageRepository,
	)
}
func TestHireCarrierUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(HireCarrierUseCaseTestSuite))
}

func (h *HireCarrierUseCaseTestSuite) TestHireCarrierForPackageDelivery_Execute() {
	ctx := context.Background()
	packageUuid := "12345678-1234-1234-1234-123456789012"
	carrierUuid := "87654321-4321-4321-4321-210987654321"

	inputPackage := entities.Package{
		UUID:        packageUuid,
		Product:     "foo",
		Weight:      10.0,
		Destination: "SP",
		Status:      entities.StatusCreated,
	}

	expectedPackage := entities.Package{
		UUID:        packageUuid,
		Product:     inputPackage.Product,
		Weight:      inputPackage.Weight,
		Destination: inputPackage.Destination,
		Status:      entities.StatusWaitingForCollection,
		CarrierUUID: carrierUuid,
	}

	h.packageRepository.On("GetPackageByUuid", ctx, packageUuid).Return(inputPackage, nil)
	h.packageRepository.On("UpdatePackage", ctx, expectedPackage).Return(expectedPackage, nil)

	result, err := h.HireCarrierUseCase.Execute(ctx, packageUuid, carrierUuid)

	h.NoError(err)
	h.Equal(expectedPackage, result)
	h.Equal(carrierUuid, result.CarrierUUID)
	h.Equal(entities.StatusWaitingForCollection, result.Status)
}

func (h *HireCarrierUseCaseTestSuite) TestHireCarrierForPackageDelivery_Execute_InvalidStatus() {
	ctx := context.Background()
	packageUuid := "12345678-1234-1234-1234-123456789012"
	carrierUuid := "87654321-4321-4321-4321-210987654321"

	inputPackage := entities.Package{
		UUID:        packageUuid,
		Product:     "foo",
		Weight:      10.0,
		Destination: "SP",
		Status:      entities.StatusDelivered, // Invalid status
	}

	h.packageRepository.On("GetPackageByUuid", ctx, packageUuid).Return(inputPackage, nil)

	result, err := h.HireCarrierUseCase.Execute(ctx, packageUuid, carrierUuid)

	h.Error(err)
	h.Equal(entities.Package{}, result)
	h.EqualError(err, "package is not in a valid state to hire a carrier")
}

func (h *HireCarrierUseCaseTestSuite) TestHireCarrierForPackageDelivery_Execute_NotFound() {
	ctx := context.Background()
	packageUuid := "12345678-1234-1234-1234-123456789012"
	carrierUuid := "87654321-4321-4321-4321-210987654321"

	h.packageRepository.On("GetPackageByUuid", ctx, packageUuid).Return(entities.Package{}, app_errors.ErrPackageNotFound)

	result, err := h.HireCarrierUseCase.Execute(ctx, packageUuid, carrierUuid)

	h.Error(err)
	h.Equal(app_errors.ErrPackageNotFound, err)
	h.Equal(entities.Package{}, result)
}

func (h *HireCarrierUseCaseTestSuite) TestHireCarrierForPackageDelivery_Execute_RepositoryError() {
	ctx := context.Background()
	packageUuid := "12345678-1234-1234-1234-123456789012"
	carrierUuid := "87654321-4321-4321-4321-210987654321"

	expectedError := errors.New("repository error")

	h.packageRepository.On("GetPackageByUuid", ctx, packageUuid).Return(entities.Package{}, expectedError)

	result, err := h.HireCarrierUseCase.Execute(ctx, packageUuid, carrierUuid)

	h.Error(err)
	h.Equal(expectedError, err)
	h.Equal(entities.Package{}, result)
}

func (h *HireCarrierUseCaseTestSuite) TestHireCarrierForPackageDelivery_Execute_UpdateError() {
	ctx := context.Background()
	packageUuid := "12345678-1234-1234-1234-123456789012"
	carrierUuid := "87654321-4321-4321-4321-210987654321"

	inputPackage := entities.Package{
		UUID:        packageUuid,
		Product:     "foo",
		Weight:      10.0,
		Destination: "SP",
		Status:      entities.StatusCreated,
	}

	packageToUpdate := entities.Package{
		UUID:        packageUuid,
		Product:     inputPackage.Product,
		Weight:      inputPackage.Weight,
		Destination: inputPackage.Destination,
		Status:      entities.StatusWaitingForCollection,
		CarrierUUID: carrierUuid,
	}

	expectedError := errors.New("repository update error")

	h.packageRepository.On("GetPackageByUuid", ctx, packageUuid).Return(inputPackage, nil)
	h.packageRepository.On("UpdatePackage", ctx, packageToUpdate).Return(entities.Package{}, expectedError)

	result, err := h.HireCarrierUseCase.Execute(ctx, packageUuid, carrierUuid)

	h.Error(err)
	h.Equal(expectedError, err)
	h.Equal(entities.Package{}, result)
}
