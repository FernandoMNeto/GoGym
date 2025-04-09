package models

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name        string    `gorm:"not null"`
	Password    string    `gorm:"unique"`
	Address     Address   `gorm:"not null"`	
	Email       string    `gorm:"unique"`
	Phone       string    `gorm:"unique"`
	LastChekIn  time.Time `gorm:"unique;not null"`
	LastChekOut time.Time `gorm:"unique;not null"`
	Biometrics  []byte    `gorm:"type:bytea"`
}
