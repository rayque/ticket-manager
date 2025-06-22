package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"shipping-management/internal/domain/app_errors"
	"shipping-management/internal/domain/entities"
	"shipping-management/internal/domain/interfaces"
	"shipping-management/internal/infrastructure/validator"
)

type PackageHandler struct {
	createUseCase      interfaces.CreatePackageUseCase
	getUseCase         interfaces.GetPackageUseCase
	updateUseCase      interfaces.UpdatePackageStatusUseCase
	packageUseCase     interfaces.PackageQuotationUseCase
	hireCarrierUseCase interfaces.HireCarrierForPackageDeliveryUseCase
}

func NewPackageHandler(
	createUseCase interfaces.CreatePackageUseCase,
	getUseCase interfaces.GetPackageUseCase,
	updateUseCase interfaces.UpdatePackageStatusUseCase,
	packageUseCase interfaces.PackageQuotationUseCase,
	hireCarrierUseCase interfaces.HireCarrierForPackageDeliveryUseCase,
) *PackageHandler {
	return &PackageHandler{
		createUseCase:      createUseCase,
		getUseCase:         getUseCase,
		updateUseCase:      updateUseCase,
		packageUseCase:     packageUseCase,
		hireCarrierUseCase: hireCarrierUseCase,
	}
}

func (p *PackageHandler) CreatePackage(c *gin.Context) {
	var req struct {
		Product     string  `json:"product" validate:"required"`
		Weight      float64 `json:"weight" validate:"required"`
		Destination string  `json:"destination" validate:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Erro na validação do JSON",
			"details": err.Error(),
		})

		return
	}

	errValidator := validator.ValidateRequest(req)
	if errValidator != nil {
		c.JSON(http.StatusBadRequest, errValidator)
		return
	}

	pack := entities.Package{
		Product:     req.Product,
		Weight:      req.Weight,
		Destination: req.Destination,
	}

	createdPackage, err := p.createUseCase.Execute(c.Request.Context(), pack)
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create package"})
		return
	}

	c.JSON(http.StatusCreated, createdPackage)
}

func (p *PackageHandler) GetPackage(c *gin.Context) {
	packageUuid := c.Param("uuid")
	if packageUuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Package UUID is required"})
		return
	}
	pkg, err := p.getUseCase.Execute(c.Request.Context(), packageUuid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Package not found"})
		return
	}
	c.JSON(http.StatusOK, pkg)
}

func (p *PackageHandler) UpdatePackageStatus(c *gin.Context) {
	var req struct {
		UUID   string          `json:"uuid" validate:"required"`
		Status entities.Status `json:"status" validate:"required,oneof=CREATED WAITING_FOR_COLLECTION COLLECTED SENT DELIVERED LOST"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid parameters",
			"details": err.Error(),
		})
		return
	}

	errValidator := validator.ValidateRequest(req)
	if errValidator != nil {
		c.JSON(http.StatusBadRequest, errValidator)
		return
	}

	updatedPackage, err := p.updateUseCase.Execute(c.Request.Context(), req.UUID, req.Status)
	if err != nil {
		if errors.Is(err, app_errors.ErrPackageNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update package status"})
		return
	}

	c.JSON(http.StatusOK, updatedPackage)
}

func (p *PackageHandler) GetQuotation(c *gin.Context) {
	packageUuid := c.Param("uuid")
	if packageUuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Package UUID is required"})
		return
	}

	quotations, err := p.packageUseCase.Execute(c.Request.Context(), packageUuid)
	if err != nil {
		if errors.Is(err, app_errors.ErrPackageNotFound) || errors.Is(err, app_errors.ErrNoCarrierFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get package quotations"})
		return
	}

	c.JSON(http.StatusOK, quotations)
}

func (p *PackageHandler) HireCarrierForPackageDelivery(c *gin.Context) {
	var req struct {
		PackageUuid string `json:"package_uuid" validate:"required"`
		CarrierUuid string `json:"carrier_uuid" validate:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid parameters",
			"details": err.Error(),
		})
		return
	}

	errValidator := validator.ValidateRequest(req)
	if errValidator != nil {
		c.JSON(http.StatusBadRequest, errValidator)
		return
	}
	pkg, err := p.hireCarrierUseCase.Execute(c.Request.Context(), req.PackageUuid, req.CarrierUuid)
	if err != nil {
		if errors.Is(err, app_errors.ErrPackageNotFound) || errors.Is(err, app_errors.ErrInvalidPackageStatus) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hire carrier for package delivery"})
		return
	}

	c.JSON(http.StatusOK, pkg)
}
