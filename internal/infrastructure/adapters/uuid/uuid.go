package uuid

import (
	"github.com/google/uuid"
	"shipping-management/internal/domain/interfaces"
)

type uuidAdapter struct{}

func NewUUIDAdapter() interfaces.UUID {
	return &uuidAdapter{}
}

func (u *uuidAdapter) NewUUID() string {
	return uuid.New().String()
}

func (u *uuidAdapter) Generate() string {
	return uuid.New().String()
}
