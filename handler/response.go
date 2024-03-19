package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/imrochamatheus/challenge-back-end-hit/models"
)

func sendError(ctx *gin.Context, code int, msg string) {
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

type CreatePlanetResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Ground      string `json:"ground"`
	Climate     string `json:"climate"`
	Appearances int64  `json:"appearances"`
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type ListPlanetsResponse struct {
	Message string          `json:"message"`
	Data    []models.Planet `json:"data"`
}
