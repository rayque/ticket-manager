package entities

type Status string

const (
	StatusCreated              Status = "CREATED"
	StatusWaitingForCollection Status = "WAITING_FOR_COLLECTION"
	StatusCollected            Status = "COLLECTED"
	StatusSent                 Status = "SENT"
	StatusDelivered            Status = "DELIVERED"
	StatusLost                 Status = "LOST"
)
