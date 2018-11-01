package main
// import "errors"
import "strings"
import "fmt"
import(
  // "golang.org/x/crypto/bcrypt"
  _ "github.com/mattn/go-sqlite3"
  "database/sql"
  )

// func login(username string, password string) (int, error){
//   return 1, nil;
// }
//
// func createUser(username string, password string, adminLevel string) (int, error){
//   return 1, nil;
// }

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

func join(strs ...string) string {//concatenates strings
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}
