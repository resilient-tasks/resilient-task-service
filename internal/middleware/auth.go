package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token faltante o inválido"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "JWT_SECRET no definido"})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Token inválido"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userId := claims["sub"]
		if userIdStr, ok := userId.(string); ok {
			c.Set("userId", userIdStr)
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Token sin userId"})
			return
		}

		role := claims["role"]
		if roleStr, ok := role.(string); ok {
			c.Set("role", roleStr)
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Token sin role"})
			return
		}

		c.Next()
	}
}
