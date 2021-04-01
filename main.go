package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", controller.LoginHandler).
		Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
