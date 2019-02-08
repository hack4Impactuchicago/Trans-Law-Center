package assets

import(
  "fmt"
  "net/http"
  "html/template"
  "Trans-Law-Center/assets/defns"
)

//function for loading the content for the form from the DB
func loadViewPage()(*defns.ViewPage, error){
  // Loading all rows of Questions from DB

  page := defns.ViewPage{Questions: nil}

  rowsQ, err := load_question_rows()
  if err != nil{return &page, err}

  var Questions []defns.Question
  for rowsQ.Next() { //for each row within the datatable

    var orderID int
    var qid, typeQ, textQ string

    if err = rowsQ.Scan(&qid, &orderID, &typeQ, &textQ); err != nil {
      return &page, err
    } else {

      rowsA, err := load_answer_to_qid(qid)
      if err != nil{return &page, err}

      var AnsList []defns.Answer

      for rowsA.Next(){

        var aid, Qid, name, textQ string

        if err = rowsA.Scan(&aid, &Qid, &name, &textQ); err != nil {
        	return &page, err
        }else{
          AnsList = append(AnsList,
            defns.Answer{AID: aid,
              QuestionID: qid,
              Name: name,
              Text: textQ})
        }
      }

      //TODO: Look into how slices are stored in Memory
      Questions = append(Questions,
        defns.Question{
          QID: qid,
          OrderID: orderID,
          Type: typeQ,
          Text: textQ,
          Answers: AnsList})
    }
  }

  //return the constructed page
  page = defns.ViewPage{Questions: Questions}
  return &page, nil
}

func loadResponsePage(r *http.Request)(*defns.ResponsePage, error){

  temp_page := defns.ResponsePage{Links: nil}
  unhashed_key, err := generate_unhashed_id(r)
  if err != nil{return &temp_page, err}

  var hashed_key string
  hashed_key = hash_function(*unhashed_key)

  fmt.Println(hashed_key)

  rows, err := load_link_with_hash(hashed_key)
  if err != nil{return &temp_page, err}

  var LinksList []defns.Link
  for rows.Next(){

    var id, url, description, Type string

    err1 := rows.Scan(&id, &url, &description, &Type);

    if err1 != nil {
      return &temp_page, err
    } else {
      LinksList = append(LinksList,
        defns.Link{URL:url,
          Description: description,
          Type: Type})
    }
  }

  page := defns.ResponsePage{Links: LinksList}
  return &page, nil
}

func ViewHandler(w http.ResponseWriter, r *http.Request){
  p, errload := loadViewPage()
  if errload != nil{
    http.Error(w, errload.Error(), http.StatusInternalServerError)
  }

  t, _ := template.ParseFiles("html/home.html")
  if err := t.Execute(w, *p); err != nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func ResultsHandler(w http.ResponseWriter, r *http.Request) {
  p, errload := loadResponsePage(r);
  if errload != nil{
    http.Error(w, errload.Error(), http.StatusInternalServerError)
  }

  t, _ := template.ParseFiles("html/links.html")
  if err := t.Execute(w, p); err != nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
