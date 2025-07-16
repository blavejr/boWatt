package main

import (
	"os"

	"log/slog"

	"github.com/blavejr/bowattServer/controllers"
	"github.com/blavejr/bowattServer/middleware"
	"github.com/blavejr/bowattServer/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", "error", err)
	}

	// Set up logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Uncomment if you are not running in a docker env
	mongoURL := os.Getenv("DATABASE_URL_LOCAL")
	// mongoURL := os.Getenv("DATABASE_URL_DOCKER")
	if mongoURL == "" {
		slog.Error("DATABASE_URL not set")
		os.Exit(1)
	}

	err = models.ConnectMongo(mongoURL)
	if err != nil {
		slog.Error("Failed to connect to MongoDB", "error", err)
		os.Exit(1)
	}

	// Get port from env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fileCollection := models.GetCollection("bowattDB", "files")
	usersCollection := models.GetCollection("bowattDB", "users")
	queriesCollection := models.GetCollection("bowattDB", "queries")

	filesControllers := controllers.NewFilesHandler(fileCollection)
	queriesControllers := controllers.NewQueriesHandler(queriesCollection, fileCollection)
	usersControllers := controllers.NewUsersHandler(usersCollection)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router.GET("/ping", func(c *gin.Context) {
		slog.Info("Health ping received")
		c.JSON(200, gin.H{"message": "pong"})
	})

	router.POST("/signup", usersControllers.CreateUser)
	router.POST("/login", usersControllers.LoginUser)

	router.Use(middleware.AuthMiddleware(usersCollection))
	{
		router.GET("/profile", func(c *gin.Context) {
			user := c.MustGet("user").(models.User)
			c.JSON(200, gin.H{"username": user.Username, "created_at": user.CreatedAt})
		})
		router.POST("/upload", filesControllers.HandleFileUpload)
		router.GET("/files", filesControllers.ListFilesForUser)
		router.POST("/query", queriesControllers.HandleQuery)
		router.GET("/query/history", queriesControllers.QueryHistoryForUser)
	}

	router.Run(":" + port)
}
