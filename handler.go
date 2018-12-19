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
        // r.Form contains the data from the form based on value keys
        // r.PostForm contains the data as a whole

        // Do stuff with the post data

        // Render applicable output page data based on form input

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
