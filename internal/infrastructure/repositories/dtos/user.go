package dtos

import "time"

type UserDTO struct {
	ID        int64  `gorm:"primaryKey;autoIncrement"`
	UUID      string `gorm:"uniqueIndex;not null"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	Phone     string
	Address   string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (UserDTO) TableName() string {
	return "users"
}
