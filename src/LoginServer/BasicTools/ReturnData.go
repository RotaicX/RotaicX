package BasicTools

import (
	"github.com/gin-gonic/gin"
	"log"
)

func ReturnData(Type int, Data string, c **gin.Context) {
	switch Type {
	case 0:
		(*c).JSON(200, gin.H{
			"Data": Data,
		})
	case 1:
		(*c).JSON(501, gin.H{
			"Error": Data,
		})
	default:
		log.Panicln("Type does not exist")
	}
}
