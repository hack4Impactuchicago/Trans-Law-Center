package src

import(
  "golang.org/x/crypto/bcrypt"
  _ "github.com/mattn/go-sqlite3"
  "database/sql"
  "log"
  "fmt"
  )

func Login(username string, password string) (int, error){
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

func CreateUser(username string, password string, adminLevel int) (int, error){
  saltAndHashed, erro := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
  strSalt := string(saltAndHashed)
  if erro != nil {
    log.Println(erro)
  }
  db, err := sql.Open("sqlite3", "database.db")
  if err != nil {
    return 0, err
  }

  // Check for existing username
  row := db.QueryRow("SELECT * FROM Users WHERE Username=?", username)

  var u, p string
  var al int
  scanerr := row.Scan(&u, &p, &al)

  if scanerr == nil {
    fmt.Println("Username Already Exists")
    db.Close()
    return 0, err
  }
  // finish checking for username

  tx, err := db.Begin()
  if err != nil {
    db.Close()
    return 0, err
  }

  stmt, err := tx.Prepare("INSERT INTO Users values(?, ?, ?)")
  if err != nil {
    db.Close()
    return 0, err
  }
  _, err = stmt.Exec(username, strSalt, adminLevel)
  if err != nil {
    db.Close()
    return 0, err
  }
  tx.Commit()
  db.Close()
  return 1, nil
}

func ChangePassword(username string, newPassword string, oldPassword string) (int, error){
  db, err := sql.Open("sqlite3", "database.db")
    if(err != nil){
      fmt.Println("error after open")
      return 0,nil
    }
  sqlStatement := `SELECT Password FROM Users WHERE username=$1;`
  row := db.QueryRow(sqlStatement,username)
  if(row == nil){
    fmt.Println("error after query")
    db.Close()
    return 0,nil
  }
  var password string
  row.Scan(&password)

  hashedPwd := []byte(password)
  bytePlainPwd := []byte(oldPassword)
  err = bcrypt.CompareHashAndPassword(hashedPwd, bytePlainPwd)

  if(err != nil){
    fmt.Println("passwords not equal")
    db.Close()
    return 0,nil
  } else {
    statement, err := db.Prepare("UPDATE Users SET Password=? WHERE Username=?")
    if(err != nil){
      fmt.Println("error after prepare")
      db.Close()
      return 0,nil
    }
    saltAndHashed, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.MinCost)
    strSalt := string(saltAndHashed)
    statement.Exec(strSalt,username)
  }
  db.Close()
  return 1, nil;
}
