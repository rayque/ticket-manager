package entities

type Package struct {
	ID          int64   `json:"id"`
	UUID        string  `json:"uuid"`
	Product     string  `json:"product"`
	Weight      float64 `json:"weight"`
	Destination string  `json:"destination"`
	Status      Status  `json:"status"`
	CarrierUUID string  `json:"carrier_uuid"`
}
