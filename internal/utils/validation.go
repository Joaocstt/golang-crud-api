package utils

import "github.com/go-playground/validator/v10"

// FormatarValidationError recebe um erro de validação e retorna um mapa com mensagens amigáveis
func FormatarValidationError(err error) map[string]string {
	// Cria um mapa vazio onde a chave será o nome do campo e o valor será a mensagem de erro
	errs := make(map[string]string)

	// Converte o erro genérico para um tipo específico de erro de validação (ValidationErrors)
	// Assim podemos percorrer os erros específicos de cada campo
	for _, err := range err.(validator.ValidationErrors) {
		// Adiciona a mensagem de erro correspondente ao campo no mapa
		errs[err.Field()] = getErrorMessage(err)
	}

	// Retorna o mapa contendo os erros formatados
	return errs
}

// getErrorMessage recebe um erro específico e retorna uma mensagem personalizada
func getErrorMessage(err validator.FieldError) string {
	// Verifica qual foi a regra de validação que falhou e retorna uma mensagem apropriada
	switch err.Tag() {
	case "required":
		return "O campo é obrigatório"
	default:
		// Caso não seja uma validação tratada, retorna uma mensagem genérica
		return "Erro no campo " + err.Field()
	}
}
