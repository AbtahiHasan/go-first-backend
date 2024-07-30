package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func errorFormatter(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error:", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	responseFormatter(w, code, errResponse{Error: msg})
}
func responseFormatter(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	w.Write(data)

}