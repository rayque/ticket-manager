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

type UpdatePackageStatusUseCaseTestSuite struct {
	suite.Suite
	PackageStatusUseCase interfaces.UpdatePackageStatusUseCase
	packageRepository    *mocks.PackageRepository
}

func (p *UpdatePackageStatusUseCaseTestSuite) SetupTest() {
	p.packageRepository = mocks.NewPackageRepository(p.T())
	p.PackageStatusUseCase = NewUpdatePackageStatus(p.packageRepository)
}

func TestUpdatePackageStatusUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UpdatePackageStatusUseCaseTestSuite))
}

func (p *UpdatePackageStatusUseCaseTestSuite) TestUpdatePackageStatusUseCase_Execute() {
	ctx := context.Background()
	uuid := "12345678-1234-1234-1234-123456789012"
	status := entities.StatusDelivered

	expectedPackage := entities.Package{
		UUID:   uuid,
		Status: status,
	}

	p.packageRepository.On("UpdatePackage", ctx, expectedPackage).Return(expectedPackage, nil)

	result, err := p.PackageStatusUseCase.Execute(ctx, uuid, status)
	p.NoError(err)
	p.Equal(expectedPackage, result)

	p.packageRepository.AssertExpectations(p.T())
}

func (p *UpdatePackageStatusUseCaseTestSuite) TestUpdatePackageStatusUseCase_Execute_Error() {
	ctx := context.Background()
	uuid := "12345678-1234-1234-1234-123456789012"
	status := entities.StatusDelivered

	expectedError := app_errors.ErrPackageNotFound

	p.packageRepository.On("UpdatePackage", ctx, entities.Package{UUID: uuid, Status: status}).
		Return(entities.Package{}, expectedError)

	result, err := p.PackageStatusUseCase.Execute(ctx, uuid, status)
	p.Equal(expectedError, err)
	p.Equal(entities.Package{}, result)

	p.packageRepository.AssertExpectations(p.T())
}
