package db

import "gorm.io/gorm"

// Comment ...
type Comment struct {
	pool *gorm.DB
}

// NewComment ...
func NewComment(pool *gorm.DB) Comment {
	return Comment{pool: pool}
}
