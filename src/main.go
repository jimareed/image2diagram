package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/images", createImages).Methods("POST")
	r.HandleFunc("/images/{id}/diagram", getDiagram).Methods("GET")
	r.HandleFunc("/images/{id}", getImage).Methods("GET")
	r.HandleFunc("/images", getImages).Methods("GET")

	log.Print("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
