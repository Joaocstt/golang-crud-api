package handlers

import (
	"api/internal/models"
	"api/internal/response"
	"api/internal/services"
	"api/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PostUser é um handler responsável por criar um novo usuário
func PostUser(ctx *gin.Context) {
	var user models.User // Declara uma variável do tipo User para armazenar os dados da requisição

	// Faz o binding do JSON recebido para a struct User
	errBadRequest := ctx.ShouldBindJSON(&user)

	// Se houver erro no binding (exemplo: JSON mal formatado), retorna um erro 400 (Bad Request)
	if errBadRequest != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errBadRequest.Error()})
		return
	}

	// Valida os dados do usuário antes de continuar
	if err := user.Validate(); err != nil {
		// Se a validação falhar, retorna um erro 400 (Bad Request) com mensagens formatadas
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.FormatarValidationError(err)})
		return
	}

	// Chama o serviço para criar o usuário no banco de dados
	res, errInternalServerError := services.CreateUserService(&user)

	// Se houver erro interno ao salvar o usuário no banco, retorna um erro 500 (Internal Server Error)
	if errInternalServerError != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errInternalServerError.Error()})
		return
	}

	// Retorna uma resposta JSON com mensagem de sucesso e os dados do usuário sem informações sensíveis
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Usuário criado com sucesso!",
		"user":    response.NewUserResponse(res), // Converte User para UserResponse para esconder dados sensíveis
	})
}

func GetUser(ctx *gin.Context) {
	res, err := services.GetUserListService()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(res) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Nenhum usuário encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Usuários listado com sucesso!",
		"user":    res,
	})
}
