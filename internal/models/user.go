package models

import (
	"api/internal/config"
)

// User representa a estrutura do usuário no sistema e no banco de dados
type User struct {
	ID        string `json:"id" gorm:"primary_key"`
	Name      string `json:"name" gorm:"unique;not null" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Telephone string `json:"telephone" validate:"required"`
}

// Validate verifica se os dados do usuário atendem às regras de validação
func (u *User) Validate() error {
	// Usa o validador global definido no pacote config para validar a struct User
	return config.Validator.Struct(u)
}
