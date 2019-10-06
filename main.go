package main

import (
  "net/http"
  "log"
  "Trans-Law-Center/api/src"
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
    http.HandleFunc("/", src.ViewHandler)
    http.HandleFunc("/home/", src.ViewHandler)
    http.HandleFunc("/results/", src.ResultsHandler)
    http.HandleFunc("/admin/", src.AdminHandler)

    //CONTENT
    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("ui/css"))))
    http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("ui/images"))))

    //SERVER

    err = http.ListenAndServe(":8080", nil)
    if err != nil {
      log.Fatal("ListenAndServe: ", err)
    }
}
