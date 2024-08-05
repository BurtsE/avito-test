package storage

import (
	"avito-test/internal/models"
	"context"
)

type HouseStorage interface {
	CreateHouse(context.Context, models.HouseBuilder) (*models.House, error)
	HouseDesc(context.Context, uint64) (*models.House, error)
	ChangeHouseUpdateTime(context.Context, uint64) error

	Flat(context.Context, uint64) (*models.Flat, error)
	CreateFlat(context.Context, models.FlatBuilder, string) (*models.Flat, error)
	UpdateFlatStatus(context.Context, uint64, string) (*models.Flat, error)

	FlatsByHouseId(context.Context, uint64) ([]*models.Flat, error)
}

type ValidationStorage interface {
	HouseExists(context.Context, uint64) (bool, error)
	FlatExists(context.Context, uint64) (bool, error)
}

type UserStorage interface {
	User(context.Context)
	RegisterUser(context.Context)
}
