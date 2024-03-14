package models

import "time"

type Planet struct {
	ID          int64  	  `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Ground      string    `json:"ground" binding:"required"`
	Climate     string    `json:"climate" binding:"required"`
	Appearances int64    `json:"appearances" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreatePlanet struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Ground      string    `json:"ground" binding:"required"`
	Climate     string    `json:"climate" binding:"required"`
	Appearances int64    `json:"appearances" binding:"required"`
}
