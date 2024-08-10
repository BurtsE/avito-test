package service

import (
	"avito-test/internal/models"
	"context"
)

type HouseService interface {
	CreateHouse(context.Context, models.HouseBuilder) (models.House, error)
	HouseDesc(context.Context, uint64) (models.House, error)

	CreateFlat(context.Context, models.FlatBuilder) (models.Flat, error)
	UpdateFlatStatus(context.Context, models.FlatStatus) (models.Flat, error)

	HouseFlats(context.Context, uint64) ([]*models.Flat, error)
}
type ValidationService interface {
	ValidateFlatBuilderData(context.Context, []byte) (models.FlatBuilder, error)
	ValidateFlatStatusData(context.Context, []byte) (models.FlatStatus, error)
	ValidateHouseData(context.Context, []byte) (models.HouseBuilder, error)
	ValidateHouse(context.Context, uint64) error
	ValidateFlat(context.Context, uint64) error
	ValidateDummyUserData(context.Context, []byte) (models.EnumRole, error)
}

type AuthentificationService interface {
	DummyAuthorize(context.Context, models.EnumRole) (string, error)
	CheckAuthorization(context.Context, []byte) (models.User, error)
}
