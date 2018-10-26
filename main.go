package main

import (
  "net/http"
)

func main() {
  createUser("jliu08", "hellomynameisjames", 2)
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
