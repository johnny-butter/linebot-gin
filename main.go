package main

import (
	"linebot-gin/controllers"
	"linebot-gin/models"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	models.ConnectDB()

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/ping", controllers.Ping)

		v1.POST("/callback", controllers.Callback)
	}

	if port, ok := os.LookupEnv("PORT"); ok {
		router.Run(":" + port)
	} else {
		router.Run(":8080")
	}
}
