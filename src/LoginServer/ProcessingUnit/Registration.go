package ProcessingUnit

import (
	"LoginServer/BasicTools/ReturnData"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Registration(RequestData map[string]interface{}, c *gin.Context) {
	ReturnData.Code200(fmt.Sprintf("Under Development, You Username %s, Password %s", RequestData["Username"], RequestData["Password"]), c)
}
