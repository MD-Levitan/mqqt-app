package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MD-Levitan/mqqt-app/config"
	"github.com/MD-Levitan/mqqt-app/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func getStringFromRequestQuery(param string, r *http.Request) (string, error) {
	logrus.Error(mux.Vars(r))
	if value, ok := mux.Vars(r)[param]; ok {
		return value, nil
	}
	return "", fmt.Errorf("cannot find such param \"%s\"", param)
}

func statusHandler(dec *json.Decoder, enc *json.Encoder, w http.ResponseWriter, r *http.Request) (err error) {
	session, err := config.GetStore().Get(r, "Rcookie")
	if err != nil {
		return err
	}

	context := getUserContext(session)
	logrus.Error(context.GetWeather())
	enc.Encode(models.Temperature{123.12})
	return nil
}

func loginHandler(dec *json.Decoder, enc *json.Encoder, w http.ResponseWriter, r *http.Request) (err error) {
	conf := config.GetConfig()
	session, err := config.GetStore().Get(r, "Rcookie")
	if err != nil {
		return err
	}

	user := models.User{}
	if err := dec.Decode(&user); err != nil {
		return err
	}

	password, err := encrypt([]byte(conf.Web.SessionKey), []byte(user.Password))
	if err != nil {
		logrus.Error(err)
		return err
	}
	user.Password = string(password)
	session.Values["Context"] = models.NewUserContext(user)
	if err := session.Save(r, w); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func userTemperatureHandler(dec *json.Decoder, enc *json.Encoder, w http.ResponseWriter, r *http.Request) (err error) {
	session, err := config.GetStore().Get(r, "Rcookie")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return err
	}
	context := getUserContext(session)
	enc.Encode(context.GetWeather().Temp.LTemperature)
	return nil
}

func userPressureHandler(dec *json.Decoder, enc *json.Encoder, w http.ResponseWriter, r *http.Request) (err error) {
	session, err := config.GetStore().Get(r, "Rcookie")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return err
	}
	context := getUserContext(session)
	enc.Encode(context.GetWeather().Press.LPressure)
	return nil
}

func userHumidityHandler(dec *json.Decoder, enc *json.Encoder, w http.ResponseWriter, r *http.Request) (err error) {
	session, err := config.GetStore().Get(r, "Rcookie")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return err
	}
	context := getUserContext(session)
	enc.Encode(context.GetWeather().Hum.LHumidity)
	return nil
}
