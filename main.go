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
    src.CreateUser("testuser", "testpsw", 2)

    //// For DB TESTING
    err := src.LoadPresetDBContent("formdb.db")
    if err != nil {
      log.Fatal("Loading Preset...: ", err)
    }

    //VIEWS

    http.HandleFunc("/home/", src.ViewHandler)
    http.HandleFunc("/results/", src.ResultsHandler)
    http.HandleFunc("/admin/", src.AdminHandler)

    //CONTENT

    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
    http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

    //SERVER

    err = http.ListenAndServe(":8080", nil)
    if err != nil {
      log.Fatal("ListenAndServe: ", err)
    }
}
