package usecases

import (
	"context"
	"errors"
	"github.com/stretchr/testify/suite"
	"shipping-management/internal/domain/entities"
	"shipping-management/internal/domain/interfaces"
	mocks "shipping-management/mocks/internal_/domain/interfaces"
	"testing"
)

type CreatePackageUseCaseTestSuite struct {
	suite.Suite
	CreatePackageUseCase interfaces.CreatePackageUseCase
	packageRepository    *mocks.PackageRepository
	uuidCreator          *mocks.UUID
}

func (c *CreatePackageUseCaseTestSuite) SetupTest() {
	c.packageRepository = mocks.NewPackageRepository(c.T())
	c.uuidCreator = mocks.NewUUID(c.T())

	c.CreatePackageUseCase = NewCreatePackage(
		c.packageRepository,
		c.uuidCreator,
	)
}

func TestCreatePackageUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CreatePackageUseCaseTestSuite))
}

func (c *CreatePackageUseCaseTestSuite) TestCreatePackage_Execute() {
	ctx := context.Background()
	expectedUUID := "12345678-1234-1234-1234-123456789012"

	inputPackage := entities.Package{
		UUID:        expectedUUID,
		Product:     "foo",
		Weight:      10.0,
		Destination: "SP",
	}

	expectedPackage := entities.Package{
		UUID:        expectedUUID,
		Product:     inputPackage.Product,
		Weight:      inputPackage.Weight,
		Destination: inputPackage.Destination,
	}

	c.uuidCreator.On("NewUUID").Return(expectedUUID)
	c.packageRepository.On("CreatePackage", ctx, inputPackage).Return(expectedPackage, nil)

	result, err := c.CreatePackageUseCase.Execute(ctx, inputPackage)

	c.NoError(err)
	c.Equal(expectedPackage, result)
	c.Equal(expectedUUID, result.UUID)
	c.Equal(inputPackage.Product, result.Product)
	c.Equal(inputPackage.Weight, result.Weight)
	c.Equal(inputPackage.Destination, result.Destination)

	c.packageRepository.AssertExpectations(c.T())
	c.uuidCreator.AssertExpectations(c.T())
}

func (c *CreatePackageUseCaseTestSuite) TestCreatePackage_Execute_RepositoryError() {
	ctx := context.Background()
	expectedUUID := "12345678-1234-1234-1234-123456789012"
	expectedError := errors.New("repository error")

	inputPackage := entities.Package{}

	c.uuidCreator.On("NewUUID").Return(expectedUUID)
	c.packageRepository.On("CreatePackage", ctx, entities.Package{
		UUID: expectedUUID,
	}).Return(entities.Package{}, expectedError)

	result, err := c.CreatePackageUseCase.Execute(ctx, inputPackage)

	c.Error(err)
	c.Equal(expectedError, err)
	c.Equal(entities.Package{}, result)

	c.packageRepository.AssertExpectations(c.T())
	c.uuidCreator.AssertExpectations(c.T())
}
