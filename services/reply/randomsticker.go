package reply

import (
	"math/rand"
	"strconv"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// https://developers.line.biz/en/docs/messaging-api/sticker-list/#sticker-definitions
type RandomSticker struct {
	PackageId    string
	MinStickerId int
	MaxStickerId int
}

func (self *RandomSticker) Message() linebot.SendingMessage {
	stickerId := rand.Intn(self.MaxStickerId-self.MinStickerId) + self.MinStickerId

	return linebot.NewStickerMessage(self.PackageId, strconv.Itoa(stickerId))
}
