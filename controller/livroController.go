package controller

import (
	"net/http"
	"strconv"

	"github.com/Ulpio/gin-api/models"
	"github.com/gin-gonic/gin"
)

func GetLivros(c *gin.Context) {
	c.JSON(http.StatusOK, models.Livros)
}

func GetLivroPorID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "ID inválido"})
		return
	}
	for _, livro := range models.Livros {
		if livro.ID == id {
			c.JSON(http.StatusOK, livro)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"msg": "Livro não encontrado!"})
}

func CreateLivro(c *gin.Context) {
	var novoLivro models.Livro
	if err := c.ShouldBindJSON(&novoLivro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Dados inválidos"})
		return
	}
	novoLivro.ID = len(models.Livros) + 1
	models.Livros = append(models.Livros, novoLivro)
	c.JSON(http.StatusCreated, novoLivro)
}

func UpdateLivro(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Id inválido"})
		return
	}
	var livroAtualizado models.Livro
	if err := c.ShouldBindJSON(&livroAtualizado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Dados inválidos"})
		return
	}
	for i, livro := range models.Livros {
		if livro.ID == id {
			livroAtualizado.ID = id
			models.Livros[i] = livroAtualizado
			c.JSON(http.StatusOK, livroAtualizado)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"msg": "Livro não encontrado"})
}

func DeleteLivro(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "ID inválido"})
		return
	}
	for i, livro := range models.Livros {
		if livro.ID == id {
			models.Livros = append(models.Livros[:i], models.Livros[i+1:]...)
			c.JSON(http.StatusNoContent, nil)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"msg": "Livro não encontrado"})
}

func GetLivros(c *gin.Context) {
	rows, err := db.Query("SELECT id, titulo, autor, ano FROM livros")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var livros []models.Livro
	for rows.Next() {
		var livro models.Livro
		if err := rows.Scan(&livro.ID, &livro.Titulo, &livro.Autor, &livro.Ano); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		livros = append(livros, livro)
	}
	c.JSON(http.StatusOK, livros)
}
