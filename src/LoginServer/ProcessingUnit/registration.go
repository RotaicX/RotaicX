package ProcessingUnit

import "github.com/gin-gonic/gin"

func Registration(Username string, Password string, c **gin.Context) {
	(*c).JSON(200, gin.H{
		"Data": "Under Development",
	})
}
