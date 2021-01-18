package main

import (
	"fmt"
	"net/http"
)

func uploadFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading File")
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFiles)
	http.ListenAndServe(":8080", nil)
}
func main() {
	fmt.Println("Golang File Upload")
	setupRoutes()
}
