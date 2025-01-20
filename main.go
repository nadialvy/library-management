package main

import (
	"library-management/database"
	"library-management/handlers"

	"github.com/gin-gonic/gin"
)


func main(){
	database.InitDatabase()

	// setup gin
	r := gin.Default()

	// Routes
	r.POST("/books", handlers.CreateBook)
	r.GET("/books", handlers.GetBooks)
	r.GET("/books/:name", handlers.GetBookByName)
	r.PUT("/books/:id", handlers.UpdateBook)
	r.DELETE("/books/:id", handlers.DeleteBook)

	r.Run(":8080")
}
