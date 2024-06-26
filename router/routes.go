package router

import (
	"github.com/gin-gonic/gin"
	"github.com/imrochamatheus/challenge-back-end-hit/handler"
)

func initializeRoutes(router *gin.Engine) {
	basePath := "/api/v1"

	handler.InitializeHandlers()

	v1 := router.Group(basePath)
	{
		v1.GET("/planet", handler.ListPlanetsHandler)
		v1.POST("/planet", handler.CreatePlanetHandler)
		v1.DELETE("/planet/:id", handler.DeletePlanetHandler)
	}
}
