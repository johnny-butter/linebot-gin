package reply

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// https://developers.line.biz/en/docs/messaging-api/sticker-list/#sticker-definitions
type RandomSticker struct {
	PackageId    string
	MinStickerId int
	MaxStickerId int
}

func (self *RandomSticker) Messages() []linebot.SendingMessage {
	messages := []linebot.SendingMessage{}

	rand.Seed(time.Now().Unix())
	stickerId := rand.Intn(self.MaxStickerId-self.MinStickerId) + self.MinStickerId

	messages = append(messages, linebot.NewStickerMessage(self.PackageId, strconv.Itoa(stickerId)))

	return messages
}
