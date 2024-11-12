package models

type Livro struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
	Ano    int    `json:"ano"`
}
