package src

import(
  "crypto/sha1"
  "encoding/hex"
  "database/sql"
  //"fmt"
  "net/http"
)

func hash_function(s string) string{
  h := sha1.New()
  h.Write([]byte(s))
  result := hex.EncodeToString(h.Sum(nil))
  return result
}

func insert_map(disporder_map map[int]string, order int, key string)(map[int]string){
  if _, exists := disporder_map[order]; exists {
    return disporder_map
  }else{
    disporder_map[order] = key
    return disporder_map
  }
}

func generate_unhashed_id(r *http.Request)(*string, error){

  if err := r.ParseForm(); err != nil {return nil, err}
  disporder_map := map[int]string{}

  db, err := sql.Open("sqlite3", "formdb.db")
  if err != nil {return nil, err}

  var unhashed_key string
  var disporder int

  for key, _ := range r.Form {  // range over map

    order, err := db.Query(`SELECT DispOrder from Questions where Id=?`,key)
    if err != nil {return nil, err}

    for order.Next(){
      err := order.Scan(&disporder)
      if err != nil {
        db.Close()
        return nil, err
      }else{
        disporder_map = insert_map(disporder_map, disporder, key)
      }
    }
  }

  i := 0
  var values []string

  for i <= len(disporder_map){

    key := disporder_map[i]
    values = r.Form[key]

    for _, value := range values{    // range over []string

      rows, errA := db.Query(`SELECT * from Answers where Id=?`,value)

      if errA != nil {
        db.Close()
        return nil, errA
      }

      var aid, qid, name, textQ string

      for rows.Next(){
        err := rows.Scan(&aid, &qid, &name, &textQ)
        //fmt.Printf("aid: %s ",aid)
        if err != nil {
          db.Close()
          return nil, err
        }else{
          unhashed_key = unhashed_key + aid
          //fmt.Printf("unhashed-key: %s",unhashed_key)
        }
      }
    }

    i += 1
  }

  db.Close()
  return &unhashed_key, nil

}
