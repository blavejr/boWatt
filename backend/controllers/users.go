package controllers

import (
	"net/http"
	"time"

	"github.com/blavejr/bowattServer/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UsersHandler struct {
	UserCollection *mongo.Collection
}

func NewUsersHandler(userColl *mongo.Collection) *UsersHandler {
	return &UsersHandler{UserCollection: userColl}
}

func (handler *UsersHandler) CreateUser(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if user already exists
	var existing models.User
	err := handler.UserCollection.Database().Collection("users").
		FindOne(c, bson.M{"username": input.Username}).
		Decode(&existing)

	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		Username:  input.Username,
		Password:  string(hashedPassword),
		Token:     uuid.NewString(),
		CreatedAt: time.Now().Unix(),
	}

	_, err = handler.UserCollection.Database().Collection("users").InsertOne(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created",
		"token":   user.Token,
	})
}

func (handler *UsersHandler) LoginUser(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	err := handler.UserCollection.Database().Collection("users").
		FindOne(c, bson.M{"username": input.Username}).
		Decode(&user)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate new token
	newToken := uuid.NewString()
	_, err = handler.UserCollection.Database().Collection("users").
		UpdateOne(c,
			bson.M{"username": input.Username},
			bson.M{"$set": bson.M{"token": newToken}},
		)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   newToken,
	})
}
