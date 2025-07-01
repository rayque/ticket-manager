package http

import (
	"github.com/gin-gonic/gin"
	"shipping-management/internal/infrastructure/http/handlers"
)

func RegisterRoutes(r *gin.Engine, packageHandler *handlers.PackageHandler, userHandler *handlers.UserHandler) {
	packageRoutes := r.Group("/package")
	{
		packageRoutes.POST("", packageHandler.CreatePackage)
		packageRoutes.GET(":uuid", packageHandler.GetPackage)
		packageRoutes.PATCH("update/status", packageHandler.UpdatePackageStatus)
		packageRoutes.POST("hire/carrier", packageHandler.HireCarrierForPackageDelivery)
		packageRoutes.GET("quotation/:uuid", packageHandler.GetQuotation)
	}

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", userHandler.CreateUser)
		userRoutes.GET("", userHandler.GetAllUsers)
		userRoutes.GET("/:uuid", userHandler.GetUser)
		userRoutes.PUT("/:uuid", userHandler.UpdateUser)
		userRoutes.DELETE("/:uuid", userHandler.DeleteUser)
	}
}
