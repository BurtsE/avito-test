package storage

import "avito-test/internal/models"

type HouseStorage interface {
	CreateHouse(models.House) (uint64, error)
	GetHouseDesc(uint64) (models.House, error)
}
