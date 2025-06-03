package main

import (
	"log"
	"net/http"
)

func main() {
	route := http.NewServeMux()
	route.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("This is the health api"))
	})
	log.Println("Server listening")
	if err := http.ListenAndServe(":8080", route); err != nil {
		log.Fatal(err)
	}
}
