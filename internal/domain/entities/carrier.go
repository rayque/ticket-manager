package entities

type Carrier struct {
	UUID   string   `json:"uuid"`
	Name   string   `json:"name"`
	Region []Region `json:"regions"`
}

type Region struct {
	Name         string  `json:"name"`
	DeliveryTime int     `json:"delivery_time_day"`
	PricePerKg   float64 `json:"price_per_kg"`
}

var StateRegion = map[string]string{
	"RS": "SUL", "SC": "SUL", "PR": "SUL",

	"SP": "SUDESTE", "RJ": "SUDESTE", "MG": "SUDESTE", "ES": "SUDESTE",

	"MT": "CENTRO-OESTE", "MS": "CENTRO-OESTE", "GO": "CENTRO-OESTE", "DF": "CENTRO-OESTE",

	"BA": "NORDESTE", "SE": "NORDESTE", "AL": "NORDESTE", "PE": "NORDESTE",
	"PB": "NORDESTE", "RN": "NORDESTE", "CE": "NORDESTE", "PI": "NORDESTE", "MA": "NORDESTE",

	"AM": "NORTE", "RR": "NORTE", "AP": "NORTE", "PA": "NORTE", "TO": "NORTE", "RO": "NORTE", "AC": "NORTE",
}
