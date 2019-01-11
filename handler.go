package assets

import(
  "fmt"
  "net/http"
  "html/template"

)

func formHandler(writer http.ResponseWriter, request *http.Request) {
    t, err := template.ParseFiles("/html/home.html")

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
        // request.Form contains the data from the form based on value keys
        // request.PostForm contains the data as a whole

        for key, value := range request.Form {

        }

        /*Do stuff with the post data - that is, the processed post data which should be split into
          - question text: the question being asked / id representing the question - these should theoretically be the column ids
          - question answer: the id / content that the user answered with [for radio, value; for checkbox, listof ID, etc]

          - To get keys atm:
          for key, value := range request.Form {}
        */

        /*

        DB Question Table:

        _____________________________________
        |    Q1     |    Q2     |     Q3    | ...
        _____________________________________
        |     A1a   |     A2a   |     A3a   | ...
        |     A1b   |     A2b   |     A3b   | ...
        |     A1c   |     A2c   |     A3c   | ...
        |     A1d   |     A2d   |     A3d   | ..
              ...         ...         ...

        */

        // Process the information from the questions into a unique qid-answer pair
        // This can be done by essentially keeping track of all possible answers for each question
        // This can be a mutable data structure / a database table, but NOT sure which would work better

        /*

        DB Response Table:

        _____________________________________
        |    R1ID   |    R2ID   |    R3ID   |
        _____________________________________
        |  R1struct |  R2struct |  R3struct |

        Struct form should be of:


        */


        // Render applicable output page data based on form input based on struct

        assets.setOutput()



    default:
        return
    }
}

func handler(w http.ResponseWriter, req *http.Request){
  t, err := template.New("home.html").ParseFiles("html/home.html")
  if err != nil {
    fmt.Println(err)
  }
  t.Execute(w, nil)
}
