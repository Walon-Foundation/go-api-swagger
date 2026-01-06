package middlewares

import (
	"net/http"
	"strings"

	"github.com/Walon-Foundation/go-gin-doc/cmd/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


// AuthMiddleware verifies the Clerk session JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Empty Bearer token"})
			return
		}

		// Verify the session JWT
		claims, err := utils.VerifyToken(token)
		if err != nil || err == jwt.ErrTokenExpired  {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":"invalid token"})
			return
		}

		// Add claims or user ID to context for handlers
		c.Set("userId", claims.UserdId)

		c.Next()
	}
}