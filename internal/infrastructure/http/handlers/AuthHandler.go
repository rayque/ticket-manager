package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"shipping-management/internal/domain/entities"
	"shipping-management/internal/domain/interfaces"
)

type AuthHandler struct {
	loginUseCase    interfaces.LoginUseCase
	registerUseCase interfaces.RegisterUseCase
	validator       *validator.Validate
}

func NewAuthHandler(
	loginUseCase interfaces.LoginUseCase,
	registerUseCase interfaces.RegisterUseCase,
	validator *validator.Validate,
) *AuthHandler {
	return &AuthHandler{
		loginUseCase:    loginUseCase,
		registerUseCase: registerUseCase,
		validator:       validator,
	}
}

// Login godoc
// @Summary Login de usuário
// @Description Autentica um usuário e retorna um token JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param loginRequest body entities.LoginRequest true "Dados de login"
// @Success 200 {object} entities.LoginResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /auth/login [post]
func (a *AuthHandler) Login(c *gin.Context) {
	var loginRequest entities.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	if err := a.validator.Struct(loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados de validação inválidos",
			"details": err.Error(),
		})
		return
	}

	response, err := a.loginUseCase.Execute(c.Request.Context(), loginRequest)
	if err != nil {
		if err.Error() == "credenciais inválidas" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Email ou senha incorretos",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro interno do servidor",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Register godoc
// @Summary Registro de usuário
// @Description Registra um novo usuário no sistema
// @Tags auth
// @Accept json
// @Produce json
// @Param registerRequest body entities.RegisterRequest true "Dados de registro"
// @Success 201 {object} entities.User
// @Failure 400 {object} map[string]interface{}
// @Failure 409 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /auth/register [post]
func (a *AuthHandler) Register(c *gin.Context) {
	var registerRequest entities.RegisterRequest

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	if err := a.validator.Struct(registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados de validação inválidos",
			"details": err.Error(),
		})
		return
	}

	user, err := a.registerUseCase.Execute(c.Request.Context(), registerRequest)
	if err != nil {
		if err.Error() == "email já está em uso" {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Email já está em uso",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro interno do servidor",
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}
