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
	GetSingle(id int) (models.Post, error)
	Create(post models.Post) (id int, err error)
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

// GetSingle ...
func (handler Post) GetSingle(ctx echo.Context) error {
	var id int
	if err := echo.PathParamsBinder(ctx).Int("id", &id).BindError(); err != nil {
		return fmt.Errorf("unable to bind the path parameters: %w", err)
	}

	post, err := handler.Storage.GetSingle(id)
	if err != nil {
		return fmt.Errorf("unable to get the post: %w", err)
	}

	return ctx.JSON(http.StatusOK, post)
}

// Create ...
func (handler Post) Create(ctx echo.Context) error {
	var post models.Post
	if err := ctx.Bind(&post); err != nil {
		return fmt.Errorf("unable to bind the body: %w", err)
	}

	id, err := handler.Storage.Create(post)
	if err != nil {
		return fmt.Errorf("unable to create the post: %w", err)
	}

	post.ID = uint(id)
	return ctx.JSON(http.StatusOK, post)
}
