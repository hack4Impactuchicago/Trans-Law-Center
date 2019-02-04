package assets

import(
  "crypto/sha1"
  "encoding/hex"
  "database/sql"
  "fmt"
  "strconv"
  "Trans-Law-Center/assets/defns"
  "errors"
  "net/http"
)

func hash_function(s string) string{
  h := sha1.New()
  h.Write([]byte(s))
  result := hex.EncodeToString(h.Sum(nil))
  return result
}

func length(list *defns.Order) (int){
  i := 0
  for list != nil {
    i += 1
  }
  return i
}

func insert_list(list *defns.Order, ind int, key int) (*defns.Order, error) {

    new_ord := defns.Order{Content: ind, Key: key, Next: nil}
    len_of := length(list)

    if len_of > ind { return nil, errors.New("invalid Length") }
    if list == nil{ return &new_ord, nil }

    head := list
    var prev *defns.Order
    prev = nil

    for list != nil {
      if list.Content < ind {
        if prev == nil{
          new_ord.Next = list
          head = list
        } else {
          prev.Next = &new_ord
          new_ord.Next = list
        }
        break
      }
      prev = list
      list = list.Next
    }

    return head, nil
}

func generate_unhashed_id(r *http.Request)(*string, error){

  if err := r.ParseForm(); err != nil {return nil, err}

  db, err := sql.Open("sqlite3", "formdb.db")
  if err != nil {return nil, err}

  var order_list *defns.Order
  order_list = nil

  var unhashed_key string
  for key, _ := range r.Form {  // range over map

    order, err := db.Query(`SELECT DispOrder from Questions where Id=?`,key)
    if err != nil {return nil, err}

    for order.Next(){
      err := order.Scan(&disporder)
      if err != nil {
        db.Close()
        return nil, err
      }else{
        order_list, err = insert_list(order_list, key, disporder)
        if err != nil {return nil, err}
      }
    }
  }

  for ordered_key := range order_list{
    values := r.Form[ordered_key.Key]
    for value := range values{    // range over []string

      row, errA := db.Query(`SELECT * from Answers where QuestionId=? AND Id=?`,
        key,value)
      if errA != nil {
        db.Close()
        return nil, errA
      }

      var aid, qid int
      var name, textQ string

      for row.Next(){
        err := row.Scan(&aid, &qid, &name, &textQ)
        if err != nil {
          db.Close()
          return nil, err
        }else{
          unhashed_key = unhashed_key + strconv.Itoa(aid)
        }
      }
    }
  }

  db.Close()
  return &unhashed_key, nil

}
