package models

import "time"

type House struct {
	UUID               string    `json:"uuid"`
	Address            string    `json:"address"`
	ConstructionDate   time.Time `json:"construction_date"`
	Developer          string    `json:"developer,omitempty"`
	InitializationDate time.Time `json:"initialization_date"`
	LastUpdateTime     time.Time `json:"last_update_time"`
}
