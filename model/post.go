package model

import "github.com/jinzhu/gorm"

type Post struct {
	Model
	Title   string `json:"title"`
	Content string `gorm:"type:TEXT" json:"content"`
	User    User   `gorm:"foreignkey:Author"`
	Author  uint   `json:author`
}

func (e *Post) Create(db *gorm.DB) (err error) {
	err = db.Create(e).Error
	return
}

func FindPostById(db *gorm.DB, id uint64) (post Post, err error) {
	err = db.First(&post, id).Error
	return
}

func FindPosts(db *gorm.DB) (posts []Post, err error) {
	err = db.Find(&posts).Error
	return
}

func DeletePost(db *gorm.DB, id uint64) (post Post, err error) {
	err = db.Delete(&post, id).Error
	return
}
