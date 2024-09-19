package main

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/spaaws/book-api/config"
    "github.com/spaaws/book-api/models"
    "github.com/spaaws/book-api/routes"
    "os"
)

func main() {
    // Conecta ao banco de dados
    config.Connect()
    config.DB.AutoMigrate(&models.Book{})

    // Pega a porta das variáveis de ambiente ou usa a porta 8080 como fallback
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Cria uma nova instância do roteador GIN
    r := gin.Default()

    // Configura o middleware CORS
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"}, // Permite todas as origens. Ajuste conforme necessário.
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: true,
    }))

    // Configura as rotas
    routes.SetupRoutes(r)

    // Inicia o servidor na porta definida
    r.Run(":" + port)
}