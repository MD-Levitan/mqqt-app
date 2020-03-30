package main

import (
	"encoding/gob"
	"net/http"

	"github.com/MD-Levitan/mqqt-app/config"
	"github.com/MD-Levitan/mqqt-app/models"
	"github.com/MD-Levitan/mqqt-app/router"

	"github.com/sirupsen/logrus"
)

// func main() {
// 	c := make(chan os.Signal, 1)
// 	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

// 	subscriber := NewMQTTSubscriber("tcp", "172.17.0.2", 1883, "User", "/api/v1/#")
// 	subscriber.wait()

// 	<-c
// }

func main() {

	models.InitGlobalContainer()

	gob.Register(&models.UserContext{})
	gob.Register(&models.User{})

	if err := config.InitConfig("config.yaml"); err != nil {
		logrus.Error("Cannot open config: ", err)
		return
	}

	if err := config.InitStore(); err != nil {
		logrus.Error("Cannot init store")
		return
	}

	http.ListenAndServe(":10000", router.MakeRouter())
}
