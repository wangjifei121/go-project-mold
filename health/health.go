package health

import (
	"net/http"
	"projectDemo/global"

	"github.com/gin-gonic/gin"
)


func Health(c *gin.Context) {
	message := "OK"
	global.ProjectLog.Info("router health request info"," OK")
	c.String(http.StatusOK, "\n"+message)
}
