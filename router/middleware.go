package router

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSONHandler(handler func(dec *json.Decoder, enc *json.Encoder, w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		encoder := json.NewEncoder(w)
		if err := handler(decoder, encoder, w, r); err != nil {
			w.WriteHeader(400)
			//encoder.Encode(models.Error{err.Error()})
			return
		}
	}
}

func authHandler(dec *json.Decoder, enc *json.Encoder, w http.ResponseWriter, r *http.Request) (err error) {
	fmt.Printf("authHandler")
	return fmt.Errorf("test error")
}
