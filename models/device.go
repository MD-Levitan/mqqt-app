package models

import (
	"encoding/json"
	"fmt"
)

type WorkTime struct {
	Time uint32 `json:"work_time"`
}

type Battery struct {
	Battery float32 `json:"battery"`
}

type AdminInfo struct {
	Info string `json:"info"`
}

type UserInfo struct {
	Info string `json:"info"`
}

type Model struct {
	Model uint8 `json:"model"`
}

type Status struct {
	Status uint8 `json:"status"`
}

type Device struct {
	Time    WorkTime
	Battery Battery
	User    UserInfo
	Admin   AdminInfo
	Model   Model
	Status  Status
}

type UDevice struct {
	Info   string `json:"info"`
	Status string `json:"status"`
}

type ADevice struct {
	Info    string  `json:"info"`
	Status  string  `json:"status"`
	Time    uint32  `json:"time"`
	Battery float32 `json:"battery"`
	Model   uint8   `json:"model"`
}

type DeviceContainer struct {
	User  *UDevice `json:"udev"`
	Admin *ADevice `json:"adev"`
}

func statusToStr(status Status) string {
	switch status.Status {
	case 0:
		return "Disconnect"
	case 1:
		return "Connect"
	case 2:
		return "Re-connecting"
	default:
		return "Undefiend"
	}
}

func ConvertDevice(device Device) *DeviceContainer {
	fmt.Printf("Get device %v\n", device)
	container := &DeviceContainer{User: nil, Admin: nil}
	if device.User.Info != "" {
		container.User = &UDevice{Info: device.User.Info, Status: statusToStr(device.Status)}
	}
	if device.Time.Time != 0 {
		container.Admin = &ADevice{Info: device.Admin.Info, Status: statusToStr(device.Status),
			Time: device.Time.Time, Battery: device.Battery.Battery, Model: device.Model.Model,
		}
	}
	return container
}

func NewDevice() *Device {
	return &Device{}
}

func (device *Device) UpdateDeviceByTopic(topic TopicType, data []byte) {
	fmt.Printf("Update Device %v - %v\n", topic, string(data))
	switch topic {
	case UserInfoTopic:
		temp := UserInfo{}
		json.Unmarshal([]byte(data), &temp)
		device.User = temp

	case UserStatusTopic:
		temp := Status{}
		json.Unmarshal([]byte(data), &temp)
		device.Status = temp

	case AdminInfoTopic:
		temp := AdminInfo{}
		json.Unmarshal([]byte(data), &temp)
		device.Admin = temp

	case AdminModelTopic:
		temp := Model{}
		json.Unmarshal([]byte(data), &temp)
		device.Model = temp

	case AdminBatteryTopic:
		temp := Battery{}
		json.Unmarshal([]byte(data), &temp)
		device.Battery = temp

	case AdminTimeTopic:
		temp := WorkTime{}
		json.Unmarshal([]byte(data), &temp)
		device.Time = temp
	}
}
