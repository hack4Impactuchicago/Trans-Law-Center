package main

import(
  //"golang.org/x/crypto/bcrypt"
  _ "github.com/mattn/go-sqlite3"
  //"database/sql"
  )

func login(username string, password string) (int, error){
  return 1, nil;
}

func createUser(username string, password string, adminLevel string) (int, error){
  return 1, nil;
}

func changePassword(username string, newPassword string, oldPassword string) (int, error){
  return 1, nil;
}
