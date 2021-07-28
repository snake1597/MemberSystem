package models

import (
	"gorm.io/gorm"
)

//binding:"required"
type User struct {
	ID       int    `gorm:"primaryKey; id" json:"id"`
	Account  string `gorm:"account" json:"account" `
	Password string `gorm:"password" json:"password" `
	Name     string `gorm:"name" json:"name" `
	Birthday string `gorm:"birthday" json:"birthday" `
}

func (user *User) Insert(db *gorm.DB) (err error) {
	err = db.Create(&user).Error
	return
}

func (user *User) Update(db *gorm.DB, account interface{}) (err error) {
	err = db.Table("users").Where("account = ?", account).Updates(user).Error
	return
}

func (user *User) FindOne(db *gorm.DB, account interface{}) (err error) {
	err = db.Table("users").Where("account = ?", account).Find(&user).Error
	return
}
