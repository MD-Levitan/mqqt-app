package main

import (
	"fmt"
	"net/http"

	"github.com/MD-Levitan/mqqt-app/router"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
)

type MQTTSubscriber struct {
	Options *MQTT.ClientOptions
	Client  MQTT.Client
}

func NewMQTTSubscriber(protocol string, host string, port uint16, clientID string, topic string) *MQTTSubscriber {
	var handler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
		logrus.Printf("MSG: %s\n", msg.Payload())
	}

	options := MQTT.NewClientOptions().AddBroker(fmt.Sprintf("%s://%s:%d", protocol, host, port))
	options.SetClientID(clientID)
	options.SetDefaultPublishHandler(handler)

	options.OnConnect = func(c MQTT.Client) {
		if token := c.Subscribe(topic, 0, handler); token.Wait() && token.Error() != nil {
			logrus.Error(token.Error())
		}
	}
	client := MQTT.NewClient(options)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logrus.Error(token.Error())
	} else {
		logrus.Printf("Connected to server\n")
	}

	sub := &MQTTSubscriber{Options: options, Client: client}
	return sub
}

func (subcriber MQTTSubscriber) wait() {
	logrus.Print("%s", subcriber.Client)
}

// func main() {
// 	c := make(chan os.Signal, 1)
// 	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

// 	subscriber := NewMQTTSubscriber("tcp", "172.17.0.2", 1883, "User", "/api/v1/#")
// 	subscriber.wait()

// 	<-c
// }

func main() {
	http.ListenAndServe(":10000", router.MakeRouter())
}
