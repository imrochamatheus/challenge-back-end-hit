package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imrochamatheus/challenge-back-end-hit/models"
	"github.com/imrochamatheus/challenge-back-end-hit/utils"
)

func ListPlanetsHandler(ctx *gin.Context) {
	path := "./queries/select_planets.sql"

	query, err := utils.ReadQueryFile(path)

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return 
	}

	stmt, err := db.Query(query)
	
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	defer stmt.Close()

	planets := []models.Planet{}

	for stmt.Next() {
		var row models.Planet

		if err := stmt.Scan(
			&row.ID,
			&row.Name,
			&row.Ground,
			&row.Climate,
			&row.Appearances,
			&row.CreatedAt,
			&row.UpdatedAt,
		); err != nil {
			log.Printf("error scanning row: %s", err)
			continue
		}

		planets = append(planets, row)
	}

	sendSuccess(ctx, http.StatusOK, "list planets", planets)
}
