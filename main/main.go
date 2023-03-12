package main

import (
	"github.com/CarlosBarbosaFilho/api-rest-go-gin/database"
	"github.com/CarlosBarbosaFilho/api-rest-go-gin/routes"
)

func main() {
	database.ConnectionDB()
	routes.HandleRequests()

}
