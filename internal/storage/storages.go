package storage

import "avito-test/internal/models"

type HouseStorage interface {
	CreateHouse(models.HouseBuilder) (*models.House, error)
	HouseDesc(uint64) (*models.House, error)

	CreateFlat(models.FlatBuilder, string) (*models.Flat, error)
	UpdateFlatStatus(uint64, string) (*models.Flat, error)

	FlatsByHouseId(uint64) ([]*models.Flat, error)
}

type ValidationStorage interface {
	HouseExists(uint64) (bool, error)
	FlatExists(uint64) (bool, error)
}

type UserStorage interface {
	User()
	RegisterUser()
}
