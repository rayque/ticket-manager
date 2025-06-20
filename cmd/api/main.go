package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"shipping-management/internal/application/usecases"
	"shipping-management/internal/infrastructure/adapters/database"
	"shipping-management/internal/infrastructure/adapters/uuid"
	"shipping-management/internal/infrastructure/config"
	"shipping-management/internal/infrastructure/http/handlers"
	repository "shipping-management/internal/infrastructure/repositories"
)

func main() {
	ctx := context.Background()
	r := gin.Default()

	appConfig := config.NewConfig()

	mongoDB, err := database.NewMongoDatabase(ctx, appConfig)
	if err != nil {
		panic("failed to connect to MongoDB: " + err.Error())
	}

	postgresDB, err := database.NewPostgres(appConfig)
	if err != nil {
		panic("failed to connect to postgres database: " + err.Error())
	}

	uuidAdapter := uuid.NewUUIDAdapter()

	packageRepository := repository.NewPackageRepository(postgresDB)
	carrierRepository := repository.NewCarrierRepository(mongoDB)

	createUseCase := usecases.NewCreatePackage(packageRepository, uuidAdapter)
	getUseCase := usecases.NewGetPackage(packageRepository)
	updateUseCase := usecases.NewUpdatePackageStatus(packageRepository)
	packageUseCase := usecases.NewPackageQuotationUseCase(packageRepository, carrierRepository)
	hireCarrierUseCase := usecases.NewHireCarrierForPackageDelivery(packageRepository)

	packageHandler := handlers.NewPackageHandler(
		createUseCase,
		getUseCase,
		updateUseCase,
		packageUseCase,
		hireCarrierUseCase,
	)

	r.POST("/create", packageHandler.CreatePackage)
	r.GET("/get/:uuid", packageHandler.GetPackage)
	r.PATCH("/update/status", packageHandler.UpdatePackageStatus)
	r.GET("/quotation/:uuid", packageHandler.GetQuotation)
	r.POST("/hire/carrier", packageHandler.HireCarrierForPackageDelivery)
	//http.RegisterRoutes(r)
	err = r.Run(":8080")
	if err != nil {
		return
	}
}
