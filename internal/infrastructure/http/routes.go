package http

import (
	"github.com/gin-gonic/gin"
	"shipping-management/internal/infrastructure/http/handlers"
)

func RegisterRoutes(r *gin.Engine, packageHandler *handlers.PackageHandler) {
	packageRoutes := r.Group("/package")
	{
		packageRoutes.POST("", packageHandler.CreatePackage)
		packageRoutes.GET(":uuid", packageHandler.GetPackage)
		packageRoutes.PATCH("update/status", packageHandler.UpdatePackageStatus)
		packageRoutes.POST("hire/carrier", packageHandler.HireCarrierForPackageDelivery)
		packageRoutes.GET("quotation/:uuid", packageHandler.GetQuotation)
	}
}
