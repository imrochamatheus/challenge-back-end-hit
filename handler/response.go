package handler

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/imrochamatheus/challenge-back-end-hit/models"
)

func sendError(ctx *gin.Context, code int, msg string) {
	log.Printf(msg)

	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, code int, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type ListPlanetsResponse struct {
	Message string        	`json:"message"`
	Data    []models.Planet `json:"data"`
}

type CreatePlanetResponse struct {
	Message string              `json:"message"`
	Data    models.CreatePlanet `json:"data"`
}
