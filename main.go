package main

import (
  "net/http"
  "fmt"
  "Trans-Law-Center/assets"
)

func main() {
  // Login Testing

  assets.createUser("jliu08", "hellomynameisjames", 2)
  
  loginSuccess, _ := assets.login("jliu09", "hellomynameisnotjames")
  if loginSuccess == 1 {
    fmt.Println("Login succeeded")
  } else {
    fmt.Println("Login failed or error occurred")
  }
  assets.changePassword("jliu08", "hellomynameisnotjames", "hellomynameisjames")
  loginSuccess, _ = assets.login("jliu08", "hellomynameisnotjames")

  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)

  // Dummy response testing



}
