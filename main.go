package main

import (
	"restapi/database"
	"restapi/server"
)

func main() {

	database.CreateClient()
	server.Start()
}
