package main

import (
	"api-go-rest/database"
	"api-go-rest/routes"
	"fmt"
)

func main() {
	/*
		models.Personalidades = []models.Personalidade{
			{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
			{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},

			teste git
		}*/

	database.ConectaComBD()

	fmt.Println("Iniciando o Servidor Rest com Go")
	routes.HandleRequest()
}
