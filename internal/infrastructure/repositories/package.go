package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"shipping-management/internal/domain/app_errors"
	"shipping-management/internal/domain/entities"
	"shipping-management/internal/domain/interfaces"
	"shipping-management/internal/infrastructure/repositories/dtos"
	"time"
)

type PostgresPackageRepository struct {
	db *gorm.DB
}

func NewPackageRepository(postgresDB *gorm.DB) interfaces.PackageRepository {
	return &PostgresPackageRepository{
		db: postgresDB,
	}
}

func (r *PostgresPackageRepository) CreatePackage(ctx context.Context, pkg entities.Package) (entities.Package, error) {
	p := dtos.FromEntityToPackageDto(pkg)
	p.Status = string(entities.StatusCreated)

	result := r.db.WithContext(ctx).Create(&p)
	if result.Error != nil {
		return entities.Package{}, result.Error
	}
	pkg.ID = p.ID
	pkg.Status = entities.Status(p.Status)
	return pkg, nil
}

func (r *PostgresPackageRepository) GetPackageByUuid(ctx context.Context, uuid string) (entities.Package, error) {
	var pkg dtos.PackageDto

	result := r.db.WithContext(ctx).First(&pkg, "uuid = ?", uuid)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entities.Package{}, app_errors.ErrPackageNotFound
		}
		return entities.Package{}, result.Error
	}

	return pkg.ToEntity(), nil
}

func (r *PostgresPackageRepository) UpdatePackage(ctx context.Context, p entities.Package) (entities.Package, error) {
	var pkg dtos.PackageDto

	result := r.db.WithContext(ctx).First(&pkg, "uuid = ?", p.UUID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entities.Package{}, app_errors.ErrPackageNotFound
		}
		return entities.Package{}, result.Error
	}

	pkg.Status = string(p.Status)
	pkg.CarrierUUID = p.CarrierUUID
	pkg.UpdatedAt = time.Now()

	updateResult := r.db.WithContext(ctx).Save(&pkg)
	if updateResult.Error != nil {
		return entities.Package{}, updateResult.Error
	}

	return pkg.ToEntity(), nil
}
