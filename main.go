package main

import (
	"github.com/Natanael-devops/api-orcamento-familiar/database"
	"github.com/Natanael-devops/api-orcamento-familiar/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.CarregaRotas()
}
