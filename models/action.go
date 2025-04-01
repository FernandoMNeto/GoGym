package models

type Action struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"not null;unique"`
}
