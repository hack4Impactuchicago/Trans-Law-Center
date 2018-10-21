package main

import(
  "fmt"
  "net/http"
  "html/template"

)

func handler(w http.ResponseWriter, req *http.Request){
  t, err := template.New("home.html").ParseFiles("html/home.html")
  if err != nil {
    fmt.Println(err)
  }
  t.Execute(w, nil)
}
