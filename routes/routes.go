package routes

import (
	"github.com/Natanael-devops/api-orcamento-familiar/controllers"
	"github.com/gin-gonic/gin"
)

func CarregaRotas() {
	r := gin.Default()
	//rotas receitas
	r.GET("/receitas/", controllers.ExibeTodasReceitas)
	r.GET("/receitas/:id", controllers.DetalhaReceitaPorID)
	r.GET("/receitas", controllers.BuscaReceitaPorDescricao)
	r.GET("/receitas/:id/:mes", controllers.BuscaReceitaPorMes)
	r.POST("/receitas", controllers.CriaNovaReceita)
	r.DELETE("/receitas/:id", controllers.DeletaReceita)
	r.PUT("/receitas/:id", controllers.EditaReceita)
	// rotas despesas
	r.GET("/despesas/", controllers.ExibeTodasDespesas)
	r.GET("/despesas/:id", controllers.DetalhaDespesaPorID)
	r.GET("/despesas", controllers.BuscaDespesaPorDescricao)
	r.GET("/despesas/:id/:mes", controllers.BuscaDespesaPorMes)
	r.POST("/despesas", controllers.CriaNovaDespesa)
	r.DELETE("/despesas/:id", controllers.DeletaDespesa)
	r.PUT("/despesas/:id", controllers.EditaDespesa)
	//rotas resumo
	r.GET("/resumo/:id/:mes", controllers.DevolveResumoMensal)

	//
	r.Run()
}
