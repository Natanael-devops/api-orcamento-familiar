package routes

import (
	"github.com/Natanael-devops/api-orcamento-familiar/controllers"
	"github.com/gin-gonic/gin"
)

func CarregaRotas() {
	r := gin.Default()
	//rotas receitas
	r.GET("/receitas", controllers.ExibeTodasReceitas)
	r.GET("/receitas/:id", controllers.DetalhaReceitaPorID)
	r.POST("/receitas", controllers.CriaNovaReceita)
	r.DELETE("/receitas/:id", controllers.DeletaReceita)
	r.PUT("/receitas/:id", controllers.EditaReceita)
	// rotas despesas
	r.GET("/despesas", controllers.ExibeTodasDespesas)
	r.GET("/despesas/:id", controllers.DetalhaDespesaPorID)
	r.POST("/despesas", controllers.CriaNovaDespesa)
	r.DELETE("/despesas/:id", controllers.DeletaDespesa)
	r.PUT("/despesas/:id", controllers.EditaDespesa)
	r.Run()
}
