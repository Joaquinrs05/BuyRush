package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email" gorm:"unique"`
	Photo	 string `json:"photo"`
	Username string `json:"username" gorm:"unique"`			
	Password string `json:"password"`
}
