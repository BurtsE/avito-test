package models

type HouseBuilder struct {
	Address          *string `json:"address"`
	ConstructionDate *int    `json:"year"`
	Developer        *string `json:"developer,omitempty"`
}
