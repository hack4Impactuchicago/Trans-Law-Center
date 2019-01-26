package defns

//Response Page Structs
type Link struct{
  URL string
  Description string
  Type string
}

type ResponsePage struct {
  //Title string
  //Description string
	Links []Link
}

//View Page Structs
type Answer struct{
  AID int
  QuestionID int
  Name string
  Text string
}

type Question struct{
  QID int
  OrderID int
  Type string
  Text string
  Answers []Answer
}

type ViewPage struct {
  //Title string
  //Description string
	Questions []Question
}
