package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint           `gorm:"primaryKey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`

	Email           string         `gorm:"uniqueIndex;not null"`
	PasswordHash    string         `gorm:"not null"`
	Role            string         `gorm:"default:user"`
	IsEmailVerified bool           `gorm:"default:false"`
}
