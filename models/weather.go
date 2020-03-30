package models

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

type Temperature struct {
	Temperature float32 `json:"temperature"`
}

type Pressure struct {
	Pressure uint16 `json:"pressure"`
}

type Humidity struct {
	Humidity uint8 `json:"humidity"`
}

type TemperatureData struct {
	LTemperature []Temperature `json:"temperature"`
}

type PressureData struct {
	LPressure []Pressure `json:"pressure"`
}

type HumidityData struct {
	LHumidity []Humidity `json:"humidity"`
}

type Weather struct {
	Temp  TemperatureData
	Press PressureData
	Hum   HumidityData
}

func NewWeather() *Weather {
	return &Weather{}
}

func (weather *Weather) UpdateWeatherByTopic(topic TopicType, data []byte) {
	switch topic {
	case UserTemperatureTopic:
		temp := Temperature{}
		json.Unmarshal([]byte(data), &temp)
		weather.Temp.LTemperature = append(weather.Temp.LTemperature, temp)
		logrus.Print(&weather)
		logrus.Print(weather)
		logrus.Print(temp)
		logrus.Print(weather.Temp.LTemperature)

	case UserPressureTopic:
		press := Pressure{}
		json.Unmarshal([]byte(data), &press)
		weather.Press.LPressure = append(weather.Press.LPressure, press)

	case UserHumidityTopic:
		hum := Humidity{}
		json.Unmarshal([]byte(data), &hum)
		weather.Hum.LHumidity = append(weather.Hum.LHumidity, hum)

	}
}
