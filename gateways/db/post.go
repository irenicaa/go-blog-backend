package db

import (
	"github.com/irenicaa/go-blog-backend/models"
	"gorm.io/gorm"
)

// Post ...
type Post struct {
	pool *gorm.DB
}

// NewPost ...
func NewPost(pool *gorm.DB) Post {
	return Post{pool: pool}
}

// GetAll ...
func (db Post) GetAll() ([]models.Post, error) {
	var posts []models.Post
	if err := db.pool.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

// GetSingle ...
func (db Post) GetSingle(id int) (models.Post, error) {
	var post models.Post
	if err := db.pool.Where("id = ?", id).First(&post).Error; err != nil {
		return models.Post{}, err
	}

	return post, nil
}

// Create ...
func (db Post) Create(post models.Post) (id int, err error) {
	if err := db.pool.Create(&post).Error; err != nil {
		return 0, err
	}

	return int(post.ID), nil
}

// Update ...
func (db Post) Update(id int, post models.Post) error {
	panic("not yet implemented")
}

// Delete ...
func (db Post) Delete(id int) error {
	panic("not yet implemented")
}
