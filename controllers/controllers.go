package controllers

import (
	"net/http"

	"github.com/cali0p3/api-go-gin/database"
	"github.com/cali0p3/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "Olá " + nome + ", tudo tranquilo?",
	})
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err!= nil{
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error()})
			return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}