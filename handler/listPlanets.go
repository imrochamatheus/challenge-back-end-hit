package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imrochamatheus/challenge-back-end-hit/models"
	"github.com/imrochamatheus/challenge-back-end-hit/utils"
)

func buildParameters(ctx *gin.Context, query string, supportedParameters []string ) (string, []interface{}) {
	parameters := make([]interface{}, len(supportedParameters))

	for i, paramName := range supportedParameters {
		if param := ctx.Query(paramName); param != "" {
			query += fmt.Sprintf(" AND " + paramName + " = ?")
			parameters[i] = param
		}
	}

	return query, parameters
}


func ListPlanetsHandler(ctx *gin.Context) {
	queryPath := "./queries/select_planets.sql"
	query, err := utils.ReadQueryFile(queryPath)

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	query, parameters := buildParameters(ctx, query, []string{"id", "name"})
	stmt, err := db.Query(query, parameters...)

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

	if len(planets) == 0 {
		sendError(ctx, http.StatusNotFound, "no planets found")
		return
	}

	sendSuccess(ctx, http.StatusOK, "list planets", planets)
}
