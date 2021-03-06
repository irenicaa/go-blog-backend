package db

import (
	"github.com/irenicaa/go-blog-backend/models"
	"gorm.io/gorm"
)

// Comment ...
type Comment struct {
	pool *gorm.DB
}

// NewComment ...
func NewComment(pool *gorm.DB) Comment {
	return Comment{pool: pool}
}

// GetAll ...
func (db Comment) GetAll(postID int) ([]models.Comment, error) {
	var comments []models.Comment
	err := db.pool.
		Where("post_id = ?", postID).
		Order("created_at").
		Find(&comments).
		Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}

// Create ...
func (db Comment) Create(postID uint, comment models.Comment) (
	id int,
	err error,
) {
	comment.PostID = postID

	if err := db.pool.Create(&comment).Error; err != nil {
		return 0, err
	}

	return int(comment.ID), nil
}
