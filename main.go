package main

import (
	"github.com/cali0p3/api-go-gin/models"
	"github.com/cali0p3/api-go-gin/routes"
)

func main() {
	models.Alunos = []models.Aluno {
		{Nome: "Calíope Longoni", CPF: "12345678901", RG: "123456789"},
		{Nome: "Anna Ball", CPF: "21345678901", RG: "987654321"},
	}
	routes.HandleRequests()
}
