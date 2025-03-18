package routes

import (
	"api/internal/handlers"
	"github.com/gin-gonic/gin"
)

// InitRoutes configura as rotas do servidor.
func InitRoutes(c *gin.Engine) {
	// Associando a rota "/" (página inicial) à função HomeController do pacote controllers.
	c.POST("/", handlers.PostUser)
}
