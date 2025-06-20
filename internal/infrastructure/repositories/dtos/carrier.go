package dtos

import "shipping-management/internal/domain/entities"

type Carrier struct {
	UUID   string   `json:"uuid" bson:"uuid"`
	Name   string   `json:"name" bson:"name"`
	Region []Region `json:"regions" bson:"regions"`
}

type Region struct {
	Name         string  `json:"name" bson:"name"`
	DeliveryTime int     `json:"delivery_time_day" bson:"delivery_time_day"`
	PricePerKg   float64 `json:"price_per_kg" bson:"price_per_kg"`
}

func CarriersDtoToEntities(carriers []Carrier) []entities.Carrier {
	result := make([]entities.Carrier, 0, len(carriers))
	for _, c := range carriers {
		result = append(result, c.ToEntity())
	}
	return result
}

func (c Carrier) ToEntity() entities.Carrier {
	return entities.Carrier{
		UUID:   c.UUID,
		Name:   c.Name,
		Region: c.ToEntityRegions(),
	}
}

func (c Carrier) ToEntityRegions() []entities.Region {
	regions := make([]entities.Region, 0, len(c.Region))
	for _, r := range c.Region {
		regions = append(regions, entities.Region{
			Name:         r.Name,
			DeliveryTime: r.DeliveryTime,
			PricePerKg:   r.PricePerKg,
		})
	}
	return regions

}
