package assets

import(
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "fmt"
    "log"
    //"Trans-Law-Center/assets/defns"
)

func CreateTable(db *sql.DB, table_type string) (error) {
    // create table if not exists
    var sql_table string

    // find a way to specify a unique key for each table
    switch table_type {
    case "Questions":
        sql_table = `
            CREATE TABLE IF NOT EXISTS Questions(
                Id INT,
                DispOrder INT,
                Type TEXT,
                Text TEXT
            );
        `
    case "Answers":
        sql_table = `
            CREATE TABLE IF NOT EXISTS Answers(
                Id INT,
                QuestionId INT,
                Name TEXT,
                Text TEXT
            );
        `
    case "Links":
        sql_table = `
            CREATE TABLE IF NOT EXISTS Links(
                Id STRING,
                Url TEXT,
                Description TEXT,
                Type TEXT
            )
        `
    case "Users":
        sql_table = `
            CREATE TABLE IF NOT EXISTS Users(
                Username TEXT,
                Password TEXT,
                AdminLevel INT
            )
        `
    }

    _, err := db.Exec(sql_table)
    if err != nil {
        log.Println(err)
        return err
    } else {
        fmt.Println("Created table of type \"" + table_type + "\" in database")
        return nil
    }
}

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
    db.Close()
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

func load_answer_to_qid(qid int)(*sql.Rows, error){

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

func LoadPresetDBContent(db_path string)(error){

  db, err := sql.Open("sqlite3", db_path)
  if err != nil {
    log.Println(err)
    return err
  }

  var insert_statement_1 string
  insert_statement_1 = `
      INSERT INTO Questions (
        Id,
        DispOrder,
        Type,
        Text
      )
      VALUES (
        '1',
        '1',
        'radio',
        'Question 1'
      );
    `

  var insert_statement_2 string
  insert_statement_2 = `
      INSERT INTO Questions (
        Id,
        DispOrder,
        Type,
        Text
      )
      VALUES (
        '2',
        '2',
        'radio',
        'Question 2'
      );
    `

  var insert_statement_3 string
  insert_statement_3 = `
      INSERT INTO Answers (
        Id,
        QuestionId,
        Name,
        Text
      )
      VALUES (
        '1',
        '1',
        'ans1',
        'Answer 1 to Question 1'
      );
    `

  var insert_statement_4 string
  insert_statement_4 = `
      INSERT INTO Answers (
        Id,
        QuestionId,
        Name,
        Text
      )
      VALUES (
        '2',
        '2',
        'ans1',
        'Answer 1 to Question 2'
      );
    `
  var insert_statement_5 string
  key := hash_function("11")
  insert_statement_5 = `
      INSERT INTO Links (
        Id,
        Url,
        Description,
        Type
      )
      VALUES (
        ?,
        'https://stackoverflow.com/questions/5952718/how-to-easily-test-posts-when-making-a-website',
        'fuck this',
        'online'
      );
    `

  _, err = db.Exec(insert_statement_1)
  if err != nil {
      return err
  }

  _, err = db.Exec(insert_statement_2)
  if err != nil {
      return err
  }

  _, err = db.Exec(insert_statement_3)
  if err != nil {
      return err
  }

  _, err = db.Exec(insert_statement_4)
  if err != nil {
      return err
  }

  tx, err := db.Begin()
  if err != nil {
    db.Close()
    return err
  }

  stmt, err := tx.Prepare(insert_statement_5)
  if err != nil {
    db.Close()
    return err
  }
  _, err = stmt.Exec(key)
  if err != nil {
    db.Close()
    return err
  }
  tx.Commit()
  db.Close()

  fmt.Println("Initialized TEST Database Tables.")
  return nil

}
