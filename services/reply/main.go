package reply

import (
	"fmt"
	"strings"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type ReplyMessage interface {
	Messages() []linebot.SendingMessage
}

type MsgSource struct {
	Type linebot.EventSourceType
	Id   string
}

func New(msg linebot.Message) ReplyMessage {
	switch message := msg.(type) {
	case *linebot.TextMessage:
		weatherKeyword := "天氣"

		if strings.HasSuffix(message.Text, weatherKeyword) {
			locationStr := strings.Replace(message.Text, weatherKeyword, "", -1)
			locationRune := []rune(locationStr)

			return &WeatherForecast{
				CountyName:   string(locationRune[:3]),
				LocationName: string(locationRune[3:]),
			}
		}

		return nil
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
