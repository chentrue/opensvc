package utils

import (
	"github.com/gin-gonic/gin"
	"opensvc/pkg/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`

}

func (g *Gin) Response (httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg: e.GMsg(errCode),
		Data: data,
	})
}