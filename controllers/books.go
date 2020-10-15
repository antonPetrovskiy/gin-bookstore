package controllers

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/RazorEdgexD/gin-bookstore/models"
)

type CreateBookInput struct {
	Score  uint `json:"score" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Score  uint `json:"score"`
	Author string `json:"author"`
}

// GET /leaderbordAll
// Find all books
func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// GET /leaderbord
// Find all books
func FindBooksTen(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	if (nil != books) {
		sort.Slice(books, func(i, j int) bool { return books[i].Score > books[j].Score })
	}

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// GET /leaderbord/:id
// Find a book
func FindBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// GET /playerPlace/:id
// Find a book
func FindBookByPlayer(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("author = ?", c.Param("author")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// POST /books
// Create new book
func CreateBook(c *gin.Context) {
	// Validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{Score: input.Score, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// PATCH /books/:id
// Update a book
func UpdateBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// PATCH /leaderbordUUID/:author
// Update a book
func UpdateBookByPlayer(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("author = ?", c.Param("author")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /books/:id
// Delete a book
func DeleteBook(c *gin.Context) {
	// Get model if exist
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
