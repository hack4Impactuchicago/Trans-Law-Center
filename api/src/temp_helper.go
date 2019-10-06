package src

import(
  "database/sql"
  "fmt"
  "log"
)

func LoadPresetDBContent(db_path string)(error){

  db, err := sql.Open("sqlite3", db_path)
  if err != nil {
    log.Println(err)
    return err
  }

  var insert_statement_1 string
  insert_statement_1 = `
      INSERT INTO Questions (Id,DispOrder,Type,Text)
      VALUES ('q1','2','radio','Question 1');
    `
  _, err = db.Exec(insert_statement_1)
  if err != nil {
      return err
  }

  var insert_statement_2 string
  insert_statement_2 = `
      INSERT INTO Questions (Id,DispOrder,Type,Text)
      VALUES ('q2','1','radio','Question 2');
    `
  _, err = db.Exec(insert_statement_2)
  if err != nil {
      return err
  }

  var insert_statement_3 string
  insert_statement_3 = `
      INSERT INTO Questions (Id,DispOrder,Type,Text)
      VALUES ('q3','3','radio','Question 3');
    `
  _, err = db.Exec(insert_statement_3)
  if err != nil {
      return err
  }

  var insert_statement_4 string
  insert_statement_4 = `
      INSERT INTO Answers (Id,QuestionId,Name,Text)
      VALUES ('q1ans1','q1','ans1','Answer 1 to Question 1');
    `
  _, err = db.Exec(insert_statement_4)
  if err != nil {
      return err
  }

  var insert_statement_5 string
  insert_statement_5 = `
      INSERT INTO Answers (Id,QuestionId,Name,Text)
      VALUES ('q1ans2','q1','ans2','Answer 2 to Question 1');
  `
  _, err = db.Exec(insert_statement_5)
  if err != nil {
      return err
  }

  var insert_statement_6 string
  insert_statement_6 = `
      INSERT INTO Answers (Id,QuestionId,Name,Text)
      VALUES ('q1ans3','q1','ans3','Answer 3 to Question 1');
  `
  _, err = db.Exec(insert_statement_6)
  if err != nil {
      return err
  }

  var insert_statement_7 string
  insert_statement_7 = `
      INSERT INTO Answers (Id,QuestionId,Name,Text)
      VALUES ('q2ans1','q2','ans1','Answer 1 to Question 2');
    `
  _, err = db.Exec(insert_statement_7)
  if err != nil {
      return err
  }

  var insert_statement_8 string
  insert_statement_8 = `
      INSERT INTO Answers (Id,QuestionId,Name,Text)
      VALUES ('q2ans2','q2','ans2','Answer 2 to Question 2');
  `
  _, err = db.Exec(insert_statement_8)
  if err != nil {
      return err
  }

  var insert_statement_9 string
  insert_statement_9 = `
      INSERT INTO Answers (Id,QuestionId,Name,Text)
      VALUES ('q3ans1','q3','ans1','Answer 1 to Question 3');
  `
  _, err = db.Exec(insert_statement_9)
  if err != nil {
      return err
  }

  var insert_statement_ten string
  key := hash_function("q2ans1q1ans1q3ans1")
  insert_statement_ten = `
      INSERT INTO Links (Id,Url,Description,Type)
      VALUES (?,'https://stackoverflow.com/questions/5952718/how-to-easily-test-posts-when-making-a-website',
        'The outputted link from combination:','online'
      );
    `
  tx, err := db.Begin()
  if err != nil {
    db.Close()
    return err
  }

  stmt, err := tx.Prepare(insert_statement_ten)
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
