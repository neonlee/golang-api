package models

import (
	"time"
)

type Pet struct {
	IdPet       int       `json:"id"`
	ClientId    int       `json:"client_id"`
	Name        string    `json:"name"`
	Race        string    `json:"race"`
	Specie      string    `json:"specie"`
	Cor         string    `json:"cor"`
	Port        string    `json:"port"`
	Weight      float64   `json:"weight"`
	Age         int       `json:"age"`
	Observation string    `json:"observation"`
	Sexy        string    `json:"sexy"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
