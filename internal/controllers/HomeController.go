package controllers

import "github.com/gin-gonic/gin"

// HomeController responsável por manipular requisições para a página initial (home).
func HomeController(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "API conectada com sucesso",
	})
}
