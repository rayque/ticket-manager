package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"shipping-management/internal/domain/app_errors"
	"shipping-management/internal/domain/entities"
	"shipping-management/internal/domain/interfaces"
	"shipping-management/internal/infrastructure/repositories/dtos"
	"strings"
)

const carriersCollection = "carriers"

type carrierRepository struct {
	coll *mongo.Collection
}

func NewCarrierRepository(
	db *mongo.Database,
) interfaces.CarrierRepository {
	return &carrierRepository{
		coll: db.Collection(carriersCollection),
	}
}

func (c carrierRepository) GetRegionByState(ctx context.Context, state string) string {
	return strings.ToUpper(entities.StateRegion[strings.ToUpper(state)])
}

func (c carrierRepository) GetCarriersByRegion(ctx context.Context, region string) ([]entities.Carrier, error) {
	filter := bson.D{{Key: "regions.name", Value: region}}
	cursor, err := c.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var carriers []dtos.Carrier
	if err = cursor.All(ctx, &carriers); err != nil {
		return nil, err
	}

	if len(carriers) == 0 {
		return nil, app_errors.ErrNoCarrierFound
	}

	var result []entities.Carrier
	for _, carrier := range carriers {
		result = append(result, carrier.ToEntity())
	}
	return dtos.CarriersDtoToEntities(carriers), nil
}
