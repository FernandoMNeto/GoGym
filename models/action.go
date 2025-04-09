package models

type Action struct {
	ID   uint   `json:"id"   gorm:"primarykey"`
	Name string `json:"name" gorm:"not null;unique"`
}
