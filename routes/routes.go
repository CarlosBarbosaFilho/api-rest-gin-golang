package routes

import (
	"github.com/CarlosBarbosaFilho/api-rest-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/personas", controllers.Personas)
	r.GET("/personas/:id", controllers.GetByPersonaId)
	r.GET("/personas/document/:document", controllers.GetByPersonaDocument)
	r.POST("/personas", controllers.CreatePersona)
	r.DELETE("/personas/:id", controllers.RemovePersona)
	r.PUT("/personas/:id", controllers.EditPersona)
	r.Run() // listen and serve on 0.0.0.0:8080
}
