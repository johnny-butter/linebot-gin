package reply

import (
	"fmt"
	"linebot-gin/models"
	"strings"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Skill struct{}

func (self *Skill) New(_ string, _ MsgSource) ReplyMessage { return self }

func (self *Skill) Messages() []linebot.SendingMessage {
	messages := []linebot.SendingMessage{}

	var keywords []models.Keyword
	models.DB.Find(&keywords)

	if len(keywords) == 0 {
		return messages
	}

	for _, v := range keywords {
		content := []string{
			fmt.Sprintf("🔑關鍵字: %v", v.Name),
			fmt.Sprintf("🔎說明: %v", v.Description),
			fmt.Sprintf("📝用法: %v", v.Usage),
		}

		messages = append(messages, linebot.NewTextMessage(strings.Join(content, "\n")))
	}

	return messages
}
