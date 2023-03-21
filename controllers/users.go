package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateUser	godoc
// @Summary	Create new user
// @Tags	users
// @Produce	json
// @Success	200
// @Router	/api/v1/user [post]
func CreateUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusCreated, gin.H{
			"message": "create users successful",
		})
	}
}

func UsersList() gin.HandlerFunc {

}