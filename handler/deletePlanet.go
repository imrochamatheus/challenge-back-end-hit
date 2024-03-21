package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imrochamatheus/challenge-back-end-hit/models"
	"github.com/imrochamatheus/challenge-back-end-hit/utils"
)

func getPlanetById(id string) error {
	stmt, err := utils.PrepareQuery(db, "./queries/select_planet_by_id.sql")

	if err != nil {
		return errors.New("planet not found")
	}

	var planet models.Planet

	row := stmt.QueryRow(id)

	if err := row.Scan(planet); err != nil && err == sql.ErrNoRows {
		return errors.New("planet not found")
	}

	return nil
}

func deletePlanet(id string) error {
	stmt, err := utils.PrepareQuery(db, "./queries/delete_planet.sql")

	if err != nil {
		return errors.New("error when prepare query")
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return errors.New("error when delete planet")
	}

	return nil
}

func DeletePlanetHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "id required param")
		return
	}

	err := getPlanetById(id)

	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = deletePlanet(id)

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, http.StatusNoContent, "delete planet", nil)
}
