package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

	UndefindedTopic TopicType = ""
)

var UserTopics = [...]TopicType{UserTemperatureTopic, UserPressureTopic, UserHumidityTopic}

func StringToTopic(str string) TopicType {
	switch str {
	case "/api/v1/user/temperature":
		return UserTemperatureTopic
	case "/api/v1/user/humidity":
		return UserHumidityTopic
	case "/api/v1/user/pressure":
		return UserPressureTopic
	}
	return UndefindedTopic
}

type Container struct {
	MQTTContainer    map[string]*MQTTSubscriber
	WeatherContainer map[string]*Weather
}

var globalContainer = Container{}

func InitGlobalContainer() {
	globalContainer.MQTTContainer = make(map[string]*MQTTSubscriber)
	globalContainer.WeatherContainer = make(map[string]*Weather)
}

func NewUserContext(user User) *UserContext {
	weather := &Weather{}
	uctx := &UserContext{User: user}
	subsciber := NewMQTTSubscriberConfig("user", UserTopics[:], weather)
	//go goMQTT(subsciber)
	globalContainer.MQTTContainer[uctx.User.Password] = subsciber
	globalContainer.WeatherContainer[uctx.User.Password] = weather
	return uctx
}

func (ctx UserContext) GetWeather() *Weather {
	if weather, ok := globalContainer.WeatherContainer[ctx.User.Password]; !ok {
		return nil
	} else {
		return weather
	}
}

/* TODO:
 *	1. Rewrite json API
 *	2. Rework&change storage
 *  3. Remove Key
 *  4. Add secure MQTT
 *  5. Add logout
 *	6. Add destroying of objects
 *	7. Add more api and admin, healt
 */
