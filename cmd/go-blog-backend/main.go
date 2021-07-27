package main

import (
	"github.com/irenicaa/go-blog-backend/gateways/db"
	"github.com/irenicaa/go-blog-backend/gateways/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	server := echo.New()

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	const dsn = "postgresql://postgres:postgres@localhost:5432" +
		"/postgres?sslmode=disable"
	dbPool, err := db.OpenDB(dsn)
	if err != nil {
		server.Logger.Fatal(err)
	}

	handlers.InitializeRoutes(server, db.NewPost(dbPool))

	if err := server.Start(":8080"); err != nil {
		server.Logger.Fatal(err)
	}
}
