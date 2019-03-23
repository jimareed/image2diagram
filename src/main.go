package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/images", createImages).Methods("POST")
	r.HandleFunc("/images/{id}", getImage)
	r.HandleFunc("/images", getImages)

	log.Print("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
