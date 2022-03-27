package main

import (
	"LoginServer/BasicTools"
	ReturnData2 "LoginServer/BasicTools/ReturnData"
	"LoginServer/ProcessingUnit"
	"flag"
	"fmt"
	RotaicxBasicTool "github.com/RotaicX/RotaicXBasicTool"
	"github.com/gin-gonic/gin"
	"os"
)

var LoginServerConfig BasicTools.TomlConfig

func main() {
	defer func() {
		err := recover()
		if err != nil {
			RotaicxBasicTool.Rlog.Errorln("A fatal error occurred and the program could not continue to run")
			os.Exit(1)
		}
	}()

	// LoginServer Init
	/// Init Args
	var ConfigFilePath = flag.String("f", "Config.toml", "You can specify the location of the configuration file here")
	flag.Parse()
	LoginServerConfig = BasicTools.LoadConfigFile(*ConfigFilePath)
	/// Init Gin
	r := gin.Default()
	r.POST("/AccountServices", AccountServices)

	// LoginServer Start !!
	err := r.Run()
	if err != nil {
		panic(err)
	}
}

func AccountServices(c *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			ReturnData2.Code500("An internal server error occurred and therefore your request could not be processed", c)
		}
	}()

	json := make(map[string]interface{})
	err := c.BindJSON(&json)
	if err != nil {
		panic(err)
	}

	switch json["type"] {
	case "registration":
		ProcessingUnit.Registration(json, c)
	case nil, "":
		ReturnData2.Code501("Type cannot be empty!!!!", c)
	default:
		ReturnData2.Code501(fmt.Sprintf("Instruction %s does not exist", json["type"]), c)
	}
}
