package main

import (
	
	"GoGinApiRest/routes"
	"GoGinApiRest/database"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}