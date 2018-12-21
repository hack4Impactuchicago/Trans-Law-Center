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

        /*Do stuff with the post data - that is, the processed post data which should be split into
          - question text: the question being asked / id representing the question - these should theoretically be the column ids
          - question answer: the id / content that the user answered with [for radio, value; for checkbox, listof ID, etc]

          - To get keys atm:
          for key, value := range request.Form {}
        */

        /*

        DB Table:

        _____________________________________
        |    Q1     |    Q2     |     Q3    |
        _____________________________________
        |     A1a   |     A2a   |     A3a   |
        |     A1b   |     A2b   |     A3b   |
        |     A1c   |     A2c   |     A3c   |
        |     A1d   |     A2d   |     A3d   |
              ...         ...         ...

        */


        // Render applicable output page data based on form input

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
