package main

import (
  "net/http"
  //"fmt"
  "log"
  "Trans-Law-Center/src"
)

func main() {
    //CREATE DB
    src.SetupFormDB("formdb.db")
    src.SetupLoginDB("database.db")

    ////FOR Login Testing
    // src.CreateUser("jliu08", "hellomynameisjames", 2)
    //
    // loginSuccess, _ := src.Login("jliu08", "hellomynameisjames")
    // if loginSuccess == 1 {
    // fmt.Println("Login succeeded")
    // } else {
    // fmt.Println("Login failed or error occurred")
    // }
    // src.ChangePassword("jliu08", "hellomynameisnotjames", "hellomynameisjames")
    // loginSuccess, _ = src.Login("jliu08", "hellomynameisnotjames")

    //FOR Testing STATIC Pages.
    // fmt.Println("Loading server on :8080")
    // fs := http.FileServer(http.Dir("html"))
    // http.Handle("/", fs)

    err := src.LoadPresetDBContent("formdb.db")
    if err != nil {
      log.Fatal("Loading Preset...: ", err)
    }

    //

    http.HandleFunc("/home/", src.ViewHandler)
    http.HandleFunc("/results/", src.ResultsHandler)

    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
    http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

    err = http.ListenAndServe(":8080", nil)
    if err != nil {
      log.Fatal("ListenAndServe: ", err)
    }
}
