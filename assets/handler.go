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
  rowsQ, errQ := AllRows("formdb.db", "Questions")

  if err != nil {
    return nil, err
  }
  defer rows.Close()
  var Questions []Question

  for rowsQ.Next() { //for each row within the datatable
    if err := rowsQ.Scan(&qid, &orderID, &typeQ, &textQ); err != nil {
    	return nil, err
    }else{

      var AnsList []Answer

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
        if err := rowsA.Scan(&aid, &qid, &name, &textQ); err != nil {
        	return nil, err
        }else{
          AnsList = append(AnsList,
            Answer{AID: aid, QuestionID: qid, Name: name, Text: textQ})
        }
      }

      //TODO: Look into how slices are stored in Memory
      Questions = append(Questions,
        Question{
          QID: qid,
          OrderID: orderID,
          Type: typeQ,
          Text: textQ,
          Answers: AnsList
        })
    }
  }
  //return the constructed page
  return ViewPage{Questions: Questions}
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
    if err := rows.Scan(&id, &url, &description, &type); err != nil {
      return nil, err
    }else{

    }
  }


}

func ViewHandler(w http.ResponseWriter, r *http.Request){
  if p, errload := loadViewPage(); errload != nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  t, _ := template.ParseFiles("/html/home.html")
  if err := t.Execute(w, p); err != nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func ResultsHandler(w http.ResponseWriter, r *http.Request) {
  if p, errload := loadResponsePage(r); errload != nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  t, _ := template.ParseFiles("/html/links.html")
  if err := t.Execute(w, p); err != nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
