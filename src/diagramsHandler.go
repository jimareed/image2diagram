package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

const maxDiagramUploadSize = 8 * 1024 * 1024 // 8 mb
const diagramUploadPath = "./tmp"
const diagramTokenLength = 12

func createDiagramsHandler(w http.ResponseWriter, r *http.Request) {
	// validate file size
	r.Body = http.MaxBytesReader(w, r.Body, maxDiagramUploadSize)
	if err := r.ParseMultipartForm(maxDiagramUploadSize); err != nil {
		diagramRenderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
		return
	}

	// parse and validate file and post parameters
	fileType := r.PostFormValue("type")
	file, _, err := r.FormFile("uploadFile")
	if err != nil {
		diagramRenderError(w, "INVALID_FILE", http.StatusBadRequest)
		return
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		diagramRenderError(w, "INVALID_FILE", http.StatusBadRequest)
		return
	}

	// check file type, detectcontenttype only needs the first 512 bytes
	filetype := http.DetectContentType(fileBytes)
	switch filetype {
	case "image/jpeg", "image/jpg":
	case "image/gif", "image/png":
	case "application/pdf":
		break
	default:
		diagramRenderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
		return
	}
	fileName := diagramRandToken(diagramTokenLength)
	fileEndings, err := mime.ExtensionsByType(fileType)
	if err != nil {
		diagramRenderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
		return
	}
	newPath := filepath.Join(diagramUploadPath, fileName+fileEndings[0])
	fmt.Printf("FileType: %s, File: %s\n", fileType, newPath)

	// write file
	newFile, err := os.Create(newPath)
	if err != nil {
		diagramRenderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
		return
	}
	defer newFile.Close() // idempotent, okay to call twice
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		diagramRenderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
		return
	}
	
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	io.WriteString(w, "{\"id\": \""+fileName+"\"}\n")
}

func diagramRenderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func diagramRandToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func getDiagramsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, "[\n")
	files, err := ioutil.ReadDir(diagramUploadPath)
	if err != nil {
		log.Fatal(err)
	}
	separator := ""
	for _, f := range files {
		io.WriteString(w, separator)
		if separator == "" {
			separator = ","
		}
		fileparts := strings.Split(f.Name(), ".")
		io.WriteString(w, "{\"id\": \""+fileparts[0]+"\"}")
		fmt.Println()
	}
	io.WriteString(w, "]\n")
}

func getDiagramHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, "{\"id\": \""+id+"\"}\n")
}

