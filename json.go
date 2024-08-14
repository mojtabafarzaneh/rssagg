package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Failed to Marshal the json response %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)

}

func responseWithErr(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Fatal("something went wrong")
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	responseWithJSON(w, code, errResponse{
		Error: msg,
	})
}
