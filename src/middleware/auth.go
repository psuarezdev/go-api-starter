package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/psuarezdev/go-api-starter/src/lib"
	"github.com/psuarezdev/go-api-starter/src/user"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			ctx.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token is required"})
			ctx.Abort()
			return
		}

		userId := lib.ValidateToken(token)

		if userId <= 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		user, err := user.GetById(uint(userId))

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			ctx.Abort()
			return
		}

		user.Password = ""
		ctx.Set("user", user)
		ctx.Next()
	}
}
