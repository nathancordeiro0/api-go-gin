package main

import (
	"github.com/joho/godotenv"
	"github.com/nathancordeiro0/api-go-gin/database"
	"github.com/nathancordeiro0/api-go-gin/routes"
)

func main() {
	godotenv.Load()
	database.ConnectDatabase()
	routes.HandleRequests()
}
