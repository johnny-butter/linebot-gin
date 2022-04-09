package reply

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"linebot-gin/models"
	"linebot-gin/services/requests"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"gorm.io/gorm"
)

type WeatherForecast struct {
	CountyName   string
	LocationName string
}

func (self *WeatherForecast) Messages() []linebot.SendingMessage {
	messages := []linebot.SendingMessage{}

	resp, err := self.getWeatherForecast()

	if err == gorm.ErrRecordNotFound {
		messages = append(messages, linebot.NewTextMessage(fmt.Sprint("無 \"", self.CountyName, self.LocationName, "\" 資料")))
		return messages
	}

	locationData := resp.Records.Locations[0].Location[0]

	content := []string{
		fmt.Sprint(locationData.LocationName, "@", resp.Records.Locations[0].LocationsName),
		fmt.Sprint("⌚", locationData.WeatherElement[1].Time[0].StartTime[5:16], "~", locationData.WeatherElement[1].Time[0].EndTime[5:16]),
		fmt.Sprint("🌡️體感: ", locationData.WeatherElement[0].Time[0].ElementValue[0].Value, locationData.WeatherElement[0].Time[0].ElementValue[0].Measures),
	}

	messages = append(messages, linebot.NewTextMessage(strings.Join(content, "\n")))
	messages = append(messages, linebot.NewTextMessage(fmt.Sprint("ℹ️", locationData.WeatherElement[1].Time[0].ElementValue[0].Value)))

	return messages
}

func (self *WeatherForecast) getWeatherForecast() (*WeatherForecastResp, error) {
	var county models.County
	qResult := models.DB.Debug().Model(&models.County{}).
		Joins("JOIN district ON county.id = district.county_id AND district.name = ?", self.LocationName).
		First(&county, "county.name = ?", self.CountyName)

	if qResult.Error != nil {
		log.Println(qResult.Error)
		return new(WeatherForecastResp), qResult.Error
	}

	tw, _ := time.LoadLocation("Asia/Taipei")
	current := time.Now().In(tw)
	currentPlus3 := current.Add(time.Hour * 3)
	timeLayout := "2006-01-02T15:04:05"

	resp, err := requests.Get(
		"https://opendata.cwb.gov.tw/api/v1/rest/datastore/"+county.CwbId,
		map[string]string{
			"Authorization": os.Getenv("CWB_AUTH_CODE"),
			"elementName":   "AT,WeatherDescription",
			"locationName":  self.LocationName,
			"timeFrom":      current.Format(timeLayout),
			"timeTo":        currentPlus3.Format(timeLayout),
		},
	)
	if err != nil {
		log.Println(err)
		return new(WeatherForecastResp), err
	}

	var ret WeatherForecastResp
	if err := json.Unmarshal(resp, &ret); err != nil {
		log.Println(err)
		return new(WeatherForecastResp), err
	}

	return &ret, nil
}

type WeatherForecastResp struct {
	Records struct {
		Locations []struct {
			LocationsName string `json:"locationsName"`
			Location      []struct {
				LocationName   string `json:"locationName"`
				WeatherElement []struct {
					Description string `json:"description"`
					Time        []struct {
						StartTime    string `json:"startTime"`
						EndTime      string `json:"endTime"`
						ElementValue []struct {
							Value    string `json:"value"`
							Measures string `json:"measures"`
						} `json:"elementValue"`
					} `json:"time"`
				} `json:"weatherElement"`
			} `json:"location"`
		} `json:"locations"`
	} `json:"records"`
}
