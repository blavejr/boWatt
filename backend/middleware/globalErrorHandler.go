package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		errs := c.Errors
		if len(errs) > 0 {
			slog.Error("Global error handler", "Errors:", errs)

			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": errs[0].Error(),
			})
		}
	}
}
