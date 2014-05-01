package main

import (
  "fmt"
  "log"
  "net/http"
  "io"
  "io/ioutil"
  "os"
)

var uploadPath string = ""

func main() {
  setupUploadPath()

  log.Println( "Upload path is set to: " + uploadPath )

  setupHandlers()

  fmt.Println("Starting server...")
  err := http.ListenAndServe(":1337", nil)

  if err != nil {
    log.Fatal( "ListenAndServe: ", err )
  }
}

/** Setup **/

func setupUploadPath() {
  uploadPath = os.Args[1]

  // Remove trailing slash
  if len(uploadPath) > 0 && string( uploadPath[ len(uploadPath) - 1 ] ) == "/" {
    log.Println("Removing trailing slash.")
    uploadPath = uploadPath[ :len(uploadPath) - 1 ]
  }

  fi, err :=  os.Stat( uploadPath )

  if err != nil {
    log.Panic( err )
  }

  if !fi.Mode().IsDir() {
    log.Println( "Upload path ( %s ) does not exist.", uploadPath )
    os.Exit(1)
  }
}

func setupHandlers() {
  http.HandleFunc("/", Index)
  http.HandleFunc("/drop", Drop)
}

/** HANDLERS **/

//Index displays a upload page for testing file-bucket
func Index( w http.ResponseWriter, req *http.Request) {
  page, _ := ioutil.ReadFile( "index.html" )
  fmt.Fprintf(w, "%s", page)
}

//Drop accepts POST requests containing a parameter named "file" and saves it to the upload path
func Drop( w http.ResponseWriter, r *http.Request) {
  log.Println("/drop")

  // Only allow POST methods
  if r.Method != "POST" {
    http.Error(w, "Invalid HTTP method", http.StatusInternalServerError)
    return
  }


  file, fileHeader, err := r.FormFile("file")

  if err != nil {
    log.Println(err)
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  filePath := uploadPath + "/" + fileHeader.Filename

  dst, err := os.Create( filePath )
  defer dst.Close()

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if _, err := io.Copy(dst, file); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  http.Redirect(w, r, "/", http.StatusFound)
}
