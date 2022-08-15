package controllers

import (
	"net/http"

	"github.com/Natanael-devops/api-orcamento-familiar/database"
	"github.com/Natanael-devops/api-orcamento-familiar/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodasReceitas(c *gin.Context) {
	var receitas []models.Receita
	database.DB.Find(&receitas)
	c.JSON(200, receitas)
}

func CriaNovaReceita(c *gin.Context) {
	var receita models.Receita
	if err := c.ShouldBindJSON(&receita); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidaDadosReceita(&receita); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	descricao := database.DB.Where("descricao = ? AND data = ?", receita.Descricao, receita.Data).Find(&receita)
	if descricao.RowsAffected > 0 {
		c.JSON(400, gin.H{"MENSAGEM": "essa descrição já existe neste mês"})
		c.JSON(400, descricao)
		return
	}

	database.DB.Create(&receita)
	c.JSON(http.StatusOK, receita)
}

func DetalhaReceitaPorID(c *gin.Context) {
	var receita models.Receita
	id := c.Params.ByName("id")
	database.DB.First(&receita, id)

	if receita.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Receita não encontrada"})
		return
	}

	c.JSON(http.StatusOK, receita)
}

func DeletaReceita(c *gin.Context) {
	var receita models.Receita
	id := c.Params.ByName("id")
	database.DB.Delete(&receita, id)
	c.JSON(http.StatusOK, gin.H{
		"data": "Receita deletada com sucesso!"})
}

func EditaReceita(c *gin.Context) {
	var receita models.Receita
	id := c.Params.ByName("id")
	database.DB.First(&receita, id)

	if err := c.ShouldBindJSON(&receita); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidaDadosReceita(&receita); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&receita).UpdateColumns(receita)
	c.JSON(http.StatusOK, receita)
}

func BuscaReceitaPorDescricao(c *gin.Context) {
	var receita []models.Receita
	descricao := c.Query("descricao")
	database.DB.Where(&models.Receita{Descricao: descricao}).Find(&receita)

	if len(receita) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Receita não encontrada"})
		return
	}

	c.JSON(http.StatusOK, receita)
}

func BuscaReceitaPorMes(c *gin.Context) {
	var receitas []models.Receita
	m := c.Param("mes")
	a := c.Param("id")
	m += "/"
	m += a
	database.DB.Where(&models.Receita{Data: m}).Find(&receitas)
	if len(receitas) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Não há receitas nesse mês!"})
		return
	}
	c.JSON(http.StatusOK, receitas)
}
