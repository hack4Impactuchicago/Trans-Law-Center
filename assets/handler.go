package assets

import(
  "net/http"
  "html/template"
  "strconv"
  "database/sql"
  "fmt"
  "Trans-Law-Center/assets/defns"
)

//function for loading the content for the form from the DB
func loadViewPage()(*defns.ViewPage, error){
  // Loading all rows of Questions from DB
  page := defns.ViewPage{Questions: nil}

  db, err := sql.Open("sqlite3", "formdb.db")
  if err != nil {
    return &page, err
  }

  rowsQ, err := db.Query(`SELECT * FROM Questions`)
  if err != nil {
    db.Close()
    return &page, err
  }

  var Questions []defns.Question
  for rowsQ.Next() { //for each row within the datatable

    var qid, orderID int
    var typeQ, textQ string
    if err = rowsQ.Scan(&qid, &orderID, &typeQ, &textQ); err != nil {
      return &page, err
    } else {
      //fmt.Printf("%d,%d,%s,%s",qid,orderID,typeQ,textQ)
      var AnsList []defns.Answer
      dbA, err := sql.Open("sqlite3", "formdb.db")
      if err != nil {
        return &page, err
      }

      rowsA, errA := dbA.Query(`SELECT * from Answers where QuestionId=?`,qid)
      if errA != nil {
        dbA.Close()
        return &page, errA
      }

      for rowsA.Next(){

        var aid, Qid int
        var name, textQ string

        if err = rowsA.Scan(&aid, &Qid, &name, &textQ); err != nil {
        	return &page, err
        }else{
          AnsList = append(AnsList,
            defns.Answer{AID: aid, QuestionID: qid, Name: name, Text: textQ})
        }
      }

      dbA.Close()

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

  db.Close()
  //return the constructed page
  page = defns.ViewPage{Questions: Questions}
  return &page, nil
}

func loadResponsePage(r *http.Request)(*defns.ResponsePage, error){

  temp_page := defns.ResponsePage{Links: nil}

  if r == nil {
    fmt.Println("Not able to parse Form")
    return &temp_page, nil
  }

  if err := r.ParseForm(); err != nil {
    fmt.Println("Not able to parse Form")
    return &temp_page, err
  }

  db, err := sql.Open("sqlite3", "formdb.db")
  if err != nil {
    return &temp_page, err
  }

  var unhashed_key string

  for key, values := range r.Form {  // range over map

    fmt.Println(key)
    fmt.Println(values)

    for _, value := range values {    // range over []string

      fmt.Println(value)

      row, errA := db.Query(
        `SELECT * from Answers where QuestionId=? AND Id=?`,
        key,
        value)
      if errA != nil {
        db.Close()
        return &temp_page, errA
      }

      var aid, qid int
      var name, textQ string

      for row.Next(){
        err := row.Scan(&aid, &qid, &name, &textQ)
        if err != nil {
          return &temp_page, err
        }else{
          fmt.Printf("********* %d", aid)
          unhashed_key = unhashed_key + strconv.Itoa(aid)
        }
      }

    }
  }


  var hashed_key string
  hashed_key = hash_function(unhashed_key)
  db.Close()

  fmt.Println(unhashed_key)
  fmt.Println(hashed_key)

  //da39a3ee5e6b4b0d3255bfef95601890afd80709

  rows, err := db.Query(
    `SELECT * from Links where Id=?`,hashed_key)

  var LinksList []defns.Link
  for rows.Next(){
    var id int
    var url, description, Type string

    err1 := rows.Scan(&id, &url, &description, &Type);
    if err1 != nil {
      return &temp_page, err
    } else {
      LinksList = append(LinksList,
        defns.Link{URL:url, Description: description, Type: Type})
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
    //fmt.Printf(errload.Error())
    //http.Error(w, errload.Error(), http.StatusInternalServerError)
  }

  t, _ := template.ParseFiles("html/links.html")
  if err := t.Execute(w, p); err != nil{
    //fmt.Printf(err.Error())
    //http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
