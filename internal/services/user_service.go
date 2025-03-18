package services

import (
	"api/internal/models"
	"api/internal/repository"
)

// CreateUserService processa a lógica de criação de um novo usuário antes de salvar no banco
func CreateUserService(user *models.User) (*models.User, error) {
	// Chama a função CreateUser no repositório para salvar o usuário no banco de dados
	return repository.CreateUser(user)
}

func GetUserListService() ([]*models.User, error) {
	return repository.GetUserList()
}
