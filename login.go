package main

import(
  "golang.org/x/crypto/bcrypt"
  _ "github.com/mattn/go-sqlite3"
  "database/sql"
  "log"
  )

func login(username string, password string) (int, error){
  db, err := sql.Open("sqlite3", "database.db")
  if err != nil {
    log.Println(err)
    return 0, err
  }

  row := db.QueryRow("SELECT Password FROM Users WHERE Username=?", username)

  var hashedPwd string
  err = row.Scan(&hashedPwd)

  switch {
    case err == sql.ErrNoRows:
      // username does not exist
      log.Println("Username does not exist")
      log.Println(err)
      db.Close()
      return 0, err
    case err != nil:
      log.Println(err)
      db.Close()
      return 0, err
  }

  byteHash := []byte(hashedPwd)
  bytePlainPwd := []byte(password)

  err = bcrypt.CompareHashAndPassword(byteHash, bytePlainPwd)
  if err != nil {
    log.Println("Wrong password")
    log.Println(err)
    // wrong password
    db.Close()
    return 0, nil
  }
  db.Close()
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
  db, err := sql.Open("sqlite3", "./database.db")
    if(err != nil){
      fmt.Println("error after open")
      return 0,nil
    }
  sqlStatement := `SELECT Password FROM users WHERE username=$1;`
  row := db.QueryRow(sqlStatement,username)
  if(row == nil){
    fmt.Println("error after query")
    db.Close()
    return 0,nil
  }
  var password string
  row.Scan(&password)
  if(password != oldPassword){
    fmt.Println("passwords not equal")
    db.Close()
    return 0,nil
  } else {
    statement, err := db.Prepare("UPDATE users SET Password=? WHERE Username=?")
    if(err != nil){
      fmt.Println("error after prepare")
      db.Close()
      return 0,nil
    }
    statement.Exec(newPassword,username)
  }
  db.Close()
  return 1, nil;
}
