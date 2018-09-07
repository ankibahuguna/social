package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Like struct {
	Model
	User   User `gorm:"foreignkey:author"`
	Author uint
}

func (l *Like) Create(db *gorm.DB) (err error) {
	err = db.Create(l).Error
	return
}
