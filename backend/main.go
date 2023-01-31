package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "os"
  "log"
)

func main() {
  r := mux.NewRouter()

  http.Handle("/", r)

  srv := &http.Server {
    Handler: r,
    Addr:    ":" + os.Getenv("PORT"),
  }

  log.Fatal(srv.ListenAndServe())
}