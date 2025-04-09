package main

import (
	"GoGym/config"
	"GoGym/routes"
)

func main() {
	config.SetupDatabase()
	
	routes.HandleRequest(config.DB)

}
