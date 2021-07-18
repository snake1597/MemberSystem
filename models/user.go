package models

type User struct {
	ID       int    `gorm:"id" json:"id"`
	Account  string `gorm:"account" json:"account" binding:"required"`
	Password string `gorm:"password" json:"password" binding:"required"`
	Name     string `gorm:"name" json:"name" binding:"required"`
	Birthday string `gorm:"birthday" json:"birthday" binding:"required"`
	Token    string `gorm:"token" json:"token" binding:"required"`
}
 

