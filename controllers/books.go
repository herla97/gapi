package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herla97/gapi/models"
	"github.com/jinzhu/gorm"
)

// GET /books
// Get all books
func FindBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var books []models.Book
	db.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}
