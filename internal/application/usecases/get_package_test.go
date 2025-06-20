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

type GetPackageUseCaseTestSuite struct {
	suite.Suite
	GetPackageUseCase interfaces.GetPackageUseCase
	packageRepository *mocks.PackageRepository
}

func (g *GetPackageUseCaseTestSuite) SetupTest() {
	g.packageRepository = mocks.NewPackageRepository(g.T())

	g.GetPackageUseCase = NewGetPackage(
		g.packageRepository,
	)
}

func TestGetPackageUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(GetPackageUseCaseTestSuite))
}

func (g *GetPackageUseCaseTestSuite) TestGetPackage_Execute() {
	ctx := context.Background()
	uuid := "12345678-1234-1234-1234-123456789012"

	expectedPackage := entities.Package{
		UUID:        uuid,
		Product:     "foo",
		Weight:      10.0,
		Destination: "SP",
	}

	g.packageRepository.On("GetPackageByUuid", ctx, uuid).Return(expectedPackage, nil)

	result, err := g.GetPackageUseCase.Execute(ctx, uuid)

	g.NoError(err)
	g.Equal(expectedPackage, result)
	g.Equal(uuid, result.UUID)
	g.Equal(expectedPackage.Product, result.Product)
	g.Equal(expectedPackage.Weight, result.Weight)
	g.Equal(expectedPackage.Destination, result.Destination)
}

func (g *GetPackageUseCaseTestSuite) TestGetPackage_Execute_NotFound() {
	ctx := context.Background()
	uuid := "12345678-1234-1234-1234-123456789012"

	g.packageRepository.On("GetPackageByUuid", ctx, uuid).Return(entities.Package{}, app_errors.ErrPackageNotFound)

	result, err := g.GetPackageUseCase.Execute(ctx, uuid)

	g.Error(err)
	g.Equal(app_errors.ErrPackageNotFound, err)
	g.Empty(result)
}
