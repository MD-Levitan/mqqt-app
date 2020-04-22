package main

import (
	"encoding/gob"
	"net/http"

	"github.com/MD-Levitan/mqqt-app/config"
	"github.com/MD-Levitan/mqqt-app/models"
	"github.com/MD-Levitan/mqqt-app/router"

	"github.com/sirupsen/logrus"
)

func main() {

	models.InitGlobalContainer()

	gob.Register(&models.UserContext{})
	gob.Register(&models.User{})

	if err := config.InitConfig("config.yaml"); err != nil {
		logrus.Error("Cannot open config: ", err)
		return
	}

	if err := config.InitTmpl(); err != nil {
		logrus.Error("Cannot init templates: ", err)
		return
	}

	if err := config.InitStore(); err != nil {
		logrus.Error("Cannot init store: ", err)
		return
	}
	//defer store.Checker(config.GetStore())

	http.ListenAndServe(":10000", router.MakeRouter())
}
