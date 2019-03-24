package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func getDiagramHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, image2Diagram(getImageFileName(id)))
}
