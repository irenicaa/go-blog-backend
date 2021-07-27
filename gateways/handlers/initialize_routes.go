package handlers

import "github.com/labstack/echo/v4"

// InitializeRoutes ...
func InitializeRoutes(server *echo.Echo, postStorage PostStorage) {
	postHandler := Post{Storage: postStorage}
	server.GET("/api/v1/posts", postHandler.GetAll)
}
