package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "os"
  "log"
  "fmt"
)

func main() {
  fmt.Println("Starting server")
  r := mux.NewRouter()

  http.Handle("/", r)

  srv := &http.Server {
    Handler: r,
    Addr:    ":" + os.Getenv("PORT"),
  }

  log.Fatal(srv.ListenAndServe())
}