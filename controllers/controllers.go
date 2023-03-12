package controllers

import (
	"net/http"

	"github.com/CarlosBarbosaFilho/api-rest-go-gin/database"
	"github.com/CarlosBarbosaFilho/api-rest-go-gin/models"
	"github.com/gin-gonic/gin"
)

func Personas(c *gin.Context) {
	var personas []models.Persona

	database.DB.Find(&personas)

	c.JSON(http.StatusAccepted, personas)
}

func Salutation(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API say:": "Hey " + name + ", very nice ?",
	})
}

func GetByPersonaDocument(c *gin.Context) {
	var persona models.Persona

	document := c.Param("document")
	database.DB.Where(&models.Persona{Document: document}).First(&persona)

	if persona.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Persona not found or not exists"})
		return
	}

	c.JSON(http.StatusOK, persona)
}

func CreatePersona(c *gin.Context) {
	var persona models.Persona
	if err := c.ShouldBindJSON(&persona); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&persona)
	c.JSON(http.StatusCreated, persona)
}

func GetByPersonaId(c *gin.Context) {
	var persona models.Persona
	id := c.Params.ByName("id")
	database.DB.First(&persona, id)

	if persona.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"not found": "Sorry, persona not found"})
		return
	}
	c.JSON(http.StatusOK, persona)

}

func RemovePersona(c *gin.Context) {
	var persona models.Persona
	id := c.Params.ByName("id")
	database.DB.Delete(&persona, id)

	c.JSON(http.StatusOK, gin.H{"message": "Persona removed with success"})

}

func EditPersona(c *gin.Context) {
	var persona models.Persona
	id := c.Params.ByName("id")

	database.DB.First(&persona, id)

	if err := c.ShouldBindJSON(&persona); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&persona).UpdateColumns(persona)
	c.JSON(http.StatusOK, gin.H{"message": "persona edited with success"})
}
