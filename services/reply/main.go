package reply

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type ReplyMessage interface {
	Message() linebot.SendingMessage
}

func New(msg linebot.Message) ReplyMessage {
	switch msg.(type) {
	case *linebot.StickerMessage:
		return &RandomSticker{
			PackageId:    "8525",
			MinStickerId: 16581290,
			MaxStickerId: 16581313,
		}
	default:
		return &RandomSticker{
			PackageId:    "6632",
			MinStickerId: 11825374,
			MaxStickerId: 11825397,
		}
	}
}
