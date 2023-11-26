package models

type TestData struct {
	Data Data
}

type Data struct {
	Grades    int        `json:"grades"`
	Questions []Question `json:"data"`
}

type Question struct {
	Number  int      `json:"number"`
	Query   string   `json:"query"`
	Options []string `json:"options"`
	Answer  string   `json:"answer"`
}

type Result struct {
	Number       int
	Query        string
	Answer       string
	Res          bool
	Lenght       int
	AnswerLenght int
}
