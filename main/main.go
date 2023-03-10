package main

import (
	"github.com/CarlosBarbosaFilho/api-rest-go-gin/models"
	"github.com/CarlosBarbosaFilho/api-rest-go-gin/routes"
)

func main() {

	models.Peronas = []models.Persona{
		{Name: "Carlos Barbosa", Email: "cbgomes@gmail.com", Phone: "83991267778", Document: "03100721403"},
	}
	routes.HandleRequests()
}
