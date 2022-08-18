package routes

import (
	"github.com/Natanael-devops/api-orcamento-familiar/controllers"
	"github.com/Natanael-devops/api-orcamento-familiar/middleware"
	"github.com/gin-gonic/gin"
)

func CarregaRotas() {
	r := gin.Default()

	//rotas receitas
	r.GET("/receitas/", middleware.Auth(), controllers.ExibeTodasReceitas)
	r.GET("/receitas/:id", middleware.Auth(), controllers.DetalhaReceitaPorID)
	r.GET("/receitas", middleware.Auth(), controllers.BuscaReceitaPorDescricao)
	r.GET("/receitas/:id/:mes", middleware.Auth(), controllers.BuscaReceitaPorMes)
	r.POST("/receitas", middleware.Auth(), controllers.CriaNovaReceita)
	r.DELETE("/receitas/:id", middleware.Auth(), controllers.DeletaReceita)
	r.PUT("/receitas/:id", middleware.Auth(), controllers.EditaReceita)
	// rotas despesas
	r.GET("/despesas/", middleware.Auth(), controllers.ExibeTodasDespesas)
	r.GET("/despesas/:id", middleware.Auth(), controllers.DetalhaDespesaPorID)
	r.GET("/despesas", middleware.Auth(), controllers.BuscaDespesaPorDescricao)
	r.GET("/despesas/:id/:mes", middleware.Auth(), controllers.BuscaDespesaPorMes)
	r.POST("/despesas", middleware.Auth(), controllers.CriaNovaDespesa)
	r.DELETE("/despesas/:id", middleware.Auth(), controllers.DeletaDespesa)
	r.PUT("/despesas/:id", middleware.Auth(), controllers.EditaDespesa)
	//rotas resumo
	r.GET("/resumo/:id/:mes", middleware.Auth(), controllers.DevolveResumoMensal)

	//
	r.POST("/", controllers.CriaNovoUsuario)
	r.POST("/login", controllers.Login)
	r.Run()
}
