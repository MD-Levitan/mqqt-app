package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MD-Levitan/mqqt-app/models"
	"github.com/gorilla/mux"
)

func getStringFromRequestQuery(param string, r *http.Request) (string, error) {
	if value, ok := mux.Vars(r)[param]; ok {
		return value, nil
	}
	return "", fmt.Errorf("cannot find such param \"%s\"", param)
}

func statusHandler(dec *json.Decoder, enc *json.Encoder, w http.ResponseWriter, r *http.Request) (err error) {
	enc.Encode(models.Temperature{123.12})
	return nil
}

func loginHandler(dec *json.Decoder, enc *json.Encoder, w http.ResponseWriter, r *http.Request) (err error) {

	var password string
	var username string

	if password, err = getStringFromRequestQuery("password", r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	if username, err = getStringFromRequestQuery("username", r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	if jwt, err := createJWT(username, password); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	} else {
		enc.Encode(models.JWT{jwt})
		return nil
	}
}

func userTemperatureHandler(dec *json.Decoder, enc *json.Encoder, w http.ResponseWriter, r *http.Request) (err error) {
	enc.Encode(models.Temperature{123.12})
	return nil
}

func userPressureHandler(dec *json.Decoder, enc *json.Encoder, w http.ResponseWriter, r *http.Request) (err error) {
	enc.Encode(models.Pressure{730})
	return nil
}

func userHumidityHandler(dec *json.Decoder, enc *json.Encoder, w http.ResponseWriter, r *http.Request) (err error) {
	enc.Encode(models.Humidity{80})
	return nil
}
