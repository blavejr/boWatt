// backend/middleware/auth.go
package middleware

import (
	"net/http"

	"github.com/blavejr/bowattServer/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthMiddleware(userColl *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}

		var user models.User
		err := userColl.FindOne(c, bson.M{"token": token}).Decode(&user)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Save user in context for handlers downstream
		c.Set("user", user)
		c.Next()
	}
}
