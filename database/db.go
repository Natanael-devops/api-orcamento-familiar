package database

import (
	"log"

	"github.com/Natanael-devops/api-orcamento-familiar/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=postgres user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Receita{})
	DB.AutoMigrate(&models.Despesa{})
	DB.AutoMigrate(&models.Usuario{})
}
