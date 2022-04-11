package controllers

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
)

func Callback(c *gin.Context) {
	bot, err := linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

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

		msgSource := &reply.MsgSource{Type: event.Source.Type}

		switch msgSource.Type {
		case linebot.EventSourceTypeUser:
			msgSource.Id = event.Source.UserID
		case linebot.EventSourceTypeGroup:
			msgSource.Id = event.Source.GroupID
		default:
			log.Println(fmt.Sprint("\"", event.Source.Type, "\" not match"))
			continue
		}

		replyInstance := reply.New(event.Message, *msgSource)

		if replyInstance == nil {
			continue
		}

		if _, err := bot.ReplyMessage(event.ReplyToken, replyInstance.Messages()...).Do(); err != nil {
			log.Println(err)
		}
	}

	c.Status(http.StatusNoContent)
}
