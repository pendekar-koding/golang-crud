package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api-crud/database"
	"rest-api-crud/models"
)

func CreateBook(c *gin.Context) {
	var book *models.Book
	err := c.ShouldBind(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	res := database.DB.Create(book)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error Creating a Book",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
	return
}

func ReadBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	res := database.DB.Find(&book, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Book Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
	return
}

func ReadBooks(c *gin.Context) {
	var books []models.Book
	res := database.DB.Find(&books)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Book Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
	return
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	err := c.ShouldBind(&book)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	var updateBook models.Book
	res := database.DB.Model(&updateBook).Where("id = ?", id).Updates(book)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update Success",
	})
	return
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	res := database.DB.Find(&book, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Book not found",
		})
		return
	}

	database.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted Success",
	})
	return
}
