package main

import (
	"example.com/main/controllers"
	"example.com/main/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/items", controllers.ItemsCreate)
	r.PUT("/items/:id", controllers.ItemsUpdate)
	r.GET("/items", controllers.ItemsIndex)
	r.GET("/items/:id", controllers.ItemsShow)
	r.DELETE("/items/:id", controllers.ItemsDelete)
	
	r.Run(":8000") // listen and serve on 0.0.0.0:8080
}