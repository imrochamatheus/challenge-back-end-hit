package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/imrochamatheus/challenge-back-end-hit/utils"
)

type swapiPlanetResponse struct {
	Count   int               `json:"count"`
	Results []swapiPlanetInfo `json:"results"`
}

type swapiPlanetInfo struct {
	Films []string `json:"films"`
}

func getPlanetAppearances(name *string) (*int64, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, "https://swapi.dev/api/planets", nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("search", strings.ToLower(*name))

	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)

	if err != nil {
		return nil, errors.New("Error getting planet appearances")
	}

	defer res.Body.Close()

	var planetResponse swapiPlanetResponse

	bodyString, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, errors.New("Error reading response: " + err.Error())
	}

	if err := json.Unmarshal(bodyString, &planetResponse); err != nil {
		return nil, errors.New("Error unmarshalling response: " + err.Error())
	}

	appearances := int64(len(planetResponse.Results[0].Films))

	if len(planetResponse.Results) == 0|| len(planetResponse.Results[0].Films) == 0 {
		return nil, errors.New("No films found for the planet")
	}

	return &appearances, nil
}

func CreatePlanetHandler(ctx *gin.Context) {
	var request CreatePlanetRequest

	err := ctx.BindJSON(&request)

	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	stmt, err := utils.PrepareQuery(db, "./queries/create_planet.sql")

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	defer stmt.Close()

	appearances, err := getPlanetAppearances(&request.Name)

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	result, err := stmt.Exec(&request.Name,
		&request.Ground,
		&request.Climate,
		&appearances)

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := result.LastInsertId()

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, http.StatusCreated, "create planet", CreatePlanetResponse{
		ID:          id,
		Name:        request.Name,
		Ground:      request.Ground,
		Climate:     request.Climate,
		Appearances: *appearances,
	})
}
