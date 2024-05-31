package routes

import (
	"log"
	"os"
	"spo_task_3/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Route() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve the FRONTEND_URL from environment variables
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173" // Default to localhost:5173 if not set
	}

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		log.Printf("Request: %s %s\n", c.Request.Method, c.Request.URL)
		c.Next()
		log.Printf("Response: %d\n", c.Writer.Status())
	})

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontendURL}, // List of allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.POST("/register", controllers.RegisterHelper)
	router.POST("/login", controllers.LoginHelper)
	log.Fatal(router.Run(":8080"))
}
