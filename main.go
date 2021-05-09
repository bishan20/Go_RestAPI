package main

import (
	"FirstGo/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/users", user.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", user.GetUser).Methods("GET")
	r.HandleFunc("/users", user.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", user.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", user.DeleteUser).Methods("Delete")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	user.InitialMigration()
	initializeRouter()
}
