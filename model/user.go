package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	Model
	Name     string `json:"name"`
	Email    string `gorm:"not null;unique" json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

func FindUsers(db *gorm.DB) (users []User, err error) {
	err = db.Find(&users).Error
	return
}

func (e *User) Create(db *gorm.DB) (err error) {
	err = db.Create(e).Error
	return
}

func FindUserById(db *gorm.DB, id uint64) (user User, err error) {
	err = db.First(&user, id).Error
	return
}

func FindUserByEmail(db *gorm.DB, email string) (user User, err error) {
	err = db.Where("Email=?", email).First(&user).Error
	return
}

func (e *User) EmailExists(db *gorm.DB) (err error) {
	var count int64
	err = db.Where("Email=?", e.Email).First(&e).Count(&count).Error
	return
}
