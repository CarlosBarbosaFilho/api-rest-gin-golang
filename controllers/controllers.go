package controllers

import (
	"github.com/CarlosBarbosaFilho/api-rest-go-gin/models"
	"github.com/gin-gonic/gin"
)

func Personas(c *gin.Context) {
	c.JSON(200, models.Peronas)
}

func Salutation(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API say:": "Hey " + name + ", very nice ?",
	})
}
