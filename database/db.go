package database

import (
	"log"
	"os"

	"github.com/Natanael-devops/api-orcamento-familiar/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	if os.Getenv("ENV") == "PROD" {
		ambiente := os.Getenv("DATABASE_URL")
		DB, err = gorm.Open(postgres.Open(ambiente), &gorm.Config{})
		if err != nil {
			log.Panic("Erro ao conectar com banco de dados")
		}
	} else {
		stringDeConexao := "host=postgres user=root password=root dbname=root port=5432 sslmode=disable"
		DB, err = gorm.Open(postgres.Open(stringDeConexao))
	}

	DB.AutoMigrate(&models.Receita{})
	DB.AutoMigrate(&models.Despesa{})
	DB.AutoMigrate(&models.Usuario{})
}
