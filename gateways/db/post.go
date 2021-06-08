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
	panic("not yet implemented")
}

// GetSingle ...
func (db Post) GetSingle(id int) (models.Post, error) {
	panic("not yet implemented")
}

// Create ...
func (db Post) Create(todo models.Post) (id int, err error) {
	panic("not yet implemented")
}

// Update ...
func (db Post) Update(id int, todo models.Post) error {
	panic("not yet implemented")
}

// Delete ...
func (db Post) Delete(id int) error {
	panic("not yet implemented")
}
