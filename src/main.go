package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/images", createImagesHandler).Methods("POST")
	r.HandleFunc("/images/{id}/diagram", getDiagramHandler).Methods("GET")
	r.HandleFunc("/images/{id}", getImageHandler).Methods("GET")
	r.HandleFunc("/images", getImagesHandler).Methods("GET")

	log.Print("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
