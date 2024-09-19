package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/spaaws/book-api/controllers"
)

func SetupRoutes(r *gin.Engine) {
    // Define as rotas
    r.GET("/books", controllers.GetBooks)
    r.GET("/books/:id", controllers.GetBooks)
    r.POST("/books", controllers.CreateBook)
    r.PUT("/books/:id", controllers.UpdateBook)
    r.DELETE("/books/:id", controllers.DeleteBook)
}