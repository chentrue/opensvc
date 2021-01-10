package routers

import (
	"github.com/gin-gonic/gin"
	"opensvc/routers/api"
)

func InitRouters () *gin.Engine{
	r := gin.New()
	stack_nova := r.Group("/nova")
	{
		stack_nova.GET("/list",api.List)
	}
	return r
}
