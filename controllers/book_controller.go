package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/spaaws/book-api/config"
    "github.com/spaaws/book-api/models"
)

func GetBooks(c *gin.Context) {
    var books []models.Book
    config.DB.Find(&books)
    c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
    id := c.Param("id")
    var book models.Book
    if err := config.DB.First(&book, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado!"})
        return
    }
    c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&book)
    c.JSON(http.StatusCreated, book)
}

func UpdateBook(c *gin.Context) {
    id := c.Param("id")
    var book models.Book
    if err := config.DB.First(&book, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado!"})
        return
    }
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Save(&book)
    c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
    id := c.Param("id")
    if err := config.DB.Delete(&models.Book{}, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado!"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Livro deletado com sucesso!"})
}