package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/psuarezdev/go-api-starter/src/auth"
	"github.com/psuarezdev/go-api-starter/src/config"
	"github.com/psuarezdev/go-api-starter/src/database"
	"github.com/psuarezdev/go-api-starter/src/user"
)

func main() {
	godotenv.Load()

	if config.GetEnvironment() == "production" {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
		f, _ := os.Create("gin.log")
		gin.DefaultWriter = io.MultiWriter(f)
	}

	database.InitDatabase()

	db := database.GetConnection()

	db.AutoMigrate(&user.User{})

	router := gin.Default()
	// router.SetTrustedProxies([]string{"<ip_address>"})
	router.Use(cors.Default())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	setupRoutes(router)

	router.Run("127.0.0.1:4000")
}

func setupRoutes(router *gin.Engine) {
	// Rutas públicas de autenticación
	auth.SetupRoutes(
		router.Group(fmt.Sprintf("%s/auth", config.API_PREFIX)),
	)

	/*authorized := router.Group(config.API_PREFIX, middleware.AuthMiddleware())
	{
		// Protected routes here
	}*/
}
