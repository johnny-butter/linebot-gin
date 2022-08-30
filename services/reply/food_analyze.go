package reply

import (
	"fmt"
	"linebot-gin/models"
	"log"
	"strings"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type FoodAnalyze struct {
	FoodName string
}

func (self *FoodAnalyze) New(param string, _ MsgSource) ReplyMessage {
	self.FoodName = param

	return self
}

func (self *FoodAnalyze) Messages() []linebot.SendingMessage {
	messages := []linebot.SendingMessage{}

	cantFindMsg := linebot.NewTextMessage(fmt.Sprint("找不到\"", self.FoodName, "\" $")).
		AddEmoji(linebot.NewEmoji(len(self.FoodName)+6, "5ac22a8c031a6752fb806d66", "027"))

	var content []string

	var food models.Food
	qResult := models.DB.Debug().Model(&models.Food{}).
		Where("name = ?", self.FoodName).
		First(&food)

	if qResult.Error != nil {
		log.Println(qResult.Error)
		messages = append(messages, cantFindMsg)
		return messages
	}

	content = append(content, fmt.Sprint("名稱: ", food.Name))
	content = append(content, fmt.Sprint("名稱(英文): ", food.NameEng))
	content = append(content, fmt.Sprint("類別: ", food.Category))
	content = append(content, fmt.Sprint("俗名: ", food.CommonNames))

	var foodIngredients []models.FoodIngredient
	qResult = models.DB.Debug().Model(&models.FoodIngredient{}).Find(&foodIngredients, "food_id = ?", food.ID)

	if qResult.Error != nil {
		log.Println(qResult.Error)
		messages = append(messages, cantFindMsg)
		return messages
	}

	content = append(content, fmt.Sprint("營養(每 100 克):"))

	for _, foodIngredient := range foodIngredients {
		content = append(content, fmt.Sprint(foodIngredient.Name, " => ", foodIngredient.Amount))
	}

	messages = append(messages, linebot.NewTextMessage(strings.Join(content, "\n")))

	return messages
}
