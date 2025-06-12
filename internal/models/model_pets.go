package models

type Pets struct {
	// ID do pet
	ID int `json:"id" example:"1"`
	// Nome do pet
	Name string `json:"name" example:"Rex"`
	// Tipo do animal
	Type string `json:"type" example:"Cachorro"`
}
