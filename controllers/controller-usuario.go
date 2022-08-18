package controllers

import (
	"net/http"

	"github.com/Natanael-devops/api-orcamento-familiar/database"
	"github.com/Natanael-devops/api-orcamento-familiar/models"
	"github.com/Natanael-devops/api-orcamento-familiar/services"
	"github.com/gin-gonic/gin"
)

func CriaNovoUsuario(c *gin.Context) {
	var usuario models.Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	email := database.DB.Where("email = ?", usuario.Email).Find(&usuario)
	if email.RowsAffected > 0 {
		c.JSON(400, gin.H{"MENSAGEM": "esse email já está cadastrado em nosso sistema!"})
		c.JSON(400, email)
		return
	}

	usuario.Senha = services.SHA256Encoder(usuario.Senha)

	database.DB.Create(&usuario)
	c.JSON(http.StatusOK, gin.H{
		"data": "Cadastro realizado com sucesso!"})
}
