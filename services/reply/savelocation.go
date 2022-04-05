package reply

import (
	"fmt"
	"linebot-gin/services/cache"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type SaveLocation struct {
	MsgSource MsgSource
	Location  string
}

func (self *SaveLocation) Messages() []linebot.SendingMessage {
	messages := []linebot.SendingMessage{}

	c := cache.NewCache(cache.Redis)
	cacheKey := fmt.Sprint(self.MsgSource.Type, ":", self.MsgSource.Id, ":location")
	if err := c.Set(cacheKey, self.Location, cache.Settings{Ttl: 300 * time.Second}); err != nil {
		return nil
	}

	// https://developers.line.biz/en/docs/messaging-api/emoji-list/#specify-emojis-in-message-object
	messages = append(messages, linebot.NewTextMessage("地點已更新 $").AddEmoji(linebot.NewEmoji(6, "5ac22a8c031a6752fb806d66", "006")))

	return messages
}
