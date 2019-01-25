package main

import (
  "net/http"
  "fmt"
  "log"
  "Trans-Law-Center/assets"
)

func main() {
    //CREATE DB
    assets.SetupFormDB("formdb.db")
    assets.SetupLoginDB("database.db")

    //FOR Login Testing
    // assets.CreateUser("jliu08", "hellomynameisjames", 2)
    //
    // loginSuccess, _ := assets.Login("jliu08", "hellomynameisjames")
    // if loginSuccess == 1 {
    // fmt.Println("Login succeeded")
    // } else {
    // fmt.Println("Login failed or error occurred")
    // }
    // assets.ChangePassword("jliu08", "hellomynameisnotjames", "hellomynameisjames")
    // loginSuccess, _ = assets.Login("jliu08", "hellomynameisnotjames")


    //FOR Testing STATIC Pages.
    // fmt.Println("Loading server on :8080")
    // fs := http.FileServer(http.Dir("html"))
    // http.Handle("/", fs)

    // http.HandleFunc("/", assets.ViewHandler)
    // http.HandleFunc("/results/", assets.ResultsHandler)
    // http.HandleFunc("/admin/, assets.SecureHandler")

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
      log.Fatal("ListenAndServe: ", err)
    }
}
