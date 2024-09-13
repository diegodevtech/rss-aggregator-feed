package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string){
	if code > 499{
		log.Println("Responding with 5XX error: ", message)
	}
	type errResponse struct {
		ErrorTitle string `json:"error"`
		ErrorDetail string `json:"detail"`
	}
	respondWithJSON(w, code, errResponse{
		ErrorTitle: "Something went wrong",
		ErrorDetail: message,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface {}){
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(data)
}