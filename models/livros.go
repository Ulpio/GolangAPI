package models

type Livro struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
	Ano    int    `json:"ano"`
}

var Livros = []Livro{
	{1, "Senhor dos Aneis", "Tolkien", 1954},
	{2, "Codigo da Vinci", "Dan Brown", 2003},
	{3, "Dom Quixote", "Miguel de Cervantes", 1605},
}
