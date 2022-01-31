package ReturnData

import (
	"github.com/gin-gonic/gin"
)

func Code200(Data string, c **gin.Context) {
	basicDataReturn(200, "Data", Data, c)
}

func Code500(Data string, c **gin.Context) {
	basicDataReturn(500, "error", Data, c)
}

func Code501(Data string, c **gin.Context) {
	basicDataReturn(501, "error", Data, c)
}

func basicDataReturn(Code int, Key string, Data string, c **gin.Context) {
	(*c).JSON(Code, gin.H{
		Key: Data,
	})
}
