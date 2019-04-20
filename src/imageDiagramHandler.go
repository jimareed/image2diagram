package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func getImageDiagramHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	bwImage := image2BlackWhiteImage(getImageFileName(id))
	diagram := blackWhiteImage2Diagram(bwImage)

	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, diagram2String(diagram))
}
