package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"shipping-management/internal/infrastructure/adapters/auth"
)

type AuthMiddleware struct {
	jwtService *auth.JWTService
}

func NewAuthMiddleware(jwtService *auth.JWTService) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService: jwtService,
	}
}

func (a *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token de autorização é obrigatório",
			})
			c.Abort()
			return
		}

		// Extrair token do cabeçalho
		token, err := a.jwtService.ExtractTokenFromHeader(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Formato de token inválido",
			})
			c.Abort()
			return
		}

		// Validar token
		claims, err := a.jwtService.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token inválido ou expirado",
			})
			c.Abort()
			return
		}

		// Adicionar informações do usuário ao contexto
		c.Set("user_id", claims.UserID)
		c.Set("user_uuid", claims.UserUUID)
		c.Set("user_email", claims.Email)

		c.Next()
	}
}

func (a *AuthMiddleware) OptionalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.Next()
			return
		}

		token, err := a.jwtService.ExtractTokenFromHeader(authHeader)
		if err != nil {
			c.Next()
			return
		}

		claims, err := a.jwtService.ValidateToken(token)
		if err != nil {
			c.Next()
			return
		}

		// Adicionar informações do usuário ao contexto se o token for válido
		c.Set("user_id", claims.UserID)
		c.Set("user_uuid", claims.UserUUID)
		c.Set("user_email", claims.Email)

		c.Next()
	}
}
