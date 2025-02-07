package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/oOSomnus/MyDrive/internal/dto"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, dto.ResponseEntity{Message: "Please login to perform the operation"},
			)
			return
		}
		tokenParts := strings.Split(token, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, dto.ResponseEntity{Message: "Please login to perform the operation"},
			)
			return
		}
		claim, err := ParseToken(tokenParts[1])
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				dto.ResponseEntity{Message: "Token expired or invalid, please login to perform the operation"},
			)
			return
		}
		c.Set("userId", claim.UserID)
		c.Next()
	}
}
