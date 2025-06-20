package entities

type Quotation struct {
	Carrier           string  `json:"carrier"`
	CarrierUUID       string  `json:"carrier_uuid"`
	Price             float64 `json:"estimated_price"`
	DeliveryTimeByDay int     `json:"estimated_delivery_time_by_day"`
}
