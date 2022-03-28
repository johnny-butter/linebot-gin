package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"linebot-gin/services/reply"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	bot, err := linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
	)

	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/callback", func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		log.Println(string(body))
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

		events, err := bot.ParseRequest(c.Request)

		if err != nil {
			if err == linebot.ErrInvalidSignature {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "invalid signature error",
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "events parse error",
			})
			return
		}

		for _, event := range events {
			if event.Type != linebot.EventTypeMessage {
				log.Println(fmt.Sprint("\"", event.Type, "\" not support"))
				continue
			}

			replyInstance := reply.New(event.Message)

			if replyInstance == nil {
				continue
			}

			if _, err := bot.ReplyMessage(event.ReplyToken, replyInstance.Message()).Do(); err != nil {
				log.Println(err)
			}
		}

		c.Status(http.StatusNoContent)
	})

	if port, ok := os.LookupEnv("PORT"); ok {
		router.Run(":" + port)
	} else {
		router.Run(":8080")
	}
}
