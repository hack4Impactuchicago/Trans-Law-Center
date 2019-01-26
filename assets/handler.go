package assets

import(
  "net/http"
  "html/template"
  //"fmt"
  "database/sql"
  "log"
  "Trans-Law-Center/assets/defns"
)

//function for loading the content for the form from the DB
func loadViewPage()(*defns.ViewPage, error){
  rowsQ, errQ := AllRows("formdb.db", "Questions") //Load all rows within DB

  if errQ != nil {
    return nil, errQ
  }
  defer rowsQ.Close()
  var Questions []defns.Question

  for rowsQ.Next() { //for each row within the datatable

    var qid, orderID int
    var typeQ, textQ string

    if err := rowsQ.Scan(&qid, &orderID, &typeQ, &textQ); err != nil {
    	return nil, err
    } else {

      var AnsList []defns.Answer

      db, errO := sql.Open("sqlite3", "formdb.db")
      if errO != nil {
        log.Println(errO)
        return nil, errO
      }

      rowsA, errA := db.Query(`SELECT * from Answers where QuestionId=?`,qid)
      if errA != nil {
        log.Println(errA)
        db.Close()
        return nil, errA
      }
      db.Close()

      for rowsA.Next(){

        var aid, Qid int
        var name, textQ string

        if err := rowsA.Scan(&aid, &Qid, &name, &textQ); err != nil {
        	return nil, err
        }else{
          AnsList = append(AnsList,
            defns.Answer{AID: aid, QuestionID: qid, Name: name, Text: textQ})
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
  page := defns.ViewPage{Questions: Questions}

  return *page, nil
}

func loadResponsePage(r *http.Request)(*ResponsePage, error){
  if err := r.ParseMultipartForm(); err != nil {
    return nil, err
  }

  db, err := sql.Open("sqlite3", "formdb.db")
  if err != nil {
    log.Println(err)
    return nil, err
  }

  var unhashed_key string

  for key, values := range r.MultipartForm {
    row, errA := db.Query(
      `SELECT * from Answers where QuestionId=? AND Id=?`,
      key,
      value)
    if errA != nil {
      log.Println(errA)
      db.Close()
      return nil, errA
    }

    if err := row.Scan(&aid, &qid, &name, &textQ); err != nil {
      return nil, err
    }else{
      unhashed_key += aid
    }
  }

  hashed_key = hash_function(unhashed_key)
  db.Close()

  rows, err := db.Query(
    `SELECT * from Links where Id=?`,hashed_key)

  var LinksList []Link

  for rows.Next(){
    err1 := rows.Scan(&id, &url, &description, &Type);
    if err1 != nil {
      return nil, err
    } else {
      LinksList = append(LinksList,
        defns.Link{URL:url, Description: description, Type: Type})
    }
  }
  return *defns.ResponsePage{Links: LinksList}
}

func ViewHandler(w http.ResponseWriter, r *http.Request){
  if *p, errload := loadViewPage(); errload != nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  t, _ := template.ParseFiles("/html/home.html")
  if err := t.Execute(w, *p); err != nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func ResultsHandler(w http.ResponseWriter, r *http.Request) {
  if *p, errload := loadResponsePage(r); errload != nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  t, _ := template.ParseFiles("/html/links.html")
  if err := t.Execute(w, *p); err != nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
