package routes

import (
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			UserRoutes(v1)
			v1.GET("/test",
				// GetListUsers godoc
				// @Summary Get list users
				// @Description Response with the list of all users as JSON.
				// @Tags users
				func(context *gin.Context) {
					context.JSON(200, gin.H{
						"message": "testkkkk",
					})
				})
		}

	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

}

func SwaggerInit() {

}
