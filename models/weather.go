package models

import (
	"encoding/json"
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

type Weather struct {
	TemperatureData []float32 `json:"temperature"`
	PressureData    []uint16  `json:"pressure"`
	HumidityData    []uint16  `json:"humidity"`
}

type TemperatureData struct {
	Temperature []float32 `json:"temperature"`
}

type PressureData struct {
	Pressure []uint16 `json:"pressure"`
}

type HumidityData struct {
	Humidity []uint16 `json:"humidity"`
}

func NewWeather() *Weather {
	return &Weather{}
}

func (weather *Weather) UpdateWeatherByTopic(topic TopicType, data []byte) {
	switch topic {
	case UserTemperatureTopic:
		temp := Temperature{}
		json.Unmarshal([]byte(data), &temp)
		weather.TemperatureData = append(weather.TemperatureData, temp.Temperature)

	case UserPressureTopic:
		press := Pressure{}
		json.Unmarshal([]byte(data), &press)
		weather.PressureData = append(weather.PressureData, press.Pressure)

	case UserHumidityTopic:
		hum := Humidity{}
		json.Unmarshal([]byte(data), &hum)
		weather.HumidityData = append(weather.HumidityData, uint16(hum.Humidity))

	}
}
