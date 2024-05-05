package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendJSONResponse(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("error while marshalling json: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func SendErrorResponse(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Internal server error (5XX): %v", msg)
	}
	type ErrorResponse struct {
		Error string `json:"error"`
	}

	SendJSONResponse(w, code, ErrorResponse{Error: msg})
}
