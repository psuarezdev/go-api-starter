package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/psuarezdev/go-api-starter/src/lib"
	"github.com/psuarezdev/go-api-starter/src/user"
)

func login(ctx *gin.Context) {
	var credentials struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
		return
	}

	userFound, err := user.GetByUsername(credentials.Username)

	if userFound == nil || err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong credentials"})
		return
	}

	passwordMatch := lib.VerifyPassword(credentials.Password, userFound.Password)

	if !passwordMatch {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong credentials"})
		return
	}

	token, err := lib.GenerateToken(userFound)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token, "user": userFound})
}

func register(ctx *gin.Context) {
	var credentials struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
		return
	}

	credentials.Password, _ = lib.HashPassword(credentials.Password)

	newUser := user.User{
		Username: credentials.Username,
		Password: credentials.Password,
	}

	if err := user.Create(&newUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	token, _ := lib.GenerateToken(&newUser)

	ctx.JSON(http.StatusCreated, gin.H{
		"token": token,
		"user":  newUser,
	})
}

func profile(ctx *gin.Context) {
	accessToken := ctx.GetHeader("Authorization")

	if accessToken == "" || !strings.HasPrefix(accessToken, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
		return
	}

	tokenString := strings.TrimPrefix(accessToken, "Bearer ")
	userId := lib.ValidateToken(tokenString)

	if userId == -1 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	userFound, err := user.GetById(uint(userId))
	if err != nil || userFound == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, userFound)
}

func SetupRoutes(router *gin.RouterGroup) {
	router.POST("/login", login)
	router.POST("/register", register)
	router.GET("/profile", profile)
}
