package main

import (
	"GoGym/config"
	"GoGym/routes"
)

func main() {
	config.Connect()

	routes.HandleRequest()
}
