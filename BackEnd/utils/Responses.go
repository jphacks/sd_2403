package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
}

func SuccessResponse(c *gin.Context, msg string, data interface{}) {
	if msg == "" {
		msg = "success"
	}
	c.JSON(
		http.StatusOK,
		Response{
			Code: 200,
			Msg:  msg,
			Data: data,
		},
	)
}
func ErrorResponse(c *gin.Context, code int, msg string, data interface{}) {
	if msg == "" {
		msg = "success"
	}
	c.JSON(
		code,
		Response{
			Code: code,
			Msg:  msg,
			Data: data,
		},
	)
	c.AbortWithStatus(code)
}
