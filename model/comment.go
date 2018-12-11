package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	Model
	Text   string `gorm:"type:TEXT" json:"text"`
	Post   Post   `gorm:"foreignkey:Post" json:"post"`
	Author uint   `gorm:"foreignkey:User" json:"author"`
}

func (c *Comment) Create(db *gorm.DB) (err error) {
	err = db.Create(c).Error
	return
}

func DeleteComment(db *gorm.DB, id uint64) (comment Comment, err error) {
	err = db.Delete(&comment, id).Error
	return
}

func GetCommentsByPostId(db *gorm.DB, postid uint64) (comment []Comment, err error) {
	err = db.Preload("Author").Where("post=?", postid).Find(&comment).Error
	return
}
