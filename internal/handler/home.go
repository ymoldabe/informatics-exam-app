package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HandlerType) Home(c *gin.Context) {

	grades := h.service.GetGrade()
	c.HTML(http.StatusOK, "home.html", gin.H{
		"Grades": grades,
	})
}
