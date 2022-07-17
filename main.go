package main

import (
	"net/http"

	"github.com/dileep9490/todoapp/Backend/database"
	"github.com/dileep9490/todoapp/Backend/handlers"
	"github.com/gorilla/mux"
)

func main() {

	database.Connect()

	route := mux.NewRouter()

	route.HandleFunc("/auth/signup", handlers.SignUP).Methods("POST")
	route.HandleFunc("/auth/login", handlers.Login).Methods("POST")
	http.ListenAndServe(":8080", route)

}
