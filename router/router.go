package router

import (
	"github.com/gorilla/mux"
	"server/handlers"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/newGame", handlers.NewGame).Methods("POST")
	r.HandleFunc("/joinGame/{gameID}", handlers.JoinGame).Methods("POST")
	r.HandleFunc("/list", handlers.List).Methods("GET")
	r.HandleFunc("/remove", handlers.Remove).Methods("DELETE")
	return r
}
