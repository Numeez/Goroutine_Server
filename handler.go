package main

import (
	"database/sql"
	"encoding/json"
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
	if err := s.makeUserTable(); err != nil {
		log.Fatal(err)
	}

	route := http.NewServeMux()
	route.HandleFunc("/getAllUsers", s.handlerGetAllUsers)
	route.HandleFunc("/addUser", s.handlerAddUser)
	route.HandleFunc("/health", s.handlerHealth)
	log.Println("Server listening")
	if err := http.ListenAndServe(s.listenAddress, route); err != nil {
		log.Fatal(err)
	}
}

func (s Server) handlerHealth(w http.ResponseWriter, r *http.Request) {
	WriteResponse(w, http.StatusOK, "This is the health handler", nil)
}

func (s Server) handlerAddUser(w http.ResponseWriter, r *http.Request) {
	var request User
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	if err := s.insertUser(request); err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	WriteResponse(w, http.StatusOK, "user inserted successfully", nil)
}

func (s Server) handlerGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.getAllUsers()
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	WriteResponse(w, http.StatusOK, "users fetched successfully", users)

}
