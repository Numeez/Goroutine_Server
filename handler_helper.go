package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

type ErrResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

func WriteResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(statusCode)
	resp := Response{
		Message: message,
	}
	respBytes, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}
	_, _ = w.Write(respBytes)
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(statusCode)
	message := "Error occured: " + err.Error()
	resp := ErrResponse{
		ErrorMessage: message,
	}
	respBytes, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}
	_, _ = w.Write(respBytes)
}
