package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/hernanrocha/minesweeper/tapcolors/controller"
)

func SetupRouter(client *redis.Client) *gin.Engine {
	// Controllers
	c := NewGameController(client)
	tc := controller.NewGameController(client)
	ws := controller.NewWebSocketController()

	// Default Engine
	r := gin.Default()

	// CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))

	// Ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Auth JWT
	// r.POST("/login", authMiddleware.LoginHandler)
	// r.POST("/register", auth.Register)

	// API v1
	v1 := r.Group("/api/v1")
	{
		// v1.Use(authMiddleware.MiddlewareFunc())

		// Refresh time can be longer than token timeout
		// auth.GET("/refresh_token", authMiddleware.RefreshHandler)

		v1.POST("/game", c.CreateGame)
		v1.GET("/game/:id", c.GetGame)
		v1.POST("/game/:id/reveal", c.RevealCell)
		v1.POST("/game/:id/flag", c.FlagCell)

		v1.POST("/tapcolors/game", tc.CreateGame)
		v1.GET("/tapcolors/ws", ws.WebSocket)
	}

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
