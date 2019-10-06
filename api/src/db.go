package src

import(
    "database/sql"
    "fmt"
    "log"
)

func clearDB(db *sql.DB) (error) {
    rows, _ := db.Query("SELECT name FROM sqlite_master WHERE type='table'")

    var tables []string

    for rows.Next() {
        var tableName string
        err := rows.Scan(&tableName)
        tables = append(tables, tableName)
        if err != nil {
            log.Fatal(err)
            log.Println(err)
            return err
        }
    }

    fmt.Printf("Clearing DB's: got %d tables\n", len(tables))

    for _, table := range tables {
        dropQuery := "DROP TABLE IF EXISTS " + table
        _, err := db.Exec(dropQuery)
        if err != nil {
            log.Println(err)
            return err
        } else {
            fmt.Println("Removed table \"" + table + "\" in database")
        }
    }
    return nil
}

func SetupLoginDB(db_path string) (error) {
    db, err := sql.Open("sqlite3", db_path)
    if err != nil {
      log.Println(err)
      return err
    }

    err = clearDB(db)
    if err != nil {
        db.Close()
        return err
    }

    err = CreateTable(db, "Users")
    if err != nil {
        db.Close()
        return err
    }
    defer db.Close()
    return nil
}

func SetupFormDB(db_path string) (error) {

    db, err := sql.Open("sqlite3", db_path)
    if err != nil {
      log.Println(err)
      return err
    }

    // clear up db
    clearDB(db)

    CreateTable(db, "Questions")
    CreateTable(db, "Answers")
    CreateTable(db, "Links")

    db.Close()
    return nil
}

func load_question_rows() (*sql.Rows, error){

  db, err := sql.Open("sqlite3", "formdb.db")
  if err != nil {
    return nil, err
  }

  rowsQ, err := db.Query(`SELECT * FROM Questions`)
  if err != nil {
    db.Close()
    return nil, err
  }

  db.Close()
  return rowsQ, nil
}

func load_answer_to_qid(qid string)(*sql.Rows, error){

  dbA, err := sql.Open("sqlite3", "formdb.db")
  if err != nil {
    return nil, err
  }

  rowsA, err := dbA.Query(`SELECT * from Answers where QuestionId=?`,qid)
  if err != nil {
    dbA.Close()
    return nil, err
  }

  dbA.Close()
  return rowsA, nil

}

func load_link_with_hash(hash string)(*sql.Rows, error){

  db, err := sql.Open("sqlite3", "formdb.db")
  if err != nil {
    return nil, err
  }

  rows, err := db.Query(`SELECT * from Links where Id=?`,hash)
  if err != nil {
    db.Close()
    return nil, err
  }

  return rows, nil

}
