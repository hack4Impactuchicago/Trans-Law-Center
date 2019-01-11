package assets

import(
  "net/http"
  "html/template"
  // "golang.org/x/crypto/bcrypt"
  // "github.com/mattn/go-sqlite3"
  // "database/sql"
  // "log"
  "fmt"
)

//fun formLoader()

//from form input, handlers the user answers to render the corresponding linked content
func FormHandler(writer http.ResponseWriter, request *http.Request) {
    _, err := template.ParseFiles("/html/home.html")

    if err != nil {
      fmt.Println(err)
      return
    }

    switch request.Method {
    case "GET":
         http.ServeFile(writer, request, "/html/home.html")
    case "POST":
        // Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
        err := request.ParseForm();
        if err != nil {
            fmt.Println(err)
            return
        }
        /*
        for key, value := range request.Form {

        }
        */

        /*Do stuff with the post data - that is, the processed post data which should be split into
          - question text: the question being asked / id representing the question - these should theoretically be the column ids
          - question answer: the id / content that the user answered with [for radio, value; for checkbox, listof ID, etc]

          - To get keys atm:
          for key, value := range request.Form {}
        */



        // Process the information from the questions into a unique qid-answer pair
        // This can be done by essentially keeping track of all possible answers for each question
        // This can be a mutable data structure / a database table, but NOT sure which would work better
        // Render applicable output page data based on form input based on struct

        // assets.setOutput()



    default:
        return
    }
}

func Handler(w http.ResponseWriter, req *http.Request){
  t, err := template.New("home.html").ParseFiles("html/home.html")
  if err != nil {
    fmt.Println(err)
  }
  t.Execute(w, nil)
}
