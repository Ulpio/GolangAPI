package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Ulpio/gin-api/models"
	"github.com/gin-gonic/gin"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func GetLivros(c *gin.Context) {
	rows, err := db.Query("SELECT id,titulo,autor,ano FROM livros")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error1": err.Error()})
		return
	}
	defer rows.Close()

	var livros []models.Livro
	for rows.Next() {
		var livro models.Livro
		if err := rows.Scan(&livro.ID, &livro.Titulo, &livro.Autor, &livro.Ano); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
			return
		}
		livros = append(livros, livro)
	}
	c.JSON(http.StatusOK, livros)
}

func GetLivroPorID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ID inválido"})
		return
	}
	var livro models.Livro
	row := db.QueryRow("SELECT id,titulo,autor,ano FROM livros WHERE id = $1", id)
	if err := row.Scan(&livro.ID, &livro.Titulo, &livro.Autor, &livro.Ano); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, livro)
}

func CreateLivro(c *gin.Context) {
	var novoLivro models.Livro
	if err := c.ShouldBindJSON(&novoLivro); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO livros(titulo,autor,ano) VALUES ($1,$2,$3) RETURNING id"
	if err := db.QueryRow(query, novoLivro.Titulo, novoLivro.Autor, novoLivro.Ano).Scan(&novoLivro.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, novoLivro)
}

func UpdateLivro(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ID inválido"})
		return
	}

	var livroAtualizado models.Livro
	if err := c.ShouldBindJSON(&livroAtualizado); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	query := "UPDATE livros SET titulo = $1, autor = $2, ano = $3 WHERE id = $4"
	res, err := db.Exec(query, livroAtualizado.Titulo, livroAtualizado.Autor, livroAtualizado.Ano, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Livro não encontrado"})
		return
	}

	livroAtualizado.ID = id
	c.JSON(http.StatusOK, livroAtualizado)
}

func DeleteLivro(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Id inválido"})
		return
	}

	query := "DELETE FROM livros WHERE id = $1"
	res, err := db.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Livro não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Livro deletado com sucesso"})
}
