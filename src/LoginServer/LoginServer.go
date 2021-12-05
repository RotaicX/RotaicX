package main

import (
	"LoginServer/BasicTools"
	"LoginServer/ProcessingUnit"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	defer func() {
		if recover() != nil {
			log.Fatalln("A fatal error occurred and the program could not continue to run")
		}
	}()
	r := gin.Default()
	r.POST("/AccountServices", AccountServices)
	r.Run()
}

func AccountServices(c *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			ReturnData.Code500("An internal server error occurred and therefore your request could not be processed", &c)
		}
	}()

	json := make(map[string]interface{})
	c.BindJSON(&json)

	switch json["type"] {
	case "registration":
		ProcessingUnit.Registration(json["username"], json["password"], &c)
	case nil, "":
		ReturnData.Code501("Type cannot be empty!!!!", &c)
	default:
		ReturnData.Code501(fmt.Sprintf("Instruction %s does not exist", json["type"]), &c)
	}
}
