package main

import (
	"database/sql"
	"log"
	"net/http"
)

type Server struct {
	listenAddress string
	dbConnection  *sql.DB
}

func NewServer(addr string, conn *sql.DB) Server {
	return Server{
		listenAddress: addr,
		dbConnection:  conn,
	}
}

func (s Server) run() {
	route := http.NewServeMux()
	route.HandleFunc("/health", s.handlerHealth)
	log.Println("Server listening")
	if err := http.ListenAndServe(s.listenAddress, route); err != nil {
		log.Fatal(err)
	}
}

func (s Server) handlerHealth(w http.ResponseWriter, r *http.Request) {
	WriteResponse(w, http.StatusOK, "This is the health handler")
}
