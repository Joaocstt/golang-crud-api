package database

import (
	"api/internal/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load("../.env")
	if err != nil {
		// Se houver erro ao carregar o arquivo .env, a aplicação é finalizada
		log.Fatal("Error loading .env file")
	}

	// Lê as variáveis de ambiente para configurar a conexão com o banco de dados MySQL
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Verifica se todas as variáveis necessárias estão presentes
	if dbUser == "" || dbHost == "" || dbPort == "" || dbName == "" {

		// Se faltar alguma variável de configuração, a aplicação é finalizada
		log.Fatal("Missing database configuration in .env file")
	}

	// Formata a string de conexão (DSN - Data Source Name) para o MySQL
	// Exemplo de DSN: "user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Tenta abrir a conexão com o banco de dados MySQL utilizando o GORM
	db, errDB := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDB != nil {
		// Se ocorrer erro na conexão, a aplicação é finalizada
		log.Fatalf("Failed to connect to database: %v", errDB)
	}

	// Rodar as migrações

	errMigrate := db.AutoMigrate(&models.User{})

	if errMigrate != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Se a conexão for bem-sucedida, exibe a mensagem de sucesso no terminal
	fmt.Printf("Database '%s' connected successfully!\n", dbName)

	DB = db

}
