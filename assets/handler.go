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
  rowsA, errA := AllRows("formdb.db", "Answers")
  if err != nil {
    return nil, err
  }
  defer rows.Close()
  var Questions []Question

  for rowsQ.Next() { //for each row within the datatable
    if err := rowsQ.Scan(&qid, &orderID, &typeQ, &textQ); err != nil {
    	return nil, err
    }else{

      rowsA, errA := db.Query(`SELECT * from Answers where QuestionId=?`,qid)
      if errA != nil {
        log.Println(errA)
        db.Close()
        return nil, errA
      } else {
        db.Close()
      }

      var AnsList []Answer
      rowsA, err := AllRows("formdb.db", "Answers")
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
  if err := r.ParseForm(); err != nil {
    return nil, err
  }
  for key, values := range r.PostForm {
    
  }
}

func ViewHandler(w http.ResponseWriter, r *http.Request){
  p, errload := loadViewPage()
  t, _ := template.ParseFiles("/html/home.html")
  if err := t.Execute(w, p); err != nil{
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func ResultsHandler(w http.ResponseWriter, r *http.Request) {
  p, _ := loadResponsePage(r)

}
