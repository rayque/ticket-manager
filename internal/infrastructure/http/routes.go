package http

import (
	"github.com/gin-gonic/gin"
	"shipping-management/internal/infrastructure/http/handlers"
	"shipping-management/internal/infrastructure/http/middleware"
)

func RegisterRoutes(
	r *gin.Engine,
	packageHandler *handlers.PackageHandler,
	userHandler *handlers.UserHandler,
	authHandler *handlers.AuthHandler,
	authMiddleware *middleware.AuthMiddleware,
) {
	// Rotas públicas de autenticação
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}

	// Rotas de pacotes (protegidas)
	packageRoutes := r.Group("/package")
	packageRoutes.Use(authMiddleware.RequireAuth())
	{
		packageRoutes.POST("", packageHandler.CreatePackage)
		packageRoutes.GET(":uuid", packageHandler.GetPackage)
		packageRoutes.PATCH("update/status", packageHandler.UpdatePackageStatus)
		packageRoutes.POST("hire/carrier", packageHandler.HireCarrierForPackageDelivery)
		packageRoutes.GET("quotation/:uuid", packageHandler.GetQuotation)
	}

	// Rotas de usuários (protegidas)
	userRoutes := r.Group("/users")
	userRoutes.Use(authMiddleware.RequireAuth())
	{
		userRoutes.POST("", userHandler.CreateUser)
		userRoutes.GET("", userHandler.GetAllUsers)
		userRoutes.GET("/:uuid", userHandler.GetUser)
		userRoutes.PUT("/:uuid", userHandler.UpdateUser)
		userRoutes.DELETE("/:uuid", userHandler.DeleteUser)
	}
}
