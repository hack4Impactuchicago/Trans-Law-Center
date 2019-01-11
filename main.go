package main

import (
  "net/http"
  "fmt"
  "Trans-Law-Center/assets"
)

func main() {
/*
  // Login Testing
  assets.CreateUser("jliu08", "hellomynameisjames", 2)

  loginSuccess, _ := assets.Login("jliu08", "hellomynameisjames")
  if loginSuccess == 1 {
    fmt.Println("Login succeeded")
  } else {
    fmt.Println("Login failed or error occurred")
  }
  // assets.ChangePassword("jliu08", "hellomynameisnotjames", "hellomynameisjames")
  // loginSuccess, _ = assets.Login("jliu08", "hellomynameisnotjames")
  */
  fmt.Println("Loading server on :8080")

  http.HandleFunc("/", assets.Handler)
  http.ListenAndServe(":8080", nil)

  // Dummy response testing


}
