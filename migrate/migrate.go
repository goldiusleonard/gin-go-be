package main

import (
	"example.com/main/initializers"
	"example.com/main/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Item{})
}