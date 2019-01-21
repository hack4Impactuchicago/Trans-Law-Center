package assets

import(
  "net/http"
  "html/template"
  "fmt"
  "database/sql"
  "log"
  "Trans-Law-Center/defns"
)

func loadViewPage()(*ViewPage, error){
  rows, err := AllRows("formdb.db", "Questions")
  if err != nil {
    return nil, err
  } else {
    
  }
}

func loadResponsePage()(*ResponsePage, error){

}

func ViewHandler(w http.ResponseWriter, r *http.Request){
  p, _ := loadViewPage()
  t, _ := template.ParseFiles("/html/home.html")
  t.Execute(w, p)
}

func ResultsHandler(w http.ResponseWriter, r *http.Request) {
  p, _ := loadPage()
  fmt.Fprintf(w, "This is the results page.")
}
