package webutils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, data interface{}, dataErr error) {
	if dataErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if data == nil {
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
