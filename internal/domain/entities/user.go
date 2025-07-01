package entities

import "time"

type User struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	UUID      string    `json:"uuid" gorm:"uniqueIndex;not null"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
