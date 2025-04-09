package models

type Address struct {
	ID uint			  `gorm:"primarykey"`	
	Street string 
	State string
	Number string  	
}