package models

type Temperature struct {
	Temp float32 `json:"temperature"`
}

type Pressure struct {
	Pressure uint16 `json:"pressure"`
}

type Humidity struct {
	Pressure uint8 `json:"humidity"`
}
