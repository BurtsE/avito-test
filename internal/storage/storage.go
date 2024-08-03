package house

import "avito-test/internal/models"

type HouseStorage interface {
	CreateHouse(models.HouseBuilder) (*models.House, error)
	GetHouseDesc(uint64) (*models.House, error)
}
