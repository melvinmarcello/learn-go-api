package main

import (
	"go-api/models"
	"go-api/routes"
)

func main() {

	models.ConnectDatabase()
	routes.Router()

}
