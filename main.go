package main

import (
	"github.com/DanielDDHM/pool-api/config"
	"github.com/DanielDDHM/pool-api/server"
)

func main() {

	server := server.App()

	config.ConnectDB()

	server.Run()
}
