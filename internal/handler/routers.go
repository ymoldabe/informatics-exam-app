package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *HandlerType) InitRouters() *gin.Engine {

	router := gin.New()
	router.LoadHTMLGlob("./ui/html/*.html")
	router.Static("/static", "./ui/static/")

	api := router.Group("/api")
	{
		api.GET("/home", h.Home)

		testing := api.Group("/testing")
		{
			testing.GET("/:grades", h.Testing)
			testing.POST("/:grades", h.TestingResult)
		}
	}

	return router
}
