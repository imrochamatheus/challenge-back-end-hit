package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/imrochamatheus/challenge-back-end-hit/models"
	"github.com/imrochamatheus/challenge-back-end-hit/utils"
)

type swapiPlanetResponse struct {
	Count   int               `json:"count"`
	Results []swapiPlanetInfo `json:"results"`
}

type swapiPlanetInfo struct {
	Films []string `json:"films"`
}

func getSWAPIResponse(name string, client *http.Client) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, "https://swapi.dev/api/planets", nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("search", strings.ToLower(name))

	req.URL.RawQuery = q.Encode()
	res, err := client.Do(req)

	if err != nil {
		return nil, errors.New("Error getting planet appearances")
	}

	return res, nil
}

func parseSWAPIResponse(res *http.Response) (*swapiPlanetResponse, error) {
	var swapiPlanetResponse swapiPlanetResponse

	bodyString, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, errors.New("Error reading response: " + err.Error())
	}

	defer res.Body.Close()

	if err := json.Unmarshal(bodyString, &swapiPlanetResponse); err != nil {
		return nil, errors.New("Error unmarshalling response: " + err.Error())
	}

	return &swapiPlanetResponse, nil
}

func getPlanetAppearances(name *string) (*int64, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	res, err := getSWAPIResponse(*name, client)

	if err != nil {
		return nil, err
	}

	swapiPlanetData, err := parseSWAPIResponse(res)

	if len(swapiPlanetData.Results) == 0 || len(swapiPlanetData.Results[0].Films) == 0 {
		return nil, errors.New("No films found for the planet")
	}

	appearances := int64(len(swapiPlanetData.Results[0].Films))

	return &appearances, nil
}

func checkIfPlanetExists(name string) error {
	query, err := utils.ReadQueryFile("./queries/select_planets.sql")

	if err != nil {
		return err
	}

	query += " AND name COLLATE SQL_Latin1_General_CP1_CI_AS LIKE CONCAT('%', ?, '%')"
	stmt, err := db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	var planet models.Planet

	row := stmt.QueryRow(strings.ToLower(name))

	if err = row.Scan(
		&planet.ID,
		&planet.Name,
		&planet.Ground,
		&planet.Climate,
		&planet.Appearances,
		&planet.CreatedAt,
		&planet.UpdatedAt); err == nil {
		return errors.New("planet already exists")
	}

	return nil
}

func createPlanet(request CreatePlanetRequest, appearances int64) (*sql.Result, error) {
	stmt, err := utils.PrepareQuery(db, "./queries/create_planet.sql")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(&request.Name,
		&request.Ground,
		&request.Climate,
		&appearances)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func CreatePlanetHandler(ctx *gin.Context) {
	var request CreatePlanetRequest

	err := ctx.BindJSON(&request)

	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	appearances, err := getPlanetAppearances(&request.Name)

	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = checkIfPlanetExists(request.Name)

	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	result, err := createPlanet(request, *appearances)

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := (*result).LastInsertId()

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
