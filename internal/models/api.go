package models

type HouseCreator struct {
	Adress       string `json:"adress,omitempty"`
	CreationDate uint64 `json:"year,omitempty"`
	Developer    string `json:"developer"`
}


