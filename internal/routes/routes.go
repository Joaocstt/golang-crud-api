package routes

import (
	"api/internal/handlers"
	"github.com/gin-gonic/gin"
)

// InitRoutes configura as rotas do servidor.
func InitRoutes(c *gin.Engine) {
	// Associando a rota "/cadastrar-usuario" à função na Handler, para executar o fluxo da criação de usuário
	c.POST("/cadastrar-usuario", handlers.PostUser)
	c.GET("/listar-usuarios", handlers.GetUser)
}
