package main

import (
	"github.com/Carlos-GitH/go-rest-api/database"
	"github.com/Carlos-GitH/go-rest-api/models"
	"github.com/Carlos-GitH/go-rest-api/routes"
)

func main() {
	models.Peronalidades = []models.Personalidade{
		{Id: 1, Nome: "Carlos", Historia: "Historia"},
		{Id: 2, Nome: "Jubileu", Historia: "Historia"},
	}
	database.ConectaComBancoDeDados()
	println("Iniciando o servidor!")
	routes.HandleRequest()
}
