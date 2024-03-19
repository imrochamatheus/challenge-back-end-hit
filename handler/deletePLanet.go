package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imrochamatheus/challenge-back-end-hit/utils"
)

func DeletePlanetHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "id required param")
	}

	stmt, err := utils.PrepareQuery(db, "./queries/delete_planet.sql")

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "error when delete planet")
		return
	}

	sendSuccess(ctx, http.StatusNoContent, "delete planet", nil)
}
