package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterInput struct {
	Username string `json:"username" bidding:"required"`
	Password string `json:"password" bidding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "valdated"})
}
