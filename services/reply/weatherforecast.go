package reply

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"linebot-gin/services/requests"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type WeatherForecast struct {
	CountyName   string
	LocationName string
}

func (self *WeatherForecast) Messages() []linebot.SendingMessage {
	messages := []linebot.SendingMessage{}

	resp, _ := self.getWeatherForecast()

	locationData := resp.Records.Locations[0].Location[0]

	content := []string{
		fmt.Sprint("⚓", locationData.LocationName, "@", resp.Records.Locations[0].LocationsName),
		fmt.Sprint("⌚", locationData.WeatherElement[1].Time[0].StartTime, "~", locationData.WeatherElement[1].Time[0].EndTime),
		fmt.Sprint("🌡️體感溫度: ", locationData.WeatherElement[0].Time[0].ElementValue[0].Value, locationData.WeatherElement[0].Time[0].ElementValue[0].Measures),
		fmt.Sprint("ℹ️", locationData.WeatherElement[1].Time[0].ElementValue[0].Value),
	}

	messages = append(messages, linebot.NewTextMessage(strings.Join(content, "\n")))

	return messages
}

func (self *WeatherForecast) getWeatherForecast() (*WeatherForecastResp, error) {
	tw, _ := time.LoadLocation("Asia/Taipei")
	current := time.Now().In(tw)
	current3 := current.Add(time.Hour * 3)

	resp, err := requests.Get(
		"https://opendata.cwb.gov.tw/api/v1/rest/datastore/"+CountyNameCwbIdMapping[self.CountyName],
		map[string]string{
			"Authorization": os.Getenv("CWB_AUTH_CODE"),
			"elementName":   "AT,WeatherDescription",
			"locationName":  self.LocationName,
			"timeFrom":      current.Format("2006-01-02T15:04:05"),
			"timeTo":        current3.Format("2006-01-02T15:04:05"),
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

// https://opendata.cwb.gov.tw/opendatadoc/CWB_Opendata_API_V1.2.pdf
var CountyNameCwbIdMapping = map[string]string{
	"宜蘭縣": "F-D0047-001",
	"桃園市": "F-D0047-005",
	"新竹縣": "F-D0047-009",
	"苗栗縣": "F-D0047-013",
	"彰化縣": "F-D0047-017",
	"南投縣": "F-D0047-021",
	"雲林縣": "F-D0047-025",
	"嘉義縣": "F-D0047-029",
	"屏東縣": "F-D0047-033",
	"臺東縣": "F-D0047-037",
	"台東縣": "F-D0047-037",
	"花蓮縣": "F-D0047-041",
	"澎湖縣": "F-D0047-045",
	"基隆市": "F-D0047-049",
	"新竹市": "F-D0047-053",
	"嘉義市": "F-D0047-057",
	"臺北市": "F-D0047-061",
	"台北市": "F-D0047-061",
	"高雄市": "F-D0047-065",
	"新北市": "F-D0047-069",
	"臺中市": "F-D0047-073",
	"台中市": "F-D0047-073",
	"臺南市": "F-D0047-077",
	"台南市": "F-D0047-077",
	"連江縣": "F-D0047-081",
	"金門縣": "F-D0047-085",
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
