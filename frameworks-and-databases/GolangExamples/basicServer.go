package main

import (
  "io"
  "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "Hello World!")
}

//Http package has a default *ServeMux. Here, we define our own
func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", hello)
  http.ListenAndServe(":3000", mux)
}
