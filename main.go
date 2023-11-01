package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
    fmt.Println("File Upload Endpoint Hit")

    // Parse our multipart form, 10 << 20 specifies a maximum
    // upload of 10 MB files.
    r.ParseMultipartForm(10 << 20)
    // FormFile returns the first file for the given key `myFile`
    // it also returns the FileHeader so we can get the Filename,
    // the Header and the size of the file
    file, handler, err := r.FormFile("myFile")
    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }
    defer file.Close()
    fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

    // Create a temporary file within our temp-images directory that follows
    // a particular naming pattern
    tempFile, err := os.Create("temp-images/upload-" + handler.Filename)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer tempFile.Close()

    // Copy the uploaded file to the temporary file
    _, err = io.Copy(tempFile, file)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Return that we have successfully uploaded our file!
    fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func setupRoutes() {
    http.HandleFunc("/upload", uploadFile)
    http.ListenAndServe(":8080", nil)
}

func main() {
    fmt.Println("Hello World")
    setupRoutes()
}
