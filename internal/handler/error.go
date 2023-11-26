package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *HandlerType) NewError(c *gin.Context, code int, message string) {
	logrus.Errorf(message)

	c.HTML(code, "error.html", gin.H{
		"ErrorCode": code,
		"ErrorMsg":  message,
	})
}
