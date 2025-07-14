package main

import (
	"os"

	"github.com/blavejr/bowattServer/controllers"
	"github.com/blavejr/bowattServer/middleware"
	"github.com/blavejr/bowattServer/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"log/slog"
)

func main() {
	// Set up logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// err := models.ConnectMongo("mongodb://mongo:27017")
	err := models.ConnectMongo("mongodb://localhost:27017")
	if err != nil {
		slog.Error("Failed to connect to MongoDB", "error", err)
		os.Exit(1)
	}

	// TODO: load these dynamically from config
	// For now, hardcoded for simplicity
	fileCollection := models.GetCollection("bowattDB", "files")
	usersCollection := models.GetCollection("bowattDB", "users")
	queriesCollection := models.GetCollection("bowattDB", "queries")
	filesControllers := controllers.NewFilesHandler(fileCollection)
	queriesControllers := controllers.NewQueriesHandler(queriesCollection, fileCollection)
	usersControllers := controllers.NewUsersHandler(usersCollection)

	router := gin.Default()

	// Allow requests from localhost:3000
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// test route
	router.GET("/ping", func(c *gin.Context) {
		slog.Info("Health ping received")
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Public routes
	router.POST("/signup", usersControllers.CreateUser)
	router.POST("/login", usersControllers.LoginUser)

	// Protected routes group
	router.Use(middleware.AuthMiddleware(usersCollection))
	{
		// Protected User routes
		router.GET("/profile", func(c *gin.Context) {
			user := c.MustGet("user").(models.User)
			c.JSON(200, gin.H{"username": user.Username, "created_at": user.CreatedAt})
		})

		// Protected File routes
		router.POST("/upload", filesControllers.HandleFileUpload)
		router.GET("/files", filesControllers.ListFilesForUser)

		// Query routes
		router.POST("/query", queriesControllers.HandleQuery)
	}

	router.Run(":8080")
}
