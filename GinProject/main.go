package main

import (
	"github.com/BAmit1234/GinProject/controllers"
	"github.com/BAmit1234/GinProject/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()
	r.LoadHTMLGlob("templates/*.html")

	// Routes
	r.GET("/books", controllers.SearchBooks)
	r.GET("/books/update/:value", controllers.SearchBook)
	r.POST("/books/update/:value", controllers.Update)
	r.POST("/books", controllers.AddBook)
	r.GET("/books/del/:id", controllers.Delete)

	// Run the server
	r.Run()
}
