package handler

type CreatePlanetRequest struct {
	Name      string    `json:"name" binding:"required"`
	Ground    string    `json:"ground" binding:"required"`
	Climate   string    `json:"climate" binding:"required"`
	Appearances int64  	`json:"appearances" binding:"required"`
}