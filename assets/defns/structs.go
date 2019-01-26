package defns

//Response Page Structs
type Link struct{
  URL string
  Description string
  type string
}

type ResponsePage struct {
	Links []Link
}

//View Page Structs
type Answer struct{
  AID integer
  QuestionID integer
  Name string
  Text string
}

type Question struct{
  QID integer
  OrderID integer
  Type string
  Text string
  Answers []Answer
}

type ViewPage struct {
  //Title string
  //Description string
	Questions []Question
}
