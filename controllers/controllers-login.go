package controllers

import (
	"net/http"

	"github.com/Natanael-devops/api-orcamento-familiar/database"
	"github.com/Natanael-devops/api-orcamento-familiar/models"
	"github.com/Natanael-devops/api-orcamento-familiar/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var log models.Login
	if err := c.ShouldBindJSON(&log); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	var usuario models.Usuario
	dbErro := database.DB.Where("email = ?", log.Email).First(&usuario).Error
	if dbErro != nil {
		c.JSON(400, gin.H{
			"error": "Usuário não encontrado"})
		return
	}

	if usuario.Senha != services.SHA256Encoder(log.Senha) {
		c.JSON(401, gin.H{
			"error": "credenciais inválidas",
		})
		return
	}

	token, err := services.NewJWTService().GeraToken(usuario.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})

}
