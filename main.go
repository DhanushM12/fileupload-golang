package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func uploadFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading File\n")
	//1. parse input, type multipart/form-data
	r.ParseMultipartForm(10 << 20)
	//2. retrieve file from posted form-data
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File from form-data")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	//3. write temp file on our server
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	//4. return whether or not this has been successful
	fmt.Fprintf(w, "Successfully Uploaded File\n")

}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFiles)
	http.ListenAndServe(":8080", nil)
}
func main() {
	fmt.Println("Golang File Upload")
	setupRoutes()
}
