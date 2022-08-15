package controllers

import (
	"net/http"

	"github.com/Natanael-devops/api-orcamento-familiar/database"
	"github.com/Natanael-devops/api-orcamento-familiar/models"
	"github.com/gin-gonic/gin"
)

func GastosCategorias(m, categorias string) int {
	var somaCategorias int
	var existe models.Despesa
	if err := database.DB.Where(&models.Despesa{Data: m, Categoria: categorias}).First(&existe).Error; err != nil {
		return 0
	}
	if existe.ID == 0 {

		return 0
	} else {
		database.DB.Table("despesas").Select("SUM(valor)").Where("data= ? AND deleted_at IS NULL AND categoria = ?", m, categorias).Find(&somaCategorias)
		return somaCategorias
	}
}

func DevolveResumoMensal(c *gin.Context) {
	var resumo models.Resumo

	m := c.Param("mes")
	a := c.Param("id")
	m += "/"
	m += a

	//somas
	var somaReceita int
	var somaDespesa int
	database.DB.Table("receita").Select("SUM(valor)").Where("data= ? AND deleted_at IS NULL", m).Find(&somaReceita)
	database.DB.Table("despesas").Select("SUM(valor)").Where("data= ? AND deleted_at IS NULL", m).Find(&somaDespesa)
	resumo.TotalReceitas = somaReceita
	resumo.TotalDespesas = somaDespesa
	resumo.Saldo = somaReceita - somaDespesa
	// ----**----

	//gastos
	var categorias = []string{"alimentação", "saúde", "moradia", "transporte", "educação", "lazer", "imprevistos", "outras"}
	resumo.GastoAlimentacao = GastosCategorias(m, categorias[0])
	resumo.GastoSaude = GastosCategorias(m, categorias[1])
	resumo.GastoMoradia = GastosCategorias(m, categorias[2])
	resumo.GastoTransporte = GastosCategorias(m, categorias[3])
	resumo.GastoEducacao = GastosCategorias(m, categorias[4])
	resumo.GastoLazer = GastosCategorias(m, categorias[5])
	resumo.GastoImprevistos = GastosCategorias(m, categorias[6])
	resumo.GastoOutras = GastosCategorias(m, categorias[7])

	c.JSON(http.StatusOK, resumo)

}
