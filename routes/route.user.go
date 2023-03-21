package routes

import (
	"backend-rltrading/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {

	//All routes related to users comes here
	router.POST("/user", controllers.CreateUser())
	router.GET("users", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "get List users",
		})
	})

	router.GET("user/:userId", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "end",
		})
	})
}
