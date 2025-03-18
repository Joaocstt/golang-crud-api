package repository

import (
	"api/internal/database"
	"api/internal/models"
	"github.com/google/uuid" // Importa a biblioteca para gerar UUIDs (Identificadores Únicos)
)

// CreateUser insere um novo usuário no banco de dados
func CreateUser(user *models.User) (*models.User, error) {
	// Gera um ID único para o usuário antes de salvá-lo no banco
	user.ID = uuid.New().String()

	// Insere o usuário no banco de dados
	res := database.DB.Create(user)

	// Se houver erro ao salvar no banco, retorna um erro
	if res.Error != nil {
		return nil, res.Error
	}

	// Se não houver erro, retorna o usuário criado
	return user, nil
}
