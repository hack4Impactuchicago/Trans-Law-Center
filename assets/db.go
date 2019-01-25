package assets

import(
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "fmt"
    "log"
    "Trans-Law-Center/defns"
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
                Id INT,
                Url TEXT,
                Description TEXT
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

func AllRows(db_path string, table_type string) (*Rows, error){

  db, err := sql.Open("sqlite3", db_path)
  if err != nil {
    log.Println(err)
    return nil, err
  }

  switch table_type {
  case condition:
    "Questions":
      command = `
        SELECT * FROM Questions
      `
    "Answers":
      command = `
        SELECT * FROM Answers
      `
    "Links":
      command = `
        SELECT * FROM Links
      `
    "Users":
      command = `
        SELECT * FROM Users
      `
  }

  rows, err := db.Query(command)
  if err != nil {
    log.Println(err)
    db.Close()
    return nil, err
  } else {
    db.Close()
    return rows, nil
  }

}
