package main

import (
	"github.com/LeonardoMedeiross/api-go-gin/database"
	"github.com/LeonardoMedeiross/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
