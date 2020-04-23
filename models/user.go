package models

import (
	"fmt"
	"math/rand"

	"github.com/MD-Levitan/mqqt-app/config"
	"github.com/sirupsen/logrus"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ID       string `json:"ID"`
}

type UserContext struct {
	User User
}

type Error struct {
	Error string `json:"error"`
}

type TopicType string

const (
	UserTemperatureTopic TopicType = "/api/v1/user/temperature"
	UserHumidityTopic    TopicType = "/api/v1/user/humidity"
	UserPressureTopic    TopicType = "/api/v1/user/pressure"
	UserInfoTopic        TopicType = "/api/v1/user/info"
	UserStatusTopic      TopicType = "/api/v1/user/status"

	AdminTimeTopic    TopicType = "/api/v1/admin/work_time"
	AdminBatteryTopic TopicType = "/api/v1/admin/battery"
	AdminInfoTopic    TopicType = "/api/v1/admin/info"
	AdminModelTopic   TopicType = "/api/v1/admin/model"

	UndefindedTopic TopicType = ""
)

var UserTopics = [...]TopicType{UserTemperatureTopic, UserPressureTopic, UserHumidityTopic,
	UserInfoTopic, UserStatusTopic}
var AdminTopics = [...]TopicType{UserTemperatureTopic, UserPressureTopic, UserHumidityTopic,
	UserInfoTopic, UserStatusTopic, AdminTimeTopic, AdminBatteryTopic,
	AdminInfoTopic, AdminModelTopic}

func StringToTopic(str string) TopicType {
	switch str {
	case "/api/v1/user/temperature":
		return UserTemperatureTopic
	case "/api/v1/user/humidity":
		return UserHumidityTopic
	case "/api/v1/user/pressure":
		return UserPressureTopic
	case "/api/v1/user/info":
		return UserInfoTopic
	case "/api/v1/user/status":
		return UserStatusTopic
	case "/api/v1/admin/work_time":
		return AdminTimeTopic
	case "/api/v1/admin/battery":
		return AdminBatteryTopic
	case "/api/v1/admin/info":
		return AdminInfoTopic
	case "/api/v1/admin/model":
		return AdminModelTopic

	}
	return UndefindedTopic
}

type Container struct {
	MQTTContainer    map[string]*MQTTSubscriber
	WeatherContainer map[string]*Weather
	DeviceContainer  map[string]*Device
	UsersContainer   map[string]string
}

var globalContainer = Container{}

func InitGlobalContainer() {
	globalContainer.MQTTContainer = make(map[string]*MQTTSubscriber)
	globalContainer.WeatherContainer = make(map[string]*Weather)
	globalContainer.DeviceContainer = make(map[string]*Device)
	globalContainer.UsersContainer = make(map[string]string)
}

func NewUserContext(user User) *UserContext {
	/* If Multiple Users is not allowed check all new users and delete old in case if such user exsit  */
	if config.GetConfig().MQQT.MultipleUsers == false {
		if userID, ok := globalContainer.UsersContainer[user.Username]; ok {
			deleteUser(userID)
		}
	}
	user.ID = fmt.Sprintf("mqtt-app-%d", rand.Int63())
	weather := &Weather{}
	device := &Device{}
	subsciber := NewMQTTSubscriberConfig(user, AdminTopics[:], weather, device)

	password, err := encrypt([]byte(config.GetConfig().Web.SessionKey), []byte(user.Password))
	if err != nil {
		logrus.Error(err)
		return nil
	}
	user.Password = string(password)
	uctx := &UserContext{User: user}
	if subsciber == nil {
		return nil
	}

	globalContainer.MQTTContainer[uctx.User.ID] = subsciber
	globalContainer.WeatherContainer[uctx.User.ID] = weather
	globalContainer.DeviceContainer[uctx.User.ID] = device
	if config.GetConfig().MQQT.MultipleUsers == false {
		globalContainer.UsersContainer[user.Username] = uctx.User.ID
	}
	return uctx
}

func deleteUser(ID string) {
	/* Remove Weather */
	delete(globalContainer.WeatherContainer, ID)
	delete(globalContainer.DeviceContainer, ID)

	/* Remove Subcriber */
	if subsciber, ok := globalContainer.MQTTContainer[ID]; !ok || subsciber == nil {
		return
	} else {
		subsciber.Client.Disconnect(0)
		delete(globalContainer.MQTTContainer, ID)
	}
}

func DeleteUserContext(uctx *UserContext) {
	deleteUser(uctx.User.ID)
}

func (uctx UserContext) CheckUser() bool {
	_, ok := globalContainer.MQTTContainer[uctx.User.ID]
	return ok
}

func (uctx UserContext) GetWeather() *Weather {
	if weather, ok := globalContainer.WeatherContainer[uctx.User.ID]; !ok {
		return nil
	} else {
		return weather
	}
}

func (uctx UserContext) GetDevice() *Device {
	if device, ok := globalContainer.DeviceContainer[uctx.User.ID]; !ok {
		return nil
	} else {
		return device
	}
}

/* TODO:
 *	2. Rework&change storage - DELAYED(check of user with same names, checkin expiration)
 *  3. Remove Key - IN WORK
 *	6. Add destroying of objects - PARTIALY DONE
 *	7. Add more api and admin, healt
 *  8. Rework UserContext
 */
