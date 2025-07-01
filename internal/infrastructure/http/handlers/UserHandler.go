package handlers

import (
	"net/http"
	"shipping-management/internal/domain/entities"
	"shipping-management/internal/domain/interfaces"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	createUserUseCase  interfaces.CreateUserUseCase
	getUserUseCase     interfaces.GetUserUseCase
	getAllUsersUseCase interfaces.GetAllUsersUseCase
	updateUserUseCase  interfaces.UpdateUserUseCase
	deleteUserUseCase  interfaces.DeleteUserUseCase
}

func NewUserHandler(
	createUserUseCase interfaces.CreateUserUseCase,
	getUserUseCase interfaces.GetUserUseCase,
	getAllUsersUseCase interfaces.GetAllUsersUseCase,
	updateUserUseCase interfaces.UpdateUserUseCase,
	deleteUserUseCase interfaces.DeleteUserUseCase,
) *UserHandler {
	return &UserHandler{
		createUserUseCase:  createUserUseCase,
		getUserUseCase:     getUserUseCase,
		getAllUsersUseCase: getAllUsersUseCase,
		updateUserUseCase:  updateUserUseCase,
		deleteUserUseCase:  deleteUserUseCase,
	}
}

type CreateUserRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type UpdateUserRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entities.User{
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
	}

	createdUser, err := uh.createUserUseCase.Execute(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func (uh *UserHandler) GetUser(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UUID is required"})
		return
	}

	user, err := uh.getUserUseCase.Execute(c.Request.Context(), uuid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) GetAllUsers(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
		return
	}

	users, err := uh.getAllUsersUseCase.Execute(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users":  users,
		"limit":  limit,
		"offset": offset,
	})
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UUID is required"})
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entities.User{
		UUID:    uuid,
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
	}

	updatedUser, err := uh.updateUserUseCase.Execute(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UUID is required"})
		return
	}

	err := uh.deleteUserUseCase.Execute(c.Request.Context(), uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
