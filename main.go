package main

import (
  "net/http"
  "fmt"
)

func main() {
  createUser("jliu08", "hellomynameisjames", 2)
  // test login function

  loginSuccess, _ := login("jliu09", "hellomynameisnotjames")
  if loginSuccess == 1 {
    fmt.Println("Login succeeded")
  } else {
    fmt.Println("Login failed or error occurred")
  }
  changePassword("jliu08", "hellomynameisnotjames", "hellomynameisjames")
  loginSuccess, _ = login("jliu08", "hellomynameisnotjames")

  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}