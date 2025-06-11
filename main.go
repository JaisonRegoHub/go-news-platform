package main

import (
	"news-platform/config"
	"news-platform/routes"
	"news-platform/utils"
)

func main() {
	config.ConnectDB()
	utils.AutoMigrate()

	r := routes.SetupRoutes()
	r.Run(":8080")
}
