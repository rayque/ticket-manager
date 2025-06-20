package dtos

import (
	"shipping-management/internal/domain/entities"
	"time"
)

type PackageDto struct {
	ID          int64   `gorm:"column:id;primaryKey"`
	UUID        string  `gorm:"column:uuid"`
	Product     string  `gorm:"column:product"`
	Weight      float64 `gorm:"column:weight"`
	Destination string  `gorm:"column:destination"`
	Status      string  `gorm:"column:status"`
	CarrierUUID string  `gorm:"column:carrier_uuid"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (PackageDto) TableName() string {
	return "packages"
}
func (p PackageDto) ToEntity() entities.Package {
	return entities.Package{
		ID:          p.ID,
		UUID:        p.UUID,
		Product:     p.Product,
		Weight:      p.Weight,
		Destination: p.Destination,
		Status:      entities.Status(p.Status),
		CarrierUUID: p.CarrierUUID,
	}
}

func FromEntityToPackageDto(entity entities.Package) PackageDto {
	return PackageDto{
		ID:          entity.ID,
		UUID:        entity.UUID,
		Product:     entity.Product,
		Weight:      entity.Weight,
		Destination: entity.Destination,
		Status:      string(entity.Status),
		CarrierUUID: entity.CarrierUUID,
	}
}
