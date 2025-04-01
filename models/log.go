package models

import "time"

type Log struct {
	ID      uint      `gorm:"primarykey"`
	Student Student   `gorm:"not null"`
	Action  Action    `gorm:"not null"`
	Date    time.Time `gorm:"not null"`
}
