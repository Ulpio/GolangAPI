package models

import "gorm.io/gorm"

type Livro struct {
	gorm.Model
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
	Ano    int    `json:"ano"`
}
