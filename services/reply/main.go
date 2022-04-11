package reply

import (
	"fmt"
	"linebot-gin/models"
	"strings"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"gorm.io/gorm"
)

type ReplyMessage interface {
	Messages() []linebot.SendingMessage
}

type KeywordReplyMessage interface {
	ReplyMessage
	New(string, MsgSource) ReplyMessage
}

type MsgSource struct {
	Type linebot.EventSourceType
	Id   string
}

var strStructMap = map[string]KeywordReplyMessage{
	"WeatherForecast": &WeatherForecast{},
	"FindFood":        &FindFood{},
	"Skill":           &Skill{},
}

func New(msg linebot.Message, msgSource MsgSource) ReplyMessage {
	switch message := msg.(type) {
	case *linebot.TextMessage:
		splitedStr := strings.SplitN(message.Text, " ", 2)

		keyword := models.Keyword{}
		qResult := models.DB.First(&keyword, "name = ?", splitedStr[0])
		if qResult.Error == gorm.ErrRecordNotFound {
			return nil
		}

		cls, ok := strStructMap[keyword.Method]
		if !ok {
			return nil
		}

		var param string
		if len(splitedStr) >= 2 {
			param = splitedStr[1]
		}

		return cls.New(param, msgSource)
	case *linebot.StickerMessage:
		return &RandomSticker{
			PackageId:    "8525",
			MinStickerId: 16581290,
			MaxStickerId: 16581313,
		}
	case *linebot.LocationMessage:
		return &SaveLocation{
			MsgSource: msgSource,
			Location:  fmt.Sprint(message.Latitude, ",", message.Longitude),
		}
	default:
		return nil
	}
}
