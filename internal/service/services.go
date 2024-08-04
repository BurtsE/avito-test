package service

import "avito-test/internal/models"

type HouseService interface {
	CreateHouse(models.HouseBuilder) (*models.House, error)
	GetHouseDesc(uint64) (*models.House, error)

	CreateFlat(models.FlatBuilder) (*models.Flat, error)
	UpdateFlatStatus(models.FlatStatus) (*models.Flat, error)
}
type ValidationService interface {
	ValidateFlatBuilderData([]byte) (models.FlatBuilder, error)
	ValidateFlatStatusData([]byte) (models.FlatStatus, error)
	ValidateHouseData([]byte) (models.HouseBuilder, error)
	ValidateHouse(uint64) error
	ValidateFlat(uint64) error
	ValidateDummyUserData([]byte) (models.EnumRole, error)
}

type AuthentificationService interface {
	DummyAuthorize(models.EnumRole) (string, error)
	CheckAuthorization([]byte) (models.EnumRole, error)
}
