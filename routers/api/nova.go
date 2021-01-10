package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"opensvc/pkg/stack"
	"opensvc/pkg/utils"
)

func List (c *gin.Context) {
	region := c.GetHeader("region")
	appG := utils.Gin{C: c}
	result, err := stack.Server_list(region)
	if err != nil {
		appG.Response(http.StatusBadRequest, 401, result)
	}
	appG.Response(http.StatusOK, 200, result)

}