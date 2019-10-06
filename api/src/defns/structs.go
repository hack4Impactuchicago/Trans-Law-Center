package defns

//Response Page Structs
type Link struct{
  URL string
  Description string
  Type string
}

type ResponsePage struct {
	Links []Link
}

//View Page Structs
type Answer struct{
  AID string
  QuestionID string
  Name string
  Text string
}

type Question struct{
  QID string
  OrderID int
  Type string
  Text string
  Answers []Answer
}

type ViewPage struct {
	Questions []Question
}

type AdminPage struct {
  Questions []Question
}

// User
type User struct {
	Password string
	Username string
}
