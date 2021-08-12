package controllers

import (
	"net/http"

	"github.com/BAmit1234/GinProject/models"
	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func SearchBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	c.HTML(http.StatusOK, "index.html", gin.H{"data": books})
	// c.JSON(http.StatusOK, gin.H{"data": books})
}

func SearchBook(c *gin.Context) {

	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("value")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.HTML(http.StatusOK, "update.html", gin.H{"id": c.Param("value"), "title": book.Title, "Author": book.Author})

	// c.JSON(http.StatusOK, gin.H{"data": book})
}
func Update(c *gin.Context) {

	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("value")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	title := c.PostForm("Bookr")
	author := c.PostForm("Authorr")
	input := models.Book{Title: title, Author: author}

	models.DB.Model(&book).Updates(input)

	c.HTML(http.StatusOK, "del.html", nil)
}

func AddBook(c *gin.Context) {
	// Validate input
	// var input CreateBookInput

	// for key, value := range c.Request.PostForm {
	// 	fmt.Println(key, value)
	// }

	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// Create book
	title := c.PostForm("Bookr")
	author := c.PostForm("Authorr")
	book := models.Book{Title: title, Author: author}
	models.DB.Create(&book)
	c.HTML(http.StatusOK, "del.html", nil)

}

func Delete(c *gin.Context) {

	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)
	c.HTML(http.StatusOK, "del.html", nil)

	// c.JSON(http.StatusOK, gin.H{"data": true})
}
