package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"opensvc/pkg/client"
	"opensvc/pkg/stack"
	"opensvc/pkg/utils"
)


func main(){

}

func List (c *gin.Context) {
	region := c.GetHeader("region")
	appG := utils.Gin{C: c}
	com := stack.Server1{
		new(client.Compute),
		}
	result, err := com.Client.List(region)
	if err != nil {
		appG.Response(http.StatusBadRequest, 401, result)
	}
	appG.Response(http.StatusOK, 200, result)

}