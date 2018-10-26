package main

import(
  "golang.org/x/crypto/bcrypt"
  _ "github.com/mattn/go-sqlite3"
  "database/sql"
  "log"
  )

func login(username string, password string) (int, error){
  return 1, nil;
}

func createUser(username string, password string, adminLevel int) (int, error){
  saltAndHashed, erro := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
  strSalt := string(saltAndHashed)
  if erro != nil {
    log.Println(erro)
  }
  db, err := sql.Open("sqlite3", "database.db")
  if err != nil {
    return 0, err
  }
  tx, err := db.Begin()
  if err != nil {
    return 0, err
  }
  stmt, err := tx.Prepare("INSERT INTO users values(?, ?, ?)")
  if err != nil {
    return 0, err
  }
  _, err = stmt.Exec(username, strSalt, adminLevel)
  if err != nil {
    return 0, err
  }
  tx.Commit()
  db.Close()
  return 1, nil
}

func changePassword(username string, newPassword string, oldPassword string) (int, error){
  return 1, nil;
}
