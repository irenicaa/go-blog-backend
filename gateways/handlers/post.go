package handlers

import (
	"fmt"
	"net/http"

	"github.com/irenicaa/go-blog-backend/models"
	"github.com/labstack/echo/v4"
)

// PostStorage ...
type PostStorage interface {
	GetAll(pagination models.Pagination) ([]models.Post, error)
}

// Post ...
type Post struct {
	Storage PostStorage
}

// GetAll ...
func (handler Post) GetAll(ctx echo.Context) error {
	posts, err := handler.Storage.GetAll(models.Pagination{})
	if err != nil {
		return fmt.Errorf("unable to get the posts: %w", err)
	}

	return ctx.JSON(http.StatusOK, posts)
}
