package src

import(
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "fmt"
    "log"
)

func CreateTable(db *sql.DB, table_type string) (error) {
    // create table if not exists
    var sql_table string

    // find a way to specify a unique key for each table
    switch table_type {
    case "Questions":
        sql_table = `
            CREATE TABLE IF NOT EXISTS Questions(
                Id TEXT,
                DispOrder INT,
                Type TEXT,
                Text TEXT
            );
        `
    case "Answers":
        sql_table = `
            CREATE TABLE IF NOT EXISTS Answers(
                Id TEXT,
                QuestionId TEXT,
                Name TEXT,
                Text TEXT
            );
        `
    case "Links":
        sql_table = `
            CREATE TABLE IF NOT EXISTS Links(
                Id TEXT,
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
