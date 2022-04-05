package reply

import (
	"encoding/json"
	"fmt"
	"linebot-gin/services/cache"
	"linebot-gin/services/requests"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type FindFood struct {
	MsgSource MsgSource
	Keyword   string
}

func (self *FindFood) Messages() []linebot.SendingMessage {
	messages := []linebot.SendingMessage{}

	cacheKey := fmt.Sprint(self.MsgSource.Type, ":", self.MsgSource.Id, ":location")
	location := self.getLocation(cacheKey)
	if len(location) == 0 {
		messages = append(messages, linebot.NewTextMessage("你在哪兒呢? $").AddEmoji(linebot.NewEmoji(7, "5ac22a8c031a6752fb806d66", "030")))
		return messages
	}

	placesResp, _ := self.getPlaces(location)

	places := placesResp.Results

	if len(places) == 0 {
		messages = append(messages, linebot.NewTextMessage(fmt.Sprint("無營業中 \"", self.Keyword, "\" 相關地點😱")))
		return messages
	}

	sort.SliceStable(places, func(i, j int) bool {
		return places[i].Rating > places[j].Rating
	})

	maxRtn := 5

	if len(places) >= maxRtn {
		places = places[0:maxRtn]
	}

	var content []string

	for _, place := range places {
		content = append(content, fmt.Sprint("店名: ", place.Name))
		content = append(content, fmt.Sprint("地址: ", place.Address))
		content = append(content, fmt.Sprint("評價: ", place.Rating, " / ", place.UserRatingsTotal))

		messages = append(messages, linebot.NewTextMessage(strings.Join(content, "\n")))
		content = content[:0]
	}

	return messages
}

func (self *FindFood) getLocation(cacheKey string) string {
	c := cache.NewCache(cache.Redis)

	return c.Get(cacheKey)
}

func (self *FindFood) getPlaces(location string) (*PlacesResp, error) {
	resp, err := requests.Get(
		"https://maps.googleapis.com/maps/api/place/nearbysearch/json",
		map[string]string{
			"key":      os.Getenv("GOOGLE_API_KEY"),
			"location": location,
			"keyword":  self.Keyword,
			"language": "zh-TW",
			"radius":   "1000",
			"type":     "restaurant",
			"opennow":  "true",
		},
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	ret := PlacesResp{}

	if err := json.Unmarshal(resp, &ret); err != nil {
		log.Println(err)
		return nil, err
	}

	return &ret, nil
}

type PlacesResp struct {
	Results []struct {
		Name             string  `json:"name"`
		Rating           float32 `json:"rating"`
		UserRatingsTotal int32   `json:"user_ratings_total"`
		Address          string  `json:"vicinity"`
		PriceLevel       int8    `json:"price_level"`
	} `json:"results"`
}
