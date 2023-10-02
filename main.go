package main

import (
	"github.com/cali0p3/api-go-gin/database"
	
	"github.com/cali0p3/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	
	routes.HandleRequests()
}
