package main

import (
	"api/internal/database"
	"api/internal/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Inicializa a conexão com o banco de dados (mysql)
	database.InitDB()

	// Cria uma nova instância do servidor Gin com configurações padrão
	// O router é um ponteiro do tipo *gin.Engine
	router := gin.Default()

	// Inicializa as rotas do servidor, passando o objeto 'router' para o pacote de rotas
	// Não precisa passar o &router porque router já é um ponteiro de gin.Engine
	// Passando o ponteiro 'router' diretamente para InitRoutes
	routes.InitRoutes(router)

	// Inicia o servidor na porta 8080 e verifica se houve algum erro
	err := router.Run(":8080")

	// Registrar se ocorreu algum erro ao tentar iniciar o servidor
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
