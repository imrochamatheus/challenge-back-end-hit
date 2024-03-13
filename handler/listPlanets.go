package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imrochamatheus/challenge-back-end-hit/models"
	"github.com/imrochamatheus/challenge-back-end-hit/utils"
)

func ListPlanetsHandler(ctx *gin.Context) {
	path := "./queries/select_planets.sql"

	query, err := utils.ReadQueryFile(path)

	if err != nil {
		fmt.Printf("unable to read file with query to list planets: %s", err)
		return
	}

	stmt, err := db.Query(query)

	if err != nil {
		fmt.Printf("error list planets: %s", err)
		return
	}

	defer stmt.Close()

	planets := []models.PlanetResponse{}

	for stmt.Next() {
		var row models.PlanetResponse

		if err := stmt.Scan(
			&row.ID,
			&row.Name,
			&row.Ground,
			&row.Climate,
			&row.CreatedAt,
			&row.UpdatedAt,
		); err != nil {
			fmt.Printf("error scanning row: %s", err)
			continue
		}

		planets = append(planets, row)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Success!",
		"data": planets,
	})
}
