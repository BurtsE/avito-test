package service

import "avito-test/internal/models"

type HouseService interface {
	CreateHouse(models.HouseBuilder) (*models.House, error)
	GetHouseDesc(uint64) (*models.House, error)

	CreateFlat(models.FlatBuilder) (*models.Flat, error)
	UpdateFlatStatus(models.FlatStatus)(*models.Flat, error)
}
