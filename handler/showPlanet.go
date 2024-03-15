package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imrochamatheus/challenge-back-end-hit/models"
	"github.com/imrochamatheus/challenge-back-end-hit/utils"
)

func ShowPlanetHandler(ctx * gin.Context) {
	id := ctx.Param("id")

	if id == ""{
		sendError(ctx, http.StatusBadRequest, "id is required")
	}

	query, err := utils.ReadQueryFile("./queries/select_planet_by_id.sql")
	
	if err != nil{
		sendError(ctx, http.StatusInternalServerError, err.Error())
	}

	stmt, err := db.Prepare(query)

	if err != nil{
		sendError(ctx, http.StatusInternalServerError, err.Error())
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	if err != nil{
		sendError(ctx, http.StatusBadRequest, err.Error())
	}
	
	var planet models.Planet

 	if err = row.Scan(
		&planet.ID,
		&planet.Name,
		&planet.Ground,
		&planet.Climate,
		&planet.Appearances,
		&planet.CreatedAt,
		&planet.UpdatedAt); err != nil{
			sendError(ctx, http.StatusNotFound, "no results math the id provided")
			return
		}

	sendSuccess(ctx, http.StatusOK, "show planet", planet)
	return
}