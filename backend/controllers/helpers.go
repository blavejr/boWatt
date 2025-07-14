package controllers

import (
	"github.com/blavejr/bowattServer/models"
	"github.com/gin-gonic/gin"
)

// CurrentUser retrieves the current user from the context
func CurrentUser(c *gin.Context) models.User {
	user, exists := c.Get("user")
	if !exists {
		return models.User{}
	}
	return user.(models.User)
}
