package main

import (
	"github.com/LeonardoMedeiross/api-go-gin/models"
	"github.com/LeonardoMedeiross/api-go-gin/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Leonardo", CPF: "1234567850", RG: "36925814"},
		{Nome: "Crystiane", CPF: "00000000000", RG: "37925814"},
	}
	routes.HandleRequest()
}
