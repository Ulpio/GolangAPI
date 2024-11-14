package routers

import (
	"github.com/Ulpio/gin-api/controller"
	"github.com/gin-gonic/gin"
)

func DefineRotas() {
	roteador := gin.Default()
	roteador.GET("/livros", controller.GetLivros)
	roteador.GET("/livros/:id", controller.GetLivroPorID)
	roteador.POST("/livros", controller.CreateLivro)
	roteador.PUT("/livros/:id", controller.UpdateLivro)
	roteador.DELETE("/livros/:id", controller.DeleteLivro)
	roteador.Run(":8000")
}
