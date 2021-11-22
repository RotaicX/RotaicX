package ProcessingUnit

import (
	"LoginServer/BasicTools"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Registration(Username interface{}, Password interface{}, c **gin.Context) {
	ReturnData.Code200(fmt.Sprintf("Under Development, You Username %s, Password %s", Username, Password), c)
}
