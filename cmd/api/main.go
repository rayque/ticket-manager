package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"shipping-management/internal/application/usecases"
	"shipping-management/internal/infrastructure/adapters/database"
	"shipping-management/internal/infrastructure/adapters/uuid"
	"shipping-management/internal/infrastructure/config"
	"shipping-management/internal/infrastructure/http"
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
	userRepository := repository.NewUserRepository(postgresDB)

	createUseCase := usecases.NewCreatePackage(packageRepository, uuidAdapter)
	getUseCase := usecases.NewGetPackage(packageRepository)
	updateUseCase := usecases.NewUpdatePackageStatus(packageRepository)
	packageUseCase := usecases.NewPackageQuotationUseCase(packageRepository, carrierRepository)
	hireCarrierUseCase := usecases.NewHireCarrierForPackageDelivery(packageRepository)

	// User use cases
	createUserUseCase := usecases.NewCreateUser(userRepository, uuidAdapter)
	getUserUseCase := usecases.NewGetUser(userRepository)
	getAllUsersUseCase := usecases.NewGetAllUsers(userRepository)
	updateUserUseCase := usecases.NewUpdateUser(userRepository)
	deleteUserUseCase := usecases.NewDeleteUser(userRepository)

	packageHandler := handlers.NewPackageHandler(
		createUseCase,
		getUseCase,
		updateUseCase,
		packageUseCase,
		hireCarrierUseCase,
	)

	userHandler := handlers.NewUserHandler(
		createUserUseCase,
		getUserUseCase,
		getAllUsersUseCase,
		updateUserUseCase,
		deleteUserUseCase,
	)

	http.RegisterRoutes(r, packageHandler, userHandler)
	err = r.Run(":8080")
	if err != nil {
		return
	}
}
