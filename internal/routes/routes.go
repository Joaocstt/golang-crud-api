package routes

import (
	"api/internal/controllers"
	"github.com/gin-gonic/gin"
)

// InitRoutes configura as rotas do servidor.
func InitRoutes(c *gin.Engine) {
	// Associando a rota "/" (página inicial) à função HomeController do pacote controllers.
	c.GET("/", controllers.HomeController)
}
