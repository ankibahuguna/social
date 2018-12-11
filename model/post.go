package model

import "github.com/jinzhu/gorm"
import "time"

type Post struct {
	Model
	Title   string `json:"title"`
	Content string `gorm:"type:TEXT" json:"content"`
	User    User   `gorm:"foreignkey:Author" json:"user"`
	Author  uint   `json:author`
}

type PostResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    uint      `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (e *Post) Create(db *gorm.DB) (err error) {
	err = db.Create(e).Error
	return
}

func FindPostById(db *gorm.DB, id uint64) (post Post, err error) {
	err = db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Table("users").Select("id,name,email")
	}).First(&post, id).Error
	return
}

func FindPosts(db *gorm.DB) (posts []PostResponse, err error) {
	err = db.Table("posts").Select("id, title, content, author, created_at, updated_at").Scan(&posts).Error
	return
}

func DeletePost(db *gorm.DB, id uint64) (post Post, err error) {
	err = db.Delete(&post, id).Error
	return
}

func FindPostsByAuthor(db *gorm.DB, authorid uint64) (post []PostResponse, err error) {
	err = db.Table("posts").Select("id, title, content, author, created_at, updated_at").Where("Author=?", authorid).Scan(&post).Error
	return
}
