package reply

import (
	"strings"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type ReplyMessage interface {
	Message() linebot.SendingMessage
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
	default:
		return nil
	}
}
