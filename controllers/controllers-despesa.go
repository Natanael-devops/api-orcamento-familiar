package controllers

import (
	"net/http"

	"github.com/Natanael-devops/api-orcamento-familiar/database"
	"github.com/Natanael-devops/api-orcamento-familiar/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodasDespesas(c *gin.Context) {
	var despesas []models.Despesa
	database.DB.Find(&despesas)
	c.JSON(200, despesas)
}

func CriaNovaDespesa(c *gin.Context) {
	var despesa models.Despesa

	if err := c.ShouldBindJSON(&despesa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidaDadosDespesa(&despesa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	descricao := database.DB.Where("descricao = ? AND data = ?", despesa.Descricao, despesa.Data).Find(&despesa)
	if descricao.RowsAffected > 0 {
		c.JSON(400, gin.H{"MENSAGEM": "essa descrição já existe neste mês"})
		c.JSON(400, descricao)
		return
	}
	despesa.Categoria = models.ValidaCategoria(&despesa)
	database.DB.Create(&despesa)
	c.JSON(http.StatusOK, despesa)
}

func DetalhaDespesaPorID(c *gin.Context) {
	var despesa models.Despesa
	id := c.Params.ByName("id")
	database.DB.First(&despesa, id)

	if despesa.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Despesa não encontrada"})
		return
	}

	c.JSON(http.StatusOK, despesa)
}

func DeletaDespesa(c *gin.Context) {
	var despesa models.Despesa
	id := c.Params.ByName("id")
	database.DB.Delete(&despesa, id)
	c.JSON(http.StatusOK, gin.H{
		"data": "Despesa deletada com sucesso!"})
}

func EditaDespesa(c *gin.Context) {
	var despesa models.Despesa
	id := c.Params.ByName("id")
	database.DB.First(&despesa, id)

	if err := c.ShouldBindJSON(&despesa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidaDadosDespesa(&despesa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&despesa).UpdateColumns(despesa)
	c.JSON(http.StatusOK, despesa)
}

func BuscaDespesaPorDescricao(c *gin.Context) {
	var despesa []models.Despesa
	descricao := c.Query("descricao")
	database.DB.Where(&models.Despesa{Descricao: descricao}).Find(&despesa)

	if len(despesa) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Despesa não encontrada"})
		return
	}

	c.JSON(http.StatusOK, despesa)
}

func BuscaDespesaPorMes(c *gin.Context) {
	var despesas []models.Despesa
	m := c.Param("mes")
	a := c.Param("id")
	m += "/"
	m += a
	database.DB.Where(&models.Despesa{Data: m}).Find(&despesas)
	if len(despesas) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Não há despesas nesse mês!"})
		return
	}
	c.JSON(http.StatusOK, despesas)
}
