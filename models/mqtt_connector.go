package models

import (
	"fmt"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"

	"github.com/MD-Levitan/mqqt-app/config"
)

type MQTTSubscriber struct {
	Options *MQTT.ClientOptions
	Client  MQTT.Client
}

func NewMQTTSubscriber(protocol string, host string, port uint16, user User, topics []TopicType, weather *Weather, device *Device) *MQTTSubscriber {
	var handler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
		//logrus.Printf("MSG: %s\n", msg.Payload())
		weather.UpdateWeatherByTopic(StringToTopic(msg.Topic()), msg.Payload())
		device.UpdateDeviceByTopic(StringToTopic(msg.Topic()), msg.Payload())
	}

	options := MQTT.NewClientOptions().AddBroker(fmt.Sprintf("%s://%s:%d", protocol, host, port))
	logrus.Printf(fmt.Sprintf("%s://%s:%d", protocol, host, port))
	// Add User Settings
	options.SetClientID(user.ID)
	options.SetUsername(user.Username)
	options.SetPassword(user.Password)

	options.SetDefaultPublishHandler(handler)

	options.OnConnect = func(c MQTT.Client) {
		if token := c.SubscribeMultiple(func(topics []TopicType) map[string]byte {
			topic_map := make(map[string]byte)
			for _, topic := range topics {
				topic_map[string(topic)] = 0
			}
			return topic_map
		}(topics), handler); token.Wait() && token.Error() != nil {
			logrus.Error(token.Error())
		}
	}
	client := MQTT.NewClient(options)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logrus.Error(token.Error())
		return nil
	} else {
		logrus.Printf("Connected to server\n")
	}

	sub := &MQTTSubscriber{Options: options, Client: client}
	return sub
}

func goMQTT(subscriber *MQTTSubscriber) {
	for subscriber.Client.IsConnected() {
		time.Sleep(20)
	}
}

func NewMQTTSubscriberConfig(user User, topics []TopicType, weather *Weather, device *Device) *MQTTSubscriber {
	conf := config.GetConfig()
	return NewMQTTSubscriber(conf.MQQT.Protocol, conf.MQQT.IP, conf.MQQT.Port, user, topics, weather, device)
}
