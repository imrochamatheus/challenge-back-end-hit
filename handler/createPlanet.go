package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imrochamatheus/challenge-back-end-hit/models"
	"github.com/imrochamatheus/challenge-back-end-hit/utils"
)

func CreatePlanetHandler(ctx *gin.Context) {
	var request CreatePlanetRequest

	err := ctx.BindJSON(&request)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	query, err := utils.ReadQueryFile("./queries/create_planet.sql")
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	defer stmt.Close()

	result, err := stmt.Exec(&request.Name,
		&request.Ground,
		&request.Climate,
		&request.Appearances)

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, http.StatusCreated, "create planet", models.CreatePlanet{
		ID:          id,
		Name:        request.Name,
		Ground:      request.Ground,
		Climate:     request.Climate,
		Appearances: request.Appearances,
	})
}
