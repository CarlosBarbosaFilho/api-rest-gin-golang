package routes

import (
	"github.com/CarlosBarbosaFilho/api-rest-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/personas", controllers.Personas)
	r.GET("/personas/:name", controllers.Salutation)
	r.Run() // listen and serve on 0.0.0.0:8080
}
