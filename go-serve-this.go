package main

import (
  "log"
  "net/http"
  "io"
  "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
  filename := r.URL.Path[1:]

  if filename == "" {
    filename = "index.html"
  }

  file, err := os.Open( filename )

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  io.Copy(w, file)
}

func main() {
  http.HandleFunc("/", handler)
  log.Println("Serving on port 8080...")
  http.ListenAndServe(":8080", nil)
}
