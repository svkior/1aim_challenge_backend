package main

import (
  "log"
  "net/http"
  "flag"
)


func main(){
  httpHost := flag.String("http", ":8090", "Description")
  flag.Parse()
  log.Printf("Starting API server on %s", *httpHost)
  http.HandleFunc("/getrooms", getrooms)
  http.Handle("/", http.FileServer(http.Dir("./build")))
  log.Fatal(http.ListenAndServe(*httpHost, nil))
}
