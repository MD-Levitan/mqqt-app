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

func userStatusHandler(dec *json.Decoder, enc *json.Encoder, w http.ResponseWriter, r *http.Request) (err error) {
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
	session, err := config.GetStore().New(r, "Rcookie")
	if err != nil {
		return err
	}

	user := models.User{}
	if err := dec.Decode(&user); err != nil {
		return err
	}

	ctx := models.NewUserContext(user)
	if ctx == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	session.Values["Context"] = ctx
	if err := session.Save(r, w); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	return nil
}

func logout(w http.ResponseWriter, r *http.Request) (err error) {
	session, err := config.GetStore().Get(r, "Rcookie")
	if err != nil || session.IsNew {
		w.WriteHeader(http.StatusUnauthorized)
		return err
	}
	context := getUserContext(session)
	models.DeleteUserContext(context)
	delete(session.Values, "Context")
	session.Options.MaxAge = -1

	if err := session.Save(r, w); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	return nil
}

func logoutHandler(dec *json.Decoder, enc *json.Encoder, w http.ResponseWriter, r *http.Request) (err error) {
	result := logout(w, r)
	return result
}

func userTemperatureHandler(dec *json.Decoder, enc *json.Encoder, w http.ResponseWriter, r *http.Request) (err error) {
	session, err := config.GetStore().Get(r, "Rcookie")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return err
	}
	context := getUserContext(session)
	if context != nil {
		if weather := context.GetWeather(); weather != nil {
			enc.Encode(&models.TemperatureData{weather.TemperatureData})
			return nil
		}
	}
	w.WriteHeader(http.StatusForbidden)
	return nil
}

func userPressureHandler(dec *json.Decoder, enc *json.Encoder, w http.ResponseWriter, r *http.Request) (err error) {
	session, err := config.GetStore().Get(r, "Rcookie")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return err
	}
	context := getUserContext(session)
	if context != nil {
		if weather := context.GetWeather(); weather != nil {
			enc.Encode(&models.PressureData{weather.PressureData})
			return nil
		}
	}
	w.WriteHeader(http.StatusForbidden)
	return nil
}

func userHumidityHandler(dec *json.Decoder, enc *json.Encoder, w http.ResponseWriter, r *http.Request) (err error) {
	session, err := config.GetStore().Get(r, "Rcookie")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return err
	}
	context := getUserContext(session)

	if context != nil {
		if weather := context.GetWeather(); weather != nil {
			enc.Encode(&models.HumidityData{weather.HumidityData})
			return nil
		}
	}
	w.WriteHeader(http.StatusForbidden)
	return nil
}

func userWeatherHandler(dec *json.Decoder, enc *json.Encoder, w http.ResponseWriter, r *http.Request) (err error) {
	session, err := config.GetStore().Get(r, "Rcookie")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return err
	}
	context := getUserContext(session)
	if context != nil {
		if weather := context.GetWeather(); weather != nil {
			enc.Encode(weather)
			return nil
		}
	}
	w.WriteHeader(http.StatusForbidden)
	return nil
}
