package models

import "time"

type House struct {
	UUID               uint64
	Adress             string
	ConstructionDate   time.Time
	Developer          string
	InitializationDate time.Time
	LastUpdateTime     time.Time
}



