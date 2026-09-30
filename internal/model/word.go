package model

type WordQuery struct {
	Paging
	Search string
}

type WordInput struct {
	Id      Id
	Word    string `json:"word" v:"required|length:1,100"`
	Explain string `json:"explain" v:"required|length:1,500"`
}
