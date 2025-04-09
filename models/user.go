package models

type User struct {
	ID       uint   `json:"id" form:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" form:"username" validate:"required,min=3,max=32" gorm:"size:32;not null;unique"`
	Password string `json:"password" form:"password" validate:"required,min=6" gorm:"not null"`
}
