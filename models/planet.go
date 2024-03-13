package models

import "time"

type Planet struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Ground    string    `json:"ground" binding:"required"`
	Climate   string    `json:"climate" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type PlanetResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Ground    string    `json:"ground" binding:"required"`
	Climate   string    `json:"climate" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}