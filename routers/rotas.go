package routers

import (
	"github.com/Ulpio/gin-api/controller"
	"github.com/Ulpio/gin-api/database"
	"github.com/gin-gonic/gin"
)

func DefineRotas() {
	db := database.InitDB()
	defer db.Close()
	controller.SetDB(db)
	roteador := gin.Default()
	roteador.GET("/livros", controller.GetLivros)
	roteador.GET("/livros/:id", controller.GetLivroPorID)
	roteador.POST("/livros", controller.CreateLivro)
	roteador.PUT("/livros/:id", controller.UpdateLivro)
	roteador.DELETE("/livros/:id", controller.DeleteLivro)
	roteador.Run(":8000")
}
