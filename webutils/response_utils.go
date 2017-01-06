package webutils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
)

func RespondWithJSON(w http.ResponseWriter, data interface{}, internalErr bool) {
	if internalErr {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if isNil(data) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, `{}`)
		return
	}
	w.WriteHeader(http.StatusOK)
	encoderErr := json.NewEncoder(w).Encode(data)
	if encoderErr != nil {
		log.Printf("error while encoding json: %v\n", encoderErr)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func isNil(a interface{}) bool {
	defer func() { recover() }()
	return a == nil || reflect.ValueOf(a).IsNil()
}
