package main

import (
	"LoginServer/ProcessingUnit"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.POST("/api", api)
	r.Run()
}

func api(c *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			c.JSON(500, gin.H{
				"Error": "An internal server error occurred and therefore your request could not be processed",
			})
		}
	}()

	json := make(map[string]interface{})
	c.BindJSON(&json)
	log.Printf("%v", json)

	switch json["type"] {
	case "registration":
		ProcessingUnit.Registration("", "", &c)
	case nil:
		c.JSON(501, gin.H{
			"Error": "type cannot be empty!!!!",
		})
	default:
		c.JSON(501, gin.H{
			"Error": fmt.Sprintf("Instruction %s does not exist", json["type"]),
		})
	}
}
