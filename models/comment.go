package models

import "gorm.io/gorm"

// Comment ...
type Comment struct {
	gorm.Model

	PostID  uint
	Post    Post
	Content string
}
