package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/gofiber/fiber/v2"
)

func main() {
	var data_store string = ""
	fmt.Println("Gin framework")
	router := gin.Default()

	router.POST("/key", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "Gin POST request success"})
		// fmt.Print(&requestBody)
		var requestBody struct {
			Message string `json:"key"`
		}
		err := c.BindJSON(&requestBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(requestBody)
		data_store = requestBody.Message
	})

	router.GET("/all", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": data_store,
		})
	})

	router.Run(":8080")
}
