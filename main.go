package main

import (
	"library-management/database"
	"library-management/handlers"
	"library-management/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDatabase()

	// setup gin
	r := gin.Default()

	auth := r.Group("/admin")
	auth.Use(middleware.AuthMiddleware())

	// Books
	r.POST("/books", handlers.CreateBook)
	r.GET("/books", handlers.GetBooks)
	r.GET("/books/:name", handlers.GetBookByName)
	r.PUT("/books/:id", handlers.UpdateBook)
	r.DELETE("/books/:id", handlers.DeleteBook)

	// Users
	auth.GET("/users", middleware.AdminOnly(), handlers.GetUsers)

	r.Run(":8080")
}
