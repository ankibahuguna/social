package db

import (
	"fmt"
	"github.com/ankibahuguna/social/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
)

var config = struct {
	DBName   string `default:"blog" env:"DBName"`
	User     string `default:"root" env:"DBUser"`
	Host     string `default:"127.0.0.1" env:"DBHost"`
		Password string `default:"123456" env:"DBPassword"`
	Port     string `default:"3306" env:"DBPort"`
}{}

func New() (db *gorm.DB) {
	configor.Load(&config)

	args := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
		)

	db, err := gorm.Open("mysql", args)

	if err != nil {
		panic(err)
	}

	db.LogMode(true)

	autoMigrate(db)

	return
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})
}
