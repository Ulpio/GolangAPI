package controller

import (
	"net/http"

	"github.com/Ulpio/gin-api/database"
	"github.com/Ulpio/gin-api/models"
	"github.com/gin-gonic/gin"
)

func GetLivros(c *gin.Context) {
	var livros []models.Livro
	database.DB.Find(&livros)
	c.JSON(http.StatusOK, livros)
}

func GetLivroPorID(c *gin.Context) {
	id := c.Param("id")
	var livro []models.Livro
	database.DB.First(&livro, id)
	c.JSON(http.StatusOK, livro)
}

func CreateLivro(c *gin.Context) {
	var novoLivro models.Livro
	if err := c.ShouldBindJSON(&novoLivro); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&novoLivro)
	c.JSON(http.StatusCreated, novoLivro)
}

func UpdateLivro(c *gin.Context) {
	id := c.Param("id")
	var livroAtualizado models.Livro
	database.DB.First(&livroAtualizado, id)
	if err := c.ShouldBindJSON(&livroAtualizado); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&livroAtualizado)
	c.JSON(http.StatusOK, livroAtualizado)
}

func DeleteLivro(c *gin.Context) {
	id := c.Param("id")
	var livro models.Livro
	database.DB.Delete(&livro, id)
	c.JSON(http.StatusOK, gin.H{"msg": "Livro deletado com sucesso"})
}
