package controllers

import (
	"example.com/main/initializers"
	"example.com/main/models"
	"github.com/gin-gonic/gin"
)

func ItemsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Name	string
		Price	float32
	}
	
	c.Bind(&body)

	// Create an item
	item := models.Item{Name: body.Name, Price: body.Price}

	result := initializers.DB.Create(&item)

	if result.Error != nil {
		c.Status(400)
		return 
	}

	c.JSON(200, gin.H{
		"item": item,
	})
}

func ItemsIndex(c *gin.Context) {
	// Get the items
	var items []models.Item
	initializers.DB.Find(&items)

	c.JSON(200, gin.H{
		"items": items,
	})
}

func ItemsShow(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the items
	var item models.Item
	initializers.DB.First(&item, id)

	// Respond with them
	c.JSON(200, gin.H{
		"item": item,
	})
}

func ItemsUpdate(c *gin.Context) {
	// Get id off the url
	id := c.Param("id")

	// Get data off req body
	var body struct {
		Name	string
		Price	float32
	}
	
	c.Bind(&body)

	// Find the post were updating
	var item models.Item
	initializers.DB.First(&item, id)

	// Update it
	initializers.DB.Model(&item).Updates(models.Item{
		Name: body.Name,
		Price: body.Price,
	})

	// Respond it
	c.JSON(200, gin.H{
		"item": item,
	})
}

func ItemsDelete(c *gin.Context) {
	// Get id off the url
	id := c.Param("id")
	
	// Delete the posts
	initializers.DB.Delete(&models.Item{}, id)

	// Respond it
	c.Status(200)
}